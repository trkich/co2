package main

import (
	"co2/common"
	"co2/services/api/sensors/handlers"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()

	common.LoadEnv()

	var router = mux.NewRouter()
	var api = router.PathPrefix("/api").Subrouter()

	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	var apiv1 = api.PathPrefix("/v1").Subrouter()

	apiv1.HandleFunc("/sensors/{uuid}/measurements", handlers.HandleMeasurements)
	apiv1.HandleFunc("/sensors/{uuid}", handlers.HandleStatus)
	apiv1.HandleFunc("/sensors/{uuid}/metrics", handlers.HandleMetrics)

	apiv1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	log.Println("API service started at port "+os.Getenv("API_PORT"))

	http.ListenAndServe(":"+os.Getenv("API_PORT"), router)
}