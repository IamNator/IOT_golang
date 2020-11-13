package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InsertHandler(res http.ResponseWriter, req *http.Request) {

	var info Customer
	json.NewDecoder(req.Body).Decode(&info)
	fmt.Fprintf(res, " data inserted : %s", info)

}

func FetchHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req) //creates a Map values passed in url
	id := vars["id"]      //Extracted from requested url

	var jsonData Customer

	json.NewEncoder("data.json").Encode(jsonData)

}
