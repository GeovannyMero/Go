package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {

}
func main() {
	router := mux.NewRouter()
	//endpoints
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
}
