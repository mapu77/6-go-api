package flights

import (
	"net/http"
	"encoding/json"
	"log"
)

func NewFlight(w http.ResponseWriter, r *http.Request) {
	var f Flight
	var err error
	err = json.NewDecoder(r.Body).Decode(&f)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	var id = make(map[string]string)
	id["id"], err = Persist(&f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(id)
	w.WriteHeader(http.StatusCreated)
}

func ListFlights(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var code string
	if len(params["code"]) > 0 {
		code = params["code"][0]
	}
	var company string
	if len(params["company"]) > 0 {
		company = params["company"][0]
	}
	var departureCity string
	if len(params["departureCity"]) > 0 {
		departureCity = params["departureCity"][0]
	}
	var arrivalCity string
	if len(params["arrivalCity"]) > 0 {
		arrivalCity = params["arrivalCity"][0]
	}
	var flights *[]Flight
	var err error
	flights, err = ListBy(code, company, departureCity, arrivalCity)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(flights)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func ListFlightDestinations(w http.ResponseWriter, r *http.Request) {
	var destinations *[]string
	var err error
	destinations, err = ListDestinations()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(destinations)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func ListFlightCompanies(w http.ResponseWriter, r *http.Request) {
	var companies *[]string
	var err error
	companies, err = ListCompanies()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(companies)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
