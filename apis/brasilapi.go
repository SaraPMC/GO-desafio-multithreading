package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AddressBrasilApiResponse struct {
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

func GetAddressFromBrasilApi(cep string) (AddressBrasilApiResponse, error) {
	url := "https://brasilapi.com.br/api/cep/v1/%s"

	response, err := http.Get(fmt.Sprintf(url, cep))
	if err != nil {
		return AddressBrasilApiResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return AddressBrasilApiResponse{}, fmt.Errorf("Erro ao buscar endereço: %s", response.Status)
	}

	var addressData AddressBrasilApiResponse
	if err := json.NewDecoder(response.Body).Decode(&addressData); err != nil {
		return AddressBrasilApiResponse{}, fmt.Errorf("Erro ao decodificar JSON: %v", err)
	}

	return addressData, nil
}
