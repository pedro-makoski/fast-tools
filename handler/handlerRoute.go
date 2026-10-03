package handler

import "os"

func Handle() {
	instructions := getInstructions()
	instructions.callComands()
}

func getInstructions() Instructions {
	argumentos := os.Args[1:]

	return Instructions{
		path: argumentos,
	}
}
