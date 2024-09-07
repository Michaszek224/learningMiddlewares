package main

import (
	"encoding/json"
	"net/http"
)

func indexFunc(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World!, go to /cats")
}

func getAllCats(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cats)
}

func getCatById(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	for _, v := range cats {
		if v.Name == name {
			json.NewEncoder(w).Encode(v)
		}
	}
}
