package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	server "weather-service/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	address :=  ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/", server.HandleCurrentWeatherResponse)

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("server is running on port %s", address)
	fmt.Println("weather service is up and running!", address)
	log.Fatal(s.ListenAndServe())
}

