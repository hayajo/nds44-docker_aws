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
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(hostname))
	})
	http.ListenAndServe(":"+port, nil)
}
