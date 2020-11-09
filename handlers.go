package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type data struct {
	id         string    `json:"id`          //customer ID
	lastUpdate time.Time `json:"lastUpdate"` //time of update
	energy     float64   `json:"energy"`     //amount of Energy consumed
}

func InsertHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Fprintln(res, "data shows", id)
}

func FetchHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id := vars["id"]
	info := data{id: id, lastUpdate: time.Now(), energy: 234.45}
	json.NewEncoder(res).Encode(info)

}

// DBNAME := "mysql"
// DBPASSWORD := "root:password1@tcp(127.0.0.1:3306)/test" //127.0.0.1:3306 == localhost:3306

// fmt.Println("Go MySQL Tutorial")

// // Open up our database connection.
// // The database is called testDb
// db, err := sql.Open(DBNAME, DBPASSWORD)

// // if there is an error opening the connection, handle it
// if err != nil {
// 	panic(err.Error())
// }

// // defer the close till after the main function has finished
// // executing
// defer db.Close()

// // perform a db.Query insert
// insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

// // if there is an error inserting, handle it
// if err != nil {
// 	panic(err.Error())
// }
// // be careful deferring Queries if you are using transactions
// defer insert.Close()
