package main

import (
	"log"
	"net/http"
)

func exampleServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html>\n<body>\n"))
	w.Write([]byte("<p>Hello OpenShift!</p>\n"))
	w.Write([]byte("</body>\n</html>\n"))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("foo\n"))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("bar\n"))
}

func readinessProbe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK\n"))
}

func livenessProbe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK\n"))
}

func main() {
	http.HandleFunc("/", exampleServer)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/ready", readinessProbe)
	http.HandleFunc("/alive", livenessProbe)
	log.Println("Starting http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
