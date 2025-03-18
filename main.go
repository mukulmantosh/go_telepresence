package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", PortLogger(HomeHandler))
	http.HandleFunc("/load_dataset", PortLogger(DataHandler))
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
