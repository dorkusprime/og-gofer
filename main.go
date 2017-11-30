package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strings"
	"github.com/dorkusprime/og-gofer/og-gofer"
	"os"
)

var (
	listenPort = envOrDefault("PORT", "8080")
)

// Takes the key name and the default value of that environmental variable
// and returns the set environment variable or, if not set, the default
func envOrDefault(keyName, default_value string) (envValue string) {
	if e := os.Getenv(keyName); len(e) == 0 {
		envValue = default_value
	} else {
		envValue = e
	}
	return
}

// HTTP Handler function – scrapes, builds response, and responds
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()

	responseObject := ogGofer.Gofer(strings.Join(query["url"], ""))
	response, _ := json.Marshal(responseObject)

	w.Write(response)
}


func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("OG Gofer Server Listening on Port %s . . . \n\n", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil)
	if err != nil {
		panic(err)
	}
}
