package flights

import (
	"github.com/mapu77/AD-Labs/6-go-api/database"
	"log"
)

func ListDestinations() (*[]string, error) {
	session, dbName, err := database.GetMongoDBSession()
	defer session.Close()

	c := session.DB(dbName).C("flights")
	var arrivals []string
	err = c.Find(nil).Distinct("arrival_city", &arrivals)
	if err != nil {
		log.Fatal(err)
	}
	var departures *[]string
	err = c.Find(nil).Distinct("departure_city", &departures)
	if err != nil {
		log.Fatal(err)
	}
	destinations := append(arrivals, *departures...)
	RemoveDuplicates(&destinations)
	return &destinations, err
}

func RemoveDuplicates(slice *[]string) {
	values := make(map[string]bool)
	j := 0
	for i, key := range *slice {
		if !values[key] {
			values[key] = true
			(*slice)[j] = (*slice)[i]
			j++
		}
	}
	*slice = (*slice)[:j]
}

func ListCompanies() (*[]string, error) {
	session, dbName, err := database.GetMongoDBSession()
	defer session.Close()

	c := session.DB(dbName).C("flights")
	var companies []string
	err = c.Find(nil).Distinct("company", &companies)
	if err != nil {
		log.Fatal(err)
	}
	return &companies, err
}
