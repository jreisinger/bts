Btsflights is a Go package and a command that scrapes data about flights from Bratislava Airport [website](https://www.bts.aero/).

To install and use the command:

```
$ go install cmd/btsflights.go
$ btsflights -type departure
Type       Number  Destination  Date        Planned  Current  Airline     Airplane
----       ------  -----------  ----        -------  -------  -------     --------
Departure  FT3128  HURGHADA     2023-06-09  21:50    21:50    FlyEgypt    Boeing 737-800 (winglets)
Departure  6D6156  ANTALYA      2023-06-10  03:30    03:30    Smartwings  Boeing 737-800 (winglets)
Departure  QS1506  HERAKLION    2023-06-10  04:10    04:10    Smartwings
<...>
```
