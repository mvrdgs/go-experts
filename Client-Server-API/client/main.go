package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Exchange struct {
	Bid   float64 `json:"bid,omitempty"`
	Error string  `json:"error,omitempty"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer res.Body.Close()

	var exchange Exchange
	err = json.NewDecoder(res.Body).Decode(&exchange)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if exchange.Error != "" {
		log.Fatalln(exchange.Error)
	}

	if _, err := f.Write([]byte(fmt.Sprintf("Dólar: R$ %.2f\n", exchange.Bid))); err != nil {
		log.Fatalln(err)
	}
}
