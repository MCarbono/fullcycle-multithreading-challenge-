package gateway

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type CNDAPICEP struct {
	Endpoint string
}

func NewCNDAPICEP(cep string) *CNDAPICEP {
	return &CNDAPICEP{
		Endpoint: "https://cdn.apicep.com/file/apicep/" + cep + ".json",
	}
}

type CDNAPICEPGatewayResponse struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func (v *CNDAPICEP) GetCEP(out chan<- CEPResponseGateway, errChan chan<- error) {
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
	var cep CDNAPICEPGatewayResponse
	err = json.Unmarshal(body, &cep)
	if err != nil {
		errChan <- err
		return
	}
	out <- NewCEP(cep.Code, cep.Address, "", "", cep.City, cep.State, v.Endpoint)
}
