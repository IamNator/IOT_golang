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
type UserController struct { //data type associated with desired methods
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController { //Embeds the object in a function that returns address of struct...(or a datatype with the desired methods)
	return &UserController{s}
}

func GetSession() *mgo.Session { //Returns the address of the newly created object
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
	uc.session.DB("iot-golang").C("data").Insert(jsonData) //Inserts the data in MongoDb
	id_map := models.Map_id{
		jsonData.UserDetails.FirstName,
		string(jsonData.ID),
	}

	uc.session.DB("iot-golang").C("id-map").Insert(id_map) //Inserts the id-firstname map into MongoDb

	//response back to client
	jsn_data, _ := json.Marshal(jsonData)
	id_mp_jsn, _ := json.Marshal(id_map)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(res, "%s\n%s", jsn_data, id_mp_jsn)

}

func (uc UserController) FetchHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var jsonData models.Customer
	//var := mux.Vars(req)
	//id := var["id"]

	if err := uc.session.DB("iot-golang").C("data").FindId(id).One(&jsonData); err != nil {
		res.WriteHeader(404) //Page not found
		return
	}

	jsn_data, _ := json.Marshal(jsonData)
	res.WriteHeader("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fwt.Fprintf(res, "%s", jsn_data)

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
