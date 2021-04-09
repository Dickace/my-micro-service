package transport

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Router() *mux.Router {
	r:= mux.NewRouter()
	s:= r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/health", ServerOK).Methods(http.MethodHead)
	s.HandleFunc("/arithmetic", func(writer http.ResponseWriter, request *http.Request) {
		requestS, _:= ioutil.ReadAll(request.Body)
		log.Print(requestS)
		str:= string(requestS)
		log.Print(str)
		fmt.Fprintf(writer,"4")
	}).Methods(http.MethodPost)
	return r
}

func ServerOK(w http.ResponseWriter, _ *http.Request){

}

func Arithmetic(w http.ResponseWriter, r *http.Request)  {
	requS, err:= ioutil.ReadAll(r.Body)
	log.Print(requS)
	if err != nil{
		log.Fatal(err)
	}
	str:= string(requS)
	s:= strings.Fields(str)
	sum, err:= strconv.ParseInt(s[0],10,64)
	for i:= 1; i < len(s)+1;i+=2{
		switch s[i] {
		case "+":
			v, err:= strconv.ParseInt(s[i+1],10,64)
			if err!=nil{
				log.Fatal(err)
			}
			sum+= v
		case "-":
			v, err:= strconv.ParseInt(s[i+1],10,64)
			if err!=nil{
				log.Fatal(err)
			}
			sum-= v
		}
	}
	fmt.Fprintf(w, strconv.FormatInt(sum, 10))
}