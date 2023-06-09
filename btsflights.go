// Package btsflights scrapes data about flights from [Bratislava Airport].
// [Bratislava Airport]: https://www.bts.aero
package btsflights

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

// Type is arrival or departure.
type Type int

const (
	Arrival Type = iota
	Departure
)

func (ft Type) String() string {
	return [...]string{"Arrival", "Departure"}[ft]
}

type Flight struct {
	Type        Type
	Number      string
	Destination string
	Date        time.Time
	TimePlanned time.Time
	TimeCurrent time.Time
}

const (
	ArrivalsURL   = "https://www.bts.aero/en/flights/arrivals-departures/current-arrivals/"
	DeparturesURL = "https://www.bts.aero/en/flights/arrivals-departures/current-departures/"
)

// Get returns the flights of arrival or departure type.
func Get(t Type) ([]Flight, error) {
	switch t {
	case Arrival:
		return parse(ArrivalsURL)
	case Departure:
		return parse(DeparturesURL)
	default:
		return nil, fmt.Errorf("unknown flight type")
	}
}

// parse gets url and parses it for flights data.
func parse(url string) ([]Flight, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	flights := visit(nil, doc)
	return flights, nil
}

// visit traverses an HTML node tree, extracts and returns data about flights.
func visit(flights []Flight, n *html.Node) []Flight {
	if n.Type == html.ElementNode && n.Data == "button" {
		var flight Flight
		for _, a := range n.Attr {
			switch a.Key {
			case "data-flight-type":
				if a.Val == "arrival" {
					flight.Type = Arrival
				} else {
					flight.Type = Departure
				}
			case "data-destination":
				flight.Destination = a.Val
			case "data-flight-number":
				flight.Number = a.Val
			case "data-flight-date":
				t, _ := time.Parse("02. 01. 2006", a.Val)
				flight.Date = t
			case "data-flight-time-planed":
				t, _ := time.Parse("15:04", a.Val)
				flight.TimePlanned = t
			case "data-flight-time-current":
				t, _ := time.Parse("15:04", a.Val)
				flight.TimeCurrent = t
			}
		}
		if !flight.Date.IsZero() {
			flights = append(flights, flight)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		flights = visit(flights, c)
	}
	return flights
}
