package services

import (
	"go-assignment-3/models"
	"go-assignment-3/params"
	"go-assignment-3/repositories"
)

type WeatherService interface {
	GetWeatherMetric() (*params.Weather, error)
	SetWeatherMetric(request params.Weather) error
}

type weatherService struct {
	weatherRepo repositories.WeatherRepository
}

func NewWeatherService(wr repositories.WeatherRepository) WeatherService {
	return &weatherService{
		weatherRepo: wr,
	}
}

func (ws *weatherService) GetWeatherMetric() (*params.Weather, error) {
	res, err := ws.weatherRepo.GetWeatherMetric()
	if err != nil {
		return nil, err
	}

	weather := params.Weather{
		Water: res.Status.Water,
		Wind:  res.Status.Wind,
	}

	return &weather, nil

}

func (ws *weatherService) SetWeatherMetric(request params.Weather) error {
	weatherModel := models.Weather{
		Status: models.Status{
			Water: request.Water,
			Wind:  request.Wind,
		},
	}

	return ws.weatherRepo.SetWeatherMetric(weatherModel)
}
