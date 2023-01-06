package main

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json := simplejson.New()
		json.Set("name", "hari")

		payload, err := json.MarshalJSON()
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	})

	router.HandleFunc("/another", func(w http.ResponseWriter, r *http.Request) {
		json := simplejson.New()
		json.Set("another", "okay")

		payload, err := json.MarshalJSON()
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	})

	http.ListenAndServe(":3000", router)
}
