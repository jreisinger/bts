package main

import (
	"flag"
	"log"

	"github.com/jreisinger/btsflights"
)

var flightsType = btsflights.TypeFlag("type", btsflights.Arrival, "flights type")

func main() {
	flag.Parse()

	flights, err := btsflights.Get(*flightsType)
	if err != nil {
		log.Fatal(err)
	}

	btsflights.Print(flights)
}
