package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type Conversion struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

var (
	timeoutError      = errors.New("transaction exceeded time limit")
	retrieveDataError = errors.New("fail to retrieve exchange data")
)

type app struct {
	db *gorm.DB
}

func New(db *gorm.DB) *app {
	return &app{
		db: db,
	}
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalln(err)
	}
	app := New(db)

	http.HandleFunc("/cotacao", app.ConsumeExchange)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (a *app) ConsumeExchange(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://economia.awesomeapi.com.br/json/last/USD-BRL",
		nil,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("unexpected error")
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: retrieveDataError.Error(),
		})
		return
	}
	defer resp.Body.Close()

	var conv Conversion

	if err := json.NewDecoder(resp.Body).Decode(&conv); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)

		return
	}

	bid, err := strconv.ParseFloat(conv.USDBRL.Bid, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	price := Price{Bid: bid}

	if err := a.create(ctx, &price); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(price)
	return
}
