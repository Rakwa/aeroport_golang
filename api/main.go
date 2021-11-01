package main

import (
	"api/controller"
	_ "api/docs"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Airport API
// @version 1.0
// @description Get information about airports
// @host localhost:3333
// @BasePath /api

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/airports", controller.GetAirports).Methods("GET")
	router.HandleFunc("/api/airports/{id}/measures", controller.GetMeasures).Methods("GET")
	router.HandleFunc("/api/airports/{id}/averages", controller.GetAverages).Methods("GET")
	router.PathPrefix("").Handler(httpSwagger.WrapHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("API availabe in localhost:3333")
	log.Fatal(http.ListenAndServe(":3333", handler))

}
