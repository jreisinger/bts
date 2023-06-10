package main

import (
	"flag"
	"log"

	"github.com/jreisinger/bts"
)

var flightsType = bts.FlightTypeFlag("type", bts.Both, "arrival or departure")

func main() {
	flag.Parse()

	flights, err := bts.GetFlights(*flightsType)
	if err != nil {
		log.Fatal(err)
	}

	flights.Print()
}
