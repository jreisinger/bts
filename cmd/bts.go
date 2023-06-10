package main

import (
	"flag"
	"log"

	"github.com/jreisinger/bts"
)

var flightsType = bts.FlightTypeFlag("type", bts.Arrival, "flights type")

func main() {
	flag.Parse()

	flights, err := bts.GetFlights(*flightsType)
	if err != nil {
		log.Fatal(err)
	}

	flights.Print()
}
