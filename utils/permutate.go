package utils

import (
	"math/rand"
)

func Permutate[T any](listOfValues []T) []T {
	listPermutated := make([]T, len(listOfValues))

	copy(listPermutated, listOfValues)

	for i := range len(listOfValues) {
		randomPos := rand.Intn(len(listOfValues))

		valueA := listPermutated[i]
		valueB := listPermutated[randomPos]

		listPermutated[i] = valueB
		listPermutated[randomPos] = valueA
	}

	return listPermutated
}
