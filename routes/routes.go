package routes

import (
	"net/http"
	"github.com/mapu77/AD-Labs/6-go-api/core/flights"
	"github.com/mapu77/AD-Labs/6-go-api/core/hotels"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

const baseUrl = "/api"
const apiVersion = "/v1"

type Routes []Route

var routes = Routes{
	Route{
		"HelloWorld",
		"GET",
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://app.swaggerhub.com/apis/mapu77/ad-go-api/1.0.0", 303)
		},
	},
	Route{
		"NewFlight",
		"POST",
		baseUrl + apiVersion + "/flights",
		flights.NewFlight,
	},
	Route{
		"GetFlights",
		"GET",
		baseUrl + apiVersion + "/flights",
		flights.ListFlights,
	},
	Route{
		"GetFlightDestinations",
		"GET",
		baseUrl + apiVersion + "/flights/destinations",
		flights.ListFlightDestinations,
	},
	Route{
		"GetFlightCompanies",
		"GET",
		baseUrl + apiVersion + "/flights/companies",
		flights.ListFlightCompanies,
	},
	Route{
		"NewHotel",
		"POST",
		baseUrl + apiVersion + "/hotels",
		hotels.NewHotel,
	},
	Route{
		"GetHotels",
		"GET",
		baseUrl + apiVersion + "/hotels",
		hotels.ListHotels,
	},
	Route{
		"GetHotelChains",
		"GET",
		baseUrl + apiVersion + "/hotels/chains",
		hotels.ListHotelChains,
	},
	Route{
		"GetHotelCities",
		"GET",
		baseUrl + apiVersion + "/hotels/cities",
		hotels.ListHotelCities,
	},
}
