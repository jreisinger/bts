Btsflights is a package and a command that scrapes data about flights from Bratislava Airport [website](https://www.bts.aero/).

To install and use the command:

```
$ go install cmd/btsflights.go
$ btsflights
Type     Number  Destination         Date        Planned  Current
----     ------  -----------         ----        -------  -------
Arrival  FT3127  HURGHADA            2023-06-09  20:50    00:00
Arrival  6D6451  MONASTIR            2023-06-09  23:00    23:35
Arrival  FR8226  ALGHERO             2023-06-09  23:25    23:15
<...>
$ btsflights -type departure
Type       Number  Destination         Date        Planned  Current
----       ------  -----------         ----        -------  -------
Departure  FT3128  HURGHADA            2023-06-09  21:50    21:50
Departure  FR6388  KAUNAS              2023-06-09  22:05    22:12
Departure  6D6156  ANTALYA             2023-06-10  03:30    03:30
<...>
```
