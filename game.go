package main

import "fmt"

func Render(game Game) {
	fmt.Printf("%s | %s | %s\n", FieldToString(game[0]), FieldToString(game[1]), FieldToString(game[2]))
	fmt.Printf("- + - + -\n")
	fmt.Printf("%s | %s | %s\n", FieldToString(game[3]), FieldToString(game[4]), FieldToString(game[5]))
	fmt.Printf("- + - + -\n")
	fmt.Printf("%s | %s | %s\n", FieldToString(game[6]), FieldToString(game[7]), FieldToString(game[8]))
}

func IsGameOver(game Game) bool {

	winner := empty

	for i := 0; i < 3; i++ {

		winner = WonH(game, i*3)
		if winner == empty {
			winner = WonV(game, i%3)
			if winner == empty {
				winner = WonD(game)
			}

		}

		if winner != empty {
			if winner == X {
				fmt.Println("Dein Gegner hat Gewonnen!")
				return true
			}
			if winner == O {
				fmt.Println("Glückwunsch, du hast Gewonnen!")
				return true
			}
			return false
		}
	}

	if !Contains(game, 0) {
		fmt.Println("Keiner hat gewonnen!")
		return true
	}

	return false
}

func Contains(s Game, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
