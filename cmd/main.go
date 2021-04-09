package main

import (
	"fmt"
	"net/http"
)

func main() {
	port:= ":8000"

	http.HandleFunc("/start", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Start")
	} )

	http.ListenAndServe(port, nil)
}

