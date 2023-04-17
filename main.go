package main

import (
	"log"
	"net/http"

	"github.com/jordation/layermon/api"
)

type Application struct {
	API *api.API
}

var app Application

func InitApp() {
	app.API = api.GetApi()
}

func startREST() {
	http.HandleFunc("/", app.API.HandleListStats)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	InitApp()
	go api.StartGrpcServer()
	go startREST()
	select {}
}
