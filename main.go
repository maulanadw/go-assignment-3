package main

import (
	"fmt"
	"go-assignment-3/controllers"
	"go-assignment-3/repositories"
	"go-assignment-3/services"
	"net/http"
)

const PORT = ":8080"

func main() {
	weatherRepo := repositories.NewWeatherRepository()
	weatherService := services.NewWeatherService(weatherRepo)
	weatherController := controllers.NewWeatherController(weatherService)

	http.HandleFunc("/", weatherController.GetWeatherMetric)
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))
	fmt.Println("Server running on port:", PORT)
	http.ListenAndServe(PORT, nil)
}
