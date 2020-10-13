package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello OpenShift")
}

func livenessProbe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health OK")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", livenessProbe)
	http.ListenAndServe(":8080", nil)
}
