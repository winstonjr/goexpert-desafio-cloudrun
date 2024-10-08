package main

import (
	"encoding/json"
	"fmt"
	"github.com/winstonjr/goexpert-desafio-cloudrun/configs"
	"github.com/winstonjr/goexpert-desafio-cloudrun/internal/infra/integration"
	"github.com/winstonjr/goexpert-desafio-cloudrun/internal/usecase"
	"log"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	viacepIntegration := integration.NewViacepIntegration()
	weatherapiIntegration := integration.NewWeatherapiIntegration(config.WeatherApiKey)
	checkWeatherUseCase := usecase.NewCheckWeatherUseCase(weatherapiIntegration, viacepIntegration)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cep := r.URL.Query().Get("cep")
		if cep == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = w.Write([]byte(`invalid zipcode`))
			return
			log.Println("erro 422")
		}
		log.Println("cep acquired", cep)
		temperature, err := checkWeatherUseCase.Execute(cep)
		if err != nil {
			log.Println("temperature error", err.Error())
			if err.Error() == "invalid zipcode" {
				log.Println("temperature error", 422)
				w.WriteHeader(http.StatusUnprocessableEntity)
				_, _ = w.Write([]byte(`invalid zipcode`))
				return
			} else if err.Error() == "can not find zipcode" {
				log.Println("temperature error", 404)
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`can not find zipcode`))
				return
			} else {
				log.Println("temperature error", 500)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		err = json.NewEncoder(w).Encode(temperature)
		if err != nil {
			log.Println("write to buffer error", 500)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	fmt.Println("Listening on port :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	//for _, cep := range os.Args[1:] {
	//	temperature, err := checkWeatherUseCase.Execute(cep)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	data, err := json.Marshal(temperature)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "Erro ao realizar parse do json: %v\n", err)
	//	}
	//	fmt.Println(string(data))
	//}
}
