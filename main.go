package main

/*
This api accepts data from an Energy meter and stores it to a mysql database
*/

import (
	"fmt"
	"github.com/IamNator/IOT_golang/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/insert", controllers.InsertHandler).Methods("POST")
	router.HandleFunc("/api/fetch/{id}", controllers.FetchHandler).Methods("GET")

	fmt.Println("Server started on part 9080")
	log.Fatal(http.ListenAndServe(":9080", router))

}
