package utils

func Revert[T any](listValues []T) []T {
	newList := make([]T, len(listValues))

	for index, value := range listValues {
		newList[len(listValues)-index-1] = value
	}

	return newList
}
