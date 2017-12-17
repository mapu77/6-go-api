package hotels

import (
	"net/http"
	"log"
	"encoding/json"
)

func NewHotel(w http.ResponseWriter, r *http.Request) {
	var f Hotel
	var err error
	err = json.NewDecoder(r.Body).Decode(&f)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var id = make(map[string]string)
	id["id"], err = Persist(&f)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(id)
	w.WriteHeader(201)
}

func ListHotels(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var name string
	if len(params["name"]) > 0 {
		name = params["name"][0]
	}
	var hotelChain string
	if len(params["hotelChain"]) > 0 {
		hotelChain = params["hotelChain"][0]
	}
	var city string
	if len(params["city"]) > 0 {
		city = params["city"][0]
	}
	var flights []Hotel
	var err error
	flights, err = ListBy(name, hotelChain, city)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(flights)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	} else {
		w.WriteHeader(200)
	}
}

func ListHotelChains(w http.ResponseWriter, r *http.Request) {
	var hotelChains *[]string
	var err error
	hotelChains, err = ListChains()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(hotelChains)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	} else {
		w.WriteHeader(200)
	}
}

func ListHotelCities(w http.ResponseWriter, r *http.Request) {
	var cities *[]string
	var err error
	cities, err = ListCities()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(cities)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	} else {
		w.WriteHeader(200)
	}
}
