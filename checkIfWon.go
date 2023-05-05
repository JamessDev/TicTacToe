package main

func WonD(game Game) int {
	if (game[0] == game[4] && game[4] == game[8]) || (game[2] == game[4] && game[4] == game[6]) {
		return (game[4])
	}

	return empty
}

func WonH(game Game, offset int) int {
	if game[offset] == game[1+offset] && game[offset] == game[offset+2] {
		return game[offset]
	}

	return empty
}

func WonV(game Game, offset int) int {
	if game[offset] == game[3+offset] && game[offset] == game[offset+6] {
		return game[offset]
	}

	return empty
}

// 0 1 2
// 3 4 5
// 6 7 8
