package flights

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/mapu77/AD-Labs/6-go-api/database"
)

type Flight struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code          string        `json:"code"`
	Company       string        `json:"company"`
	DepartureTime string        `json:"departure_time" bson:"departure_time"`
	DepartureCity string        `json:"departure_city" bson:"departure_city"`
	ArrivalTime   string        `json:"arrival_time" bson:"arrival_time"`
	ArrivalCity   string        `json:"arrival_city" bson:"arrival_city"`
}

// Make a Flight persistent and return its unique identifier as a string.
// If any error occurs, returns error code in err return value. It is nil otherwise.
func Persist(f *Flight) (string, error) {
	session, dbName, err := database.GetMongoDBSession()
	defer session.Close()

	c := session.DB(dbName).C("flights")
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

// Returns the flights matching the parameters, if any. Errors are returned in error second return value.
func ListBy(code string, company string, departureCity string, arrivalCity string) (*[]Flight, error) {
	session, dbName, err := database.GetMongoDBSession()
	defer session.Close()

	c := session.DB(dbName).C("flights")
	var args = bson.M{}
	if code != "" {
		args["code"] = code
	}
	if company != "" {
		args["company"] = company
	}
	if departureCity != "" {
		args["departure_city"] = departureCity
	}
	if arrivalCity != "" {
		args["arrival_city"] = arrivalCity
	}
	var flights []Flight
	err = c.Find(args).Sort("_id").All(&flights)
	return &flights, err
}
