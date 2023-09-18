package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Sotnasjeff/cep-multithreading-api/internal/dto"
	"github.com/go-chi/chi"
)

func GetAddress(w http.ResponseWriter, r *http.Request) {
	api1 := make(chan dto.GetBrasilApiOutput)
	api2 := make(chan dto.GetViaCepOutput)

	cep := chi.URLParam(r, "cep")

	go func() {
		req, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalf("Error in reading body %v", err)
		}
		var brasilApiResponse dto.GetBrasilApiOutput
		err = json.Unmarshal(body, &brasilApiResponse)
		if err != nil {
			log.Fatalf("Error in parsing response %v", err)
		}
		api1 <- brasilApiResponse

	}()

	go func() {
		req, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalf("Error in reading body %v", err)
		}
		var viaCepApiResponse dto.GetViaCepOutput
		err = json.Unmarshal(body, &viaCepApiResponse)
		if err != nil {
			log.Fatalf("Error in parsing response %v", err)
		}
		api2 <- viaCepApiResponse
	}()

	select {
	case brasilApi := <-api1:
		fmt.Printf("Response vindo da Brasil Api -> \nCep: %s\n State: %s\n Service: %s\n City: %s\n Street: %s\n", brasilApi.Cep, brasilApi.State, brasilApi.Service, brasilApi.City, brasilApi.Street)
		return
	case viaCep := <-api2:
		fmt.Printf("Response vindo da Via Cep -> \nCep: %s\n Logradouro:%s\n Localidade: %s\n Bairro: %s\n Complemento: %s\n", viaCep.Cep, viaCep.Logradouro, viaCep.Localidade, viaCep.Bairro, viaCep.Complemento)
		return
	case <-time.After(time.Second * 1):
		fmt.Printf("\nTimeout")
		return
	}
}
