package main

const (
	empty = 0
	X     = 1
	O     = 2
)

type Game [9]int

func FieldToString(field int) string {

	if field == empty {
		return " "
	}

	if field == X {
		return "X"
	}

	if field == O {
		return "O"
	}
	return "FALSCH"
}
