package main

import (
	"cars/config"
	"cars/handlers"
	middleware "cars/middlewares"
	"fmt"
	"net/http"
)

func main() {

	config.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/cars", handlers.CarHandler)
	mux.HandleFunc("/cars/", handlers.CarHandler)

	wrapperMux := middleware.Logger(mux)
	wrapperMux = middleware.SecurityHeaders(wrapperMux)


	// http.HandleFunc("/cars", handlers.CarHandler)
	// http.HandleFunc("/cars/", handlers.CarHandler)
	fmt.Println("http listening")
	http.ListenAndServe(":3015", wrapperMux)
}