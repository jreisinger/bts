package btsflights

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Print prints a table with flights.
func Print(flights []Flight) {
	const format = "%v\t%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Type", "Number", "Destination", "Date", "Planned", "Current")
	fmt.Fprintf(tw, format, "----", "------", "-----------", "----", "-------", "-------")
	for _, f := range flights {
		fmt.Fprintf(tw, format, f.Type, f.Number, f.Destination, f.Date.Format("2006-01-02"), f.TimePlanned.Format("15:04"), f.TimeCurrent.Format("15:04"))
	}
	tw.Flush()
}
