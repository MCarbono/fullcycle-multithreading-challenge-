package gateway

import "strings"

type CepServiceGateway interface {
	GetCEP(out chan<- CEPResponseGateway, errChan chan<- error)
}

type CEPResponseGateway struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	ServiceURL  string `json:"service-api-url"`
}

func NewCEP(cep, logradouro, complemento, bairro, localidade, uf, serviceURL string) CEPResponseGateway {
	return CEPResponseGateway{
		ServiceURL:  "URL da api para coleta dos dados: " + serviceURL,
		Cep:         cep,
		Logradouro:  logradouro,
		Complemento: complemento,
		Bairro:      bairro,
		Localidade:  localidade,
		Uf:          uf,
	}
}

func (c CEPResponseGateway) String() string {
	var s strings.Builder
	s.WriteString("-----------------------------------------------------------------------------------\n")
	s.WriteString(c.ServiceURL + "\n")
	s.WriteString("CEP: " + c.Cep + "\n")
	s.WriteString("Logradouro: " + c.Logradouro + "\n")
	s.WriteString("Complemento: " + c.Complemento + "\n")
	s.WriteString("Bairro: " + c.Bairro + "\n")
	s.WriteString("Localidade: " + c.Localidade + "\n")
	s.WriteString("UF: " + c.Uf + "\n")
	s.WriteString("-----------------------------------------------------------------------------------")
	return s.String()
}
