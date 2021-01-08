// [START container_batman]
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// register batman function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", batman)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
}

// batman responds to the request with a plain-text.
func batman(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "A wealthy American playboy\n")
}

// [END container_batman]
