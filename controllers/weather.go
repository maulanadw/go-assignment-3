package controllers

import (
	"encoding/json"
	"go-assignment-3/params"
	"go-assignment-3/services"
	"math/rand"
	"net/http"
	"text/template"
)

type WeatherController interface {
	GetWeatherMetric(w http.ResponseWriter, r *http.Request)
}

type weatherController struct {
	weatherService services.WeatherService
}

func NewWeatherController(ws services.WeatherService) WeatherController {
	return &weatherController{
		weatherService: ws,
	}
}

func (wc *weatherController) GetWeatherMetric(w http.ResponseWriter, r *http.Request) {
	var errorResponse map[string]interface{}
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		metric := params.Weather{
			Water: uint(rand.Intn(100)),
			Wind:  uint(rand.Intn(100)),
		}

		err = wc.weatherService.SetWeatherMetric(metric)
		if err != nil {
			errorResponse["error"] = err.Error()
			jsonResp, err := json.Marshal(errorResponse)
			if err != nil {
				panic(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		}

		data, err := wc.weatherService.GetWeatherMetric()
		if err != nil {
			errorResponse["error"] = err.Error()
			jsonResp, err := json.Marshal(errorResponse)
			if err != nil {
				panic(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		}
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
