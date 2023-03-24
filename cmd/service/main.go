package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"tweet_compression_registry/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	api := app.NewApi(a)
	showRoutes(&api)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), api.Router))
}

func showRoutes(api *app.Api) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(api.Router, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}
