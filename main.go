package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func parseHostAndPort(hostPort string) (string, string, error) {
	// Check if the hostPort contains ":"
	if strings.Contains(hostPort, ":") {
		host, port, err := net.SplitHostPort(hostPort)
		if err != nil {
			return "", "", err
		}
		return host, port, nil
	}
	// If no port is specified, assume it's empty but separate host
	return hostPort, "", nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	host, port, err := parseHostAndPort(r.Host)
	if err != nil {
		log.Printf("Error parsing host and port: %v", err)
		host = "unknown"
		port = "unknown"
	}

	log.Printf("Received request host=\"%s\" port=\"%s\"", host, port)
	fmt.Fprintf(w, "Hello Go!")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
