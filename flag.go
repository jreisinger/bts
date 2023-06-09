package btsflights

import (
	"flag"
	"fmt"
)

type flightTypeFlag struct{ Type }

func (f *flightTypeFlag) Set(s string) error {
	switch s {
	case "arrival":
		f.Type = Arrival
		return nil
	case "departure":
		f.Type = Departure
		return nil
	}
	return fmt.Errorf("use arrival or departure")
}

// TypeFlag defines flights type flag with the specified name, default value,
// and usage, and returns address of the flag variable.
func TypeFlag(name string, value Type, usage string) *Type {
	f := flightTypeFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Type
}
