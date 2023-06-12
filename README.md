Bts is a Go package and a command that scrapes data about flights from Bratislava Airport website https://www.bts.aero.

To install and use the command:

```
$ go install cmd/bts.go
$ bts -busy
Type       Number  Destination        Date        Planned  Current  Airline                  Airplane
----       ------  -----------        ----        -------  -------  -------                  --------
arrival    QS1505  HERAKLION          2023-06-14  11:00    00:00    Smartwings               Boeing 737-800 (winglets)
arrival    FR5041  LEEDS-BRADFORD     2023-06-14  11:10    00:00    Ryanair                  Boeing 737-800 (winglets)
arrival    FR4642  MILAN-BERGAMO      2023-06-14  11:55    00:00    Ryanair                  Boeing 737-800 (winglets)
departure  6D6320  ARAXOS / PATRAS    2023-06-14  11:00    11:00    Smartwings               Boeing 737-800 (winglets)
departure  4M842   KOSICE             2023-06-14  11:10    11:10    Mavi Gök Havacilik A.S.  Boeing 757-300
departure  FR9512  PALMA DE MALLORCA  2023-06-14  11:20    11:20    Ryanair                  Boeing 737-800 (winglets)
departure  6D6164  ANTALYA            2023-06-14  11:50    11:50    Smartwings               Boeing 737-800 (winglets)
```
