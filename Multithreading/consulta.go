package Multithreading

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func BuscaCEP(cep string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ch := make(chan Consulta)

	go func() {
		err := ConsultaViaCEP(ctx, cep, ch)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		err := ConsultaApiCep(ctx, cep, ch)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	consulta := <-ch
	cancel()
	close(ch)

	log.Println(consulta)
}

func ConsultaViaCEP(ctx context.Context, cep string, ch chan<- Consulta) error {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep)

	fmt.Println("inicio consulta Via CEP")

	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	log.Printf("Consulta Via CEP tempo de resposta: %s\n", time.Since(start))

	var payload ViaCEPResponse
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		log.Fatalln(err)
	}

	select {
	case <-ctx.Done():
		log.Println("Consulta via cep cancelada")
		return nil
	default:
		consulta := Consulta{
			Buscador:   "Via CEP",
			CEP:        payload.CEP,
			UF:         payload.UF,
			Cidade:     payload.Localidade,
			Bairro:     payload.Bairro,
			Logradouro: payload.Logradouro,
		}

		ch <- consulta
		return nil
	}
}

func ConsultaApiCep(ctx context.Context, cep string, ch chan<- Consulta) error {
	cep = fmt.Sprintf("%s-%s", cep[:5], cep[5:])

	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)

	fmt.Println("inicio consulta Api CEP")

	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	log.Printf("Consulta Api CEP tempo de resposta: %s\n", time.Since(start))

	var payload ConsultaAPIReponse
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		log.Fatalln(err)
	}

	select {
	case <-ctx.Done():
		log.Println("Consulta api cep cancelada")
		return nil
	default:
		consulta := Consulta{
			Buscador:   "Api CEP",
			CEP:        payload.Code,
			UF:         payload.State,
			Cidade:     payload.City,
			Bairro:     payload.District,
			Logradouro: payload.Address,
		}

		ch <- consulta
		return nil
	}
}
