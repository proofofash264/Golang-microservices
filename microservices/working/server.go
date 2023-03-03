package main

import (
	route "example/user/working/api/customer/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){

	mainRouter := mux.NewRouter()
	// customerRouter := mux.NewRouter()
	// mainRouter.Handle("/", customerRouter)

	mainRouter.HandleFunc("/hello", route.Hello).Methods("GET")
	mainRouter.HandleFunc("/bye", route.Bye).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", mainRouter))
}
