package transport

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r:= mux.NewRouter()
	s:= r.PathPrefix("/api/v1/").Subrouter()
	s.HandleFunc("/", ServerOK).Methods(http.MethodHead)
	return r
}

func ServerOK(w http.ResponseWriter, _*http.Request){

}