package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("HOSTNAME_PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		hostname, _ := os.Hostname()
		rw.Write([]byte(hostname))
	})
	http.ListenAndServe(":"+port, nil)
}
