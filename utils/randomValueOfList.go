package utils

import "math/rand"

func GetRandomValueOfList[T any](listOfItems []T) T {
	randomPos := rand.Intn(len(listOfItems))

	return listOfItems[randomPos]
}
