package utils

func GetKeys[K comparable, V any](mapToGet map[K]V) []K {
	keys := make([]K, 0, len(mapToGet))
	for k := range mapToGet {
		keys = append(keys, k)
	}

	return keys
}

func GetValues[K comparable, V any](mapToGet map[K]V) []V {
	values := make([]V, 0, len(mapToGet))
	for _, v := range mapToGet {
		values = append(values, v)
	}

	return values
}
