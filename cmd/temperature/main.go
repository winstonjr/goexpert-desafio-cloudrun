package main

import (
	"encoding/json"
	"fmt"
	"github.com/winstonjr/goexpert-desafio-cloudrun/internal/infra/integration"
	"github.com/winstonjr/goexpert-desafio-cloudrun/internal/usecase"
	"os"
)

func main() {
	viacepIntegration := integration.NewViacepIntegration()
	weatherapiIntegration := integration.NewWeatherapiIntegration("23b020471bba461680101942241309")
	checkWeatherUseCase := usecase.NewCheckWeatherUseCase(weatherapiIntegration, viacepIntegration)
	for _, cep := range os.Args[1:] {
		temperature, err := checkWeatherUseCase.Execute(cep)
		if err != nil {
			panic(err)
		}

		data, err := json.Marshal(temperature)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao realizar parse do json: %v\n", err)
		}
		fmt.Println(string(data))
	}
}
