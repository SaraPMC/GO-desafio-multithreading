package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AddressViaCEPResponse struct {
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func GetAddressFromViaCEPApi(cep string) (AddressViaCEPResponse, error) {
	url := "https://viacep.com.br/ws/%s/json/"

	response, err := http.Get(fmt.Sprintf(url, cep))
	if err != nil {
		return AddressViaCEPResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return AddressViaCEPResponse{}, fmt.Errorf("Erro ao buscar endereço: %s", response.Status)
	}

	var addressData AddressViaCEPResponse
	if err := json.NewDecoder(response.Body).Decode(&addressData); err != nil {
		return AddressViaCEPResponse{}, fmt.Errorf("Erro ao decodificar JSON: %v", err)
	}

	return addressData, nil
}
