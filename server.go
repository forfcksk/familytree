package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Family struct {
	Name   string  `json:"name"`
	SName  string  `json:"sname"`
	Parent *Parent `json:"parent"`
}

type Parent struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var family []Family

func getFamilys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(family)
}

func getFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range family {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Family{})
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

func main() {
	r := mux.NewRouter()
	family = append(family, Family{Name: "Danila", SName: "Nazarov", Parent: &Parent{Firstname: "Roman", Lastname: "Nazarov"}})
	family = append(family, Family{Name: "Roman", SName: "Nazarov", Parent: &Parent{Firstname: "Olga", Lastname: "Nazarova"}})
	r.HandleFunc("/family", getFamilys).Methods("GET")
	r.HandleFunc("/family/{name}", getFamily).Methods("GET")
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/check", NotImplemented).Methods("GET")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8081", r))
}
