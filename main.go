package main

import (
	"bufio"
	"errors"
	"fmt"
	"fullcycle-multithreading-challenge/gateway"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}

	ch := make(chan gateway.CEPResponseGateway)
	errCh := make(chan error)

	viaCEPServiceGateway := gateway.NewVIACEP(*input)
	cdnCEPApiServiceGateway := gateway.NewCNDAPICEP(*input)
	go viaCEPServiceGateway.GetCEP(ch, errCh)
	go cdnCEPApiServiceGateway.GetCEP(ch, errCh)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case msg := <-errCh:
		fmt.Println(msg)
	}
}

func readInput() (*string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Para saber mais informações sobre um determinado cep, digite apenas números em um desses formatos: 00000000 ou 00000-000")
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	input = strings.Replace(input, "\n", "", -1)
	input, err = validateInput(input)
	return &input, err
}

func validateInput(input string) (string, error) {
	r := regexp.MustCompile(`[0-9]{5}[0-9]{3}`)
	if r.Match([]byte(input)) {
		return input[:5] + "-" + input[5:], nil
	}
	r = regexp.MustCompile(`[0-9]{5}-[0-9]{3}`)
	if r.Match([]byte(input)) {
		return input, nil
	}
	return input, errors.New("CEP inválido")
}
