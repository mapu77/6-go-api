package database

import (
	"gopkg.in/mgo.v2"
	"os"
	"log"
)

// Returns a MongoDB session. Errors, if any, are returned in error variable.
func GetMongoDBSession() (*mgo.Session, string, error) {

	mongoUri, exists := os.LookupEnv("MONGODB_URI")
	if !exists {
		// Assume api is connecting to a locally mongo db
		mongoUri = "localhost:27017"
	}
	session, err := mgo.Dial(mongoUri)
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)
	dbName, exists := os.LookupEnv("MONGODB_NAME")
	if !exists {
		dbName = "ad-travel-agency"
	}
	return session, dbName, err
}
