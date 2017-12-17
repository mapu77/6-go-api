package routes

import (
	"net/http"
	"github.com/mapu77/AD-Labs/6-go-api/core/flights"
	"fmt"
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
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(writer, "HelloWorld!")
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
	//Route{
	//	"NewHotel",
	//	"POST",
	//	baseUrl + apiVersion + "/hotels",
	//	hotels.NewHotel,
	//},
	//Route{
	//	"GetHotels",
	//	"GET",
	//	baseUrl + apiVersion + "/hotels",
	//	hotels.ListHotels,
	//},
}
