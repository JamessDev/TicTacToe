package main

import (
	"fmt"
)

func main() {
	var test Game


	test = Game{0, 0, 0, 0, 0, 0, 0, 0, 0}

	var fieldNumber int

	user := O
	bot := X

	auswahl := ModiWahl()
	Render(test)

	for !IsGameOver(test) {

		fmt.Print("Welches Feld wollen Sie beschriften?: ")
		fmt.Scan(&fieldNumber)

		if fieldNumber >= 9 {
			fmt.Println("Die Zahl ist zu hoch")
			continue
		}

		if test[fieldNumber] != empty {
			fmt.Println("Das Feld ist belegt, bitte wähle ein anderes aus!")
			continue
		}

		test[fieldNumber] = user
		if IsGameOver(test) {
			return
		}

		switch auswahl {
		case (1):
			if user == O {
				user = X
			} else {
				user = O
			}
		case (2):

			fieldNumber := BotRandom(test)
			if fieldNumber == 42 {
				fmt.Println("Error!")
			}
			test[fieldNumber] = bot

		case (3):

			fieldNumber := BotMedium(test)
			if test[fieldNumber] != empty {
				continue
			}
			if fieldNumber == 42 {
				fmt.Println("Error!")
			}
			test[fieldNumber] = bot

		case (4):

			fieldNumber := BotHard(test)
			if test[fieldNumber] != empty {
				continue
			}
			if fieldNumber == 42 {
				fmt.Println("Error!")
			}
			test[fieldNumber] = bot
		}
		Render(test)

	}
}
