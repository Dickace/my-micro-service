package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port:= os.Getenv("PORT")

	http.HandleFunc("/start", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Start")
	} )

	http.ListenAndServe(":"+port, nil)
}

