package database

import (
	"gopkg.in/mgo.v2"
	"os"
	"log"
)

// Returns a MongoDB session. Errors, if any, are returned in error variable.
func GetMongoDBSession() (*mgo.Session, error) {

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
	return session, err
}
