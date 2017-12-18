package hotels

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/mapu77/AD-Labs/6-go-api/database"
	"log"
)

type Hotel struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name"`
	HotelChain string        `json:"hotel_chain" bson:"hotel_chain"`
	Rooms      int           `json:"rooms"`
	Street     string        `json:"street"`
	ZipCode    string        `json:"zip_code" bson:"zip_code"`
	City       string        `json:"city"`
	Country    string        `json:"country"`
}

// Make an Hotel persistent and return its unique identifier as a string.
// If any error occurs, returns error code in err return value. It is nil otherwise.
func Persist(f *Hotel) (string, error) {
	session, dbName, err := database.GetMongoDBSession()
	defer session.Close()

	c := session.DB(dbName).C("hotels")
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

// Returns the hotels matching the parameters, if any. Errors are returned in error second return value.
func ListBy(name string, hotelChain string, city string) ([]Hotel, error) {
	session, dbName, err := database.GetMongoDBSession()
	defer session.Close()

	c := session.DB(dbName).C("hotels")
	var args = bson.M{}
	if name != "" {
		args["name"] = name
	}
	if hotelChain != "" {
		args["hotel_chain"] = hotelChain
	}
	if city != "" {
		args["city"] = city
	}
	var hotels []Hotel
	err = c.Find(args).Sort("_id").All(&hotels)
	return hotels, err
}
