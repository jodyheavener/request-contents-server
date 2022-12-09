package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

type requestInfo struct {
	Path        string            `json:"path"`
	QueryParams map[string]string `json:"queryParams"`
	Headers     map[string]string `json:"headers"`
}

func main() {
	// Get the port number from the command line arguments,
	// or use the default of 2416
	port := 2416
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Invalid port number:", os.Args[1])
		}
	}

	log.Println("Listening on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a requestInfo struct with the request info.
		info := requestInfo{
			Path:    r.URL.Path,
			Headers: map[string]string{},
		}

		// Add the query parameters to the struct.
		info.QueryParams = map[string]string{}
		for key, values := range r.URL.Query() {
			info.QueryParams[key] = values[0]
		}

		// Add the request headers to the struct.
		for key, values := range r.Header {
			info.Headers[key] = values[0]
		}

		// Marshal the struct to JSON and write it to the response.
		response, err := json.Marshal(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
