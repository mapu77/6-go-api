package flights

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
)

type Flight struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code          string        `json:"code"`
	Company       string        `json:"company"`
	DepartureTime string        `json:"departure_time"`
	DepartureCity string        `json:"departure_city"`
	ArrivalTime   string        `json:"arrival_time"`
	ArrivalCity   string        `json:"arrival_city"`
}

// Make a Flight persistent and return its unique identifier as a string.
// If any error occurs, returns error code in err return value. It is nil otherwise.
func Persist(f *Flight) (string, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("ad-travel-agency").C("flights")
	err = c.Insert(&f)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]bson.ObjectId
	err = c.Find(nil).Select(bson.M{"_id": 1}).Sort("-_id").Limit(1).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result["_id"].Hex(), err
}

func ListBy(code string, company string, departureCity string, arrivalCity string) ([]Flight, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("ad-travel-agency").C("flights")
	var args = bson.M{}
	if code != "" {
		args["code"] = code
	}
	if company != "" {
		args["company"] = company
	}
	if company != "" {
		args["departureCity"] = departureCity
	}
	if company != "" {
		args["arrivalCity"] = arrivalCity
	}
	var flights []Flight
	err = c.Find(args).Sort("_id").All(&flights)
	return flights, err
}
