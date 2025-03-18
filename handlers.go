package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) error {
	return json.NewEncoder(w).Encode(map[string]string{"message": "Hello Go!"})
}

func DataHandler(w http.ResponseWriter, _ *http.Request) error {
	source := rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().UnixNano()))
	r := rand.New(source)
	randomNumber := r.IntN(10) + 1

	if os.Getenv("ENV") == "production" {
		if randomNumber > 5 {
			log.Println("Something failed at line number 20")
			return fmt.Errorf("%s", "Failed!")
		}
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(map[string]string{"message": "Data Loaded Successfully!"})
}
