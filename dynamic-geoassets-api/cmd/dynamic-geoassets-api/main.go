package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/metno/S-ENDA-documentation/dynamic-geoassets-api/internal/app"
)

func main() {
	templates := template.Must(template.ParseGlob("templates/*"))

	registrationService := app.NewService(templates, app.StaticStore)

	log.Println("Starting Registration web service...")
	go func() {
		http.ListenAndServe(":8088", registrationService.InternalRouter)
	}()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      registrationService.ExternalRouter,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
