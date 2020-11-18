package models

import (
	"gopkg.in/mgo.v2/bson"
)

type UserDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	GPS       string `json:"gps"`
}

type Meter struct {
	//	meterID                 string
	Lastupdate              string `json:"last_update"`
	CurrentPowerConsumption string `json:"current_power_consumption"`
	TotalEnegyConsumed      string `json:"total_energy_consumed"` //Total consumed in last 30 days
}

type Device struct {
	Name                    string `json:"name"`                      //Fan, TV
	Location                string `json:"location"`                  // kitchen, bathroom
	State                   string `json:"state"`                     // on/off
	CurrentPowerConsumption string `json:"current_power_consumption"` //eg 45watts
}

/*****************Main struct **************/
type Customer struct {
	ID          bson.ObjectId `json: "id"` //data type will be updated before production
	UserDetails *UserDetails  `json:"user_details"`
	Meter       *Meter        `json:"meter"`
	Devices     []Device      `json:"devices"`
}

/**************For Keeping Track of id*******************/
type Map_id struct {
	ID        bson.ObjectId `json: "id"`
	FirstName string        `json:"first_name"`
}
