package main

import (
	"net/http"
	"os"
)

func main() {
	port:= os.Getenv("PORT")

	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, _ *http.Request) {

	} )

	http.ListenAndServe(":"+port, nil)
}

