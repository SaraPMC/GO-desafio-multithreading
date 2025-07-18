package main

import (
	"desafio-multithread/apis"
	"fmt"
	"time"
)

type AddressResponse struct {
	API        string `json:"api"`
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
}

func main() {
	cep := "03077-000"
	ch := make(chan AddressResponse, 2)

	go func() {
		for {
			addressViaCEPApiResponse, errViaCEPApi := apis.GetAddressFromViaCEPApi(cep)
			if errViaCEPApi == nil {
				var addressResponse = AddressResponse{
					API:        "ViaCEPApi",
					CEP:        cep,
					Logradouro: addressViaCEPApiResponse.Logradouro,
					Bairro:     addressViaCEPApiResponse.Bairro,
					Cidade:     addressViaCEPApiResponse.Localidade,
					Estado:     addressViaCEPApiResponse.UF,
				}
				ch <- addressResponse
				break
			}
		}
	}()

	go func() {
		for {
			addressBrasilApiResponse, errBrasilApi := apis.GetAddressFromBrasilApi(cep)
			if errBrasilApi == nil {
				var addressResponse = AddressResponse{
					API:        "BrasilApi",
					CEP:        cep,
					Logradouro: addressBrasilApiResponse.Street,
					Bairro:     addressBrasilApiResponse.Neighborhood,
					Cidade:     addressBrasilApiResponse.City,
					Estado:     addressBrasilApiResponse.State,
				}
				ch <- addressResponse
				break
			}
		}
	}()

	select {
	case response1 := <-ch:
		fmt.Printf("Resposta obtida através da API: %s \n CEP: %s\n Logradouro: %s\n Bairro: %s\n Cidade: %s\n Estado: %s\n",
			response1.API, response1.CEP, response1.Logradouro, response1.Bairro, response1.Cidade, response1.Estado)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout: Nenhuma resposta recebida dentro de 1 segundo.")
	}
}
