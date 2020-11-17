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

	uc := controllers.NewUserController(controllers.GetSession()) //Creates a Mongo Db session
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/insert", uc.InsertHandler).Methods("POST")
	router.HandleFunc("/api/fetch/{id}", uc.FetchHandler).Methods("GET")
	router.HandleFunc("/api/fetch/{id}", uc.DeleteHandler).Methods("DELETE")

	fmt.Println("Server started on part 9080")
	log.Fatal(http.ListenAndServe(":9080", router))

}
