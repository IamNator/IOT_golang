package main

//This api accepts data from an Energy meter and stores it to a mysql database
//Energy amount, time, and Customer Id

import (
	//	"database/sql"
	//"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/insert", InsertHandler).Methods("POST")
	router.HandleFunc("/api/fetch/{id}", FetchHandler).Methods("GET")
	fmt.Println("Server started on part 9080")

	log.Fatal(http.ListenAndServe(":9080", router))

}
