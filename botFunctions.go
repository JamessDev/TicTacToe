package main

import (
	"math/rand"
)

func BotRandom(game Game) int {
	EmptyField := []int{}
	for i, element := range game {
		if element == empty {
			EmptyField = append(EmptyField, i)
		}
	}
	return EmptyField[rand.Intn(len(EmptyField))]
}

func BotHard(game Game) int {

	if game[4] == empty {
		return 4
	} else {
		return BotMedium(game)
	}

	//return 42      Gucken ob man es braucht, ist unreachable Code!
}

func BotMedium(game Game) int {
	valid := true
	validField := false
	var field int

	for !validField {

		switch valid {

		case game[0] == empty && ((game[3] == game[6] && game[3] != empty) || (game[1] == game[2] && game[1] != empty) || (game[4] == game[8] && game[4] != empty)):
			field = 0

		case game[1] == empty && ((game[0] == game[2] && game[0] != empty) || (game[4] == game[7] && game[4] != empty)):
			field = 1
		case game[2] == empty && ((game[0] == game[1] && game[0] != empty) || (game[4] == game[6] && game[4] != empty) || (game[5] == game[8] && game[5] != empty)):
			field = 2
		case game[3] == empty && ((game[0] == game[6] && game[0] != empty) || (game[4] == game[5] && game[4] != empty)):
			field = 3

		case game[4] == empty && ((game[0] == game[8] && game[0] != empty) || (game[1] == game[7] && game[1] != empty) || (game[2] == game[6] && game[2] != empty) || (game[3] == game[5] && game[3] != empty)):
			field = 4
		case game[5] == empty && ((game[2] == game[8] && game[2] != empty) || (game[3] == game[4] && game[3] != empty)):
			field = 5
		case game[6] == empty && ((game[0] == game[3] && game[0] != empty) || (game[2] == game[4] && game[2] != empty) || (game[7] == game[8] && game[7] != empty)):
			field = 6
		case game[7] == empty && ((game[1] == game[4] && game[1] != empty) || (game[6] == game[8] && game[6] != empty)):
			field = 7

		case game[8] == empty && ((game[0] == game[4] && game[0] != empty) || (game[2] == game[5] && game[2] != empty) || (game[6] == game[7] && game[6] != empty)):
			field = 8
		default:

			field = BotRandom(game)

		}

		if game[field] == empty {
			validField = true
			return field
		}
		validField = false

	}

	return 42
}
