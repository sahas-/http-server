package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func beautifyJSON(jsonStr string) (string, error) {
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return "", err
	}

	prettyJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		log.Println("Error marshalling JSON into indented format:", err)
		return "", err
	}

	return string(prettyJSON), nil
}
func main() {
	// Define the port on which the server will listen
	port := 8888

	// Define a simple handler function for incoming HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Dump the request information to the console
		body, err := io.ReadAll(io.Reader(r.Body))
		if err != nil {
			fmt.Println("Error reading request body:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Beautify the JSON data
		prettyJSON, err := beautifyJSON(string(body))
		if err != nil {
			fmt.Println("Error beautifying JSON:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print the beautified JSON to the console
		fmt.Println(prettyJSON)
		// Respond with a simple message
		fmt.Fprintf(w, "Hello, this is the Golang REST API!")
	})

	// Start the HTTP server
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
