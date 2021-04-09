package main

import (
	"myMicroService/pkg/transport"
	"net/http"
	"os"
)

func main() {
	port:= os.Getenv("PORT")

	r:= transport.Router()
	http.ListenAndServe(":"+port, r)
}

