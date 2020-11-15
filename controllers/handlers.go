package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/IOT_golang/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
)

//For mongodb Access
type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func GetSession() *mgo.Session {
	//Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	//Check if connection err, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func (uc UserController) InsertHandler(res http.ResponseWriter, req *http.Request) {

	// file, _ := os.Open("data.json")
	// defer file.Close()

	var jsonData models.Customer

	json.NewDecoder(req.Body).Decode(&jsonData)
	jsonData.ID = "45"
	//Insert to MongoDb
	//json.NewEncoder(res).Encode(&jsonData)
	fmt.Fprintln(res, "Data inserted successfully")

}

func (uc UserController) FetchHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req) //creates a Map values passed in url
	id := vars["id"]      //Extracted from requested url

	var jsonData models.Customer

	file, _ := os.Open("controllers/data.json")
	defer file.Close()

	json.NewDecoder(file).Decode(&jsonData)
	jsonData.ID = id

	json.NewEncoder(res).Encode(&jsonData)

}
