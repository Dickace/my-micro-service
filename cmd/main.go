package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT", "4444")
	port:= os.Getenv("PORT")

	http.HandleFunc("/start", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Start")
	} )

	http.ListenAndServe(":"+port, nil)
}

