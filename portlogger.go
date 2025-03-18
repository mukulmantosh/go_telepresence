package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func PortLogger(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		host, port, err := parseHostAndPort(r.Host)
		if err != nil {
			log.Printf("Error parsing host and port: %v", err)
			host = "unknown"
			port = "unknown"
		}
		log.Printf("Received request host=\"%s\" port=\"%s\"", host, port)

		if err := f(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(
				map[string]string{"message": "Error handling request! Something went wrong."},
			)
		}

	}

}

func parseHostAndPort(hostPort string) (string, string, error) {
	if strings.Contains(hostPort, ":") {
		host, port, err := net.SplitHostPort(hostPort)
		if err != nil {
			return "", "", err
		}
		return host, port, nil
	}
	return hostPort, "", nil
}
