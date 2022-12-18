package gateway

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type VIACEP struct {
	Endpoint string
}

func NewVIACEP(cep string) *VIACEP {
	return &VIACEP{
		Endpoint: "https://viacep.com.br/ws/" + cep + "/json/",
	}
}

type VIACEPGatewayResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v *VIACEP) GetCEP(out chan<- CEPResponseGateway, errChan chan<- error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer close(out)
	defer close(errChan)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", v.Endpoint, nil)
	if err != nil {
		errChan <- err
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}
	var cep VIACEPGatewayResponse
	err = json.Unmarshal(body, &cep)
	if err != nil {
		errChan <- err
		return
	}
	out <- NewCEP(cep.Cep, cep.Logradouro, cep.Complemento, cep.Bairro, cep.Localidade, cep.Uf, v.Endpoint)
}
