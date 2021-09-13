package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var addr string

func main() {
	flag.StringVar(&addr, "addr", ":9090", "The address to listen on for HTTP requests.")
	flag.Parse()

	const endpointEnv = "ECS_CONTAINER_METADATA_URI_V4"
	endpoint := os.Getenv(endpointEnv)
	if endpoint == "" {
		log.Fatalf("%q environmental variable is not set, are you running this on ECS?", endpointEnv)
	}

	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("Failed to parse endpoint: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(endpointURL)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %v", r.URL)
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
