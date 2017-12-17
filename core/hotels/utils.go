package hotels

import (
	"github.com/mapu77/AD-Labs/6-go-api/database"
	"log"
)

func ListChains() (*[]string, error) {
	session, err := database.GetMongoDBSession()
	defer session.Close()

	c := session.DB("ad-travel-agency").C("hotels")
	var hotelChains []string
	err = c.Find(nil).Distinct("hotel_chain", &hotelChains)
	if err != nil {
		log.Fatal(err)
	}
	return &hotelChains, err
}

func ListCities() (*[]string, error) {
	session, err := database.GetMongoDBSession()
	defer session.Close()

	c := session.DB("ad-travel-agency").C("hotels")
	var cities []string
	err = c.Find(nil).Distinct("city", &cities)
	if err != nil {
		log.Fatal(err)
	}
	return &cities, err
}
