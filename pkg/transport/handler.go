package transport

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r:= mux.NewRouter()
	s:= r.PathPrefix("/api/v1/").Subrouter()
	s.HandleFunc("/", ServerOK).Methods(http.MethodHead)
	s.HandleFunc("/api/v1", Arithmetic)
	return r
}

func ServerOK(w http.ResponseWriter, _ *http.Request){

}

func Arithmetic(w http.ResponseWriter, r *http.Request)  {

}