package controllers

import (
	"encoding/json"
	"github.com/IamNator/IOT_golang/models"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func InsertHandler(res http.ResponseWriter, req *http.Request) {

	// file, _ := os.Open("data.json")
	// defer file.Close()

	var jsonData models.Customer

	json.NewDecoder(req.Body).Decode(&jsonData)
	jsonData.ID = "45"
	//Insert to MongoDb
	//json.NewEncoder(res).Encode(&jsonData)

}

func FetchHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req) //creates a Map values passed in url
	id := vars["id"]      //Extracted from requested url

	var jsonData models.Customer

	file, _ := os.Open("data.json")
	defer file.Close()

	json.NewDecoder(file).Decode(&jsonData)
	jsonData.ID = id

	json.NewEncoder(res).Encode(&jsonData)

}
