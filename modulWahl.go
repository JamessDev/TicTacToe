package main

import "fmt"

func ModiWahl() int {
	var auswahl int
	var valid bool
	for !valid {
		fmt.Println("Möchtest du Lokal(1), gegen einen zufälligen Bot(2) spielen, gegen einen medium Bot (3), oder einen schwierigen Bot(4) spielen? ")

		fmt.Scan(&auswahl)

		if auswahl == 1 || auswahl == 2 || auswahl == 3 || auswahl == 4 {
			valid = true
		}
	}
	return auswahl
}
