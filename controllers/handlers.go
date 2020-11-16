package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IamNator/IOT_golang/models"
	_ "github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	s, err := mgo.Dial("mongodb://localhost") //listens on port 27017

	//Check if connection err, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func (uc UserController) InsertHandler(res http.ResponseWriter, req *http.Request) {

	// file, _ := os.Open("data.json")
	// defer file.Close() //Closes file

	var jsonData models.Customer
	json.NewDecoder(req.Body).Decode(&jsonData)
	jsonData.ID = bson.NewObjectId()
	//Insert to MongoDb
	uc.session.DB("iot-golang").C("data").Insert(jsonData)
	jsn_data, _ := json.Marshal(jsonData)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(res, "%s", jsn_data)

}

func (uc UserController) FetchHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(req) //creates a Map values passed in url
	//id := vars["id"]      //Extracted from requested url

	var jsonData models.Customer

	file, _ := os.Open("controllers/data.json")
	defer file.Close()

	json.NewDecoder(file).Decode(&jsonData)
	//jsonData.ID = id

	json.NewEncoder(res).Encode(&jsonData)

}

func (uc UserController) DeleteHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(req) //creates a Map values passed in url
	//id := vars["id"]      //Extracted from requested url

	var jsonData models.Customer

	file, _ := os.Open("controllers/data.json")
	defer file.Close()

	json.NewDecoder(file).Decode(&jsonData)
	//jsonData.ID = id

	json.NewEncoder(res).Encode(&jsonData)

}
