package main

import (
	"fmt"
	"net/http"
)

type Cat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cats = []Cat{
	{Name: "Figo", Age: 11},
	{Name: "Liza", Age: 6},
	{Name: "Tosia", Age: 7},
}

func main() {
	router := http.NewServeMux()
	fmt.Println("Listening on :2137...")
	router.HandleFunc("/", indexFunc)
	router.HandleFunc("/cats", getAllCats)
	router.HandleFunc("/cats/{name}", getCatById)
	http.ListenAndServe(":2137", logger(jsonHeader(router)))
}
