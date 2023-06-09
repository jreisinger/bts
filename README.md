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
Arrival  6D6113  ANTALYA             2023-06-09  23:35    23:45
Arrival  FR1744  CORFU               2023-06-10  00:15    00:00
```