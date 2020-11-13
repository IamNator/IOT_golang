package main

import (
	"encoding/json"
	//	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func InsertHandler(res http.ResponseWriter, req *http.Request) {

	// file, _ := os.Open("data.json")
	// defer file.Close()

	var jsonData Customer

	jsonData.ID = "45"
	json.NewDecoder(req.Body).Decode(&jsonData)
	//Insert to MongoDb
	//json.NewEncoder(res).Encode(&jsonData)

}

func FetchHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req) //creates a Map values passed in url
	id := vars["id"]      //Extracted from requested url

	var jsonData Customer

	file, _ := os.Open("data.json")
	defer file.Close()

	jsonData.ID = id
	json.NewDecoder(file).Decode(&jsonData)

	json.NewEncoder(res).Encode(jsonData)

}
