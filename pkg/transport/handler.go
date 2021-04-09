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
	s.HandleFunc("/arithmetic", Arithmetic).Methods(http.MethodPost)
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
	s:= strings.Split(str, "+")
	sum:= int64(0)
	for i:=0; i< len(s);i++{
		t, _ := strconv.ParseInt(s[i], 10,64)
		sum += t
	}


	fmt.Fprintf(w, strconv.Itoa(int(sum)))
}

func SplitMultyply(s string, sum int) int {

	return sum
}