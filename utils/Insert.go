package utils

func Insert[T any](slice []T, elem T, position int) []T {

	if position > len(slice) {
		return append(slice, elem)
	}

	tmp := append(slice[:position], elem)
	return append(tmp, slice[(position+1):]...)
}
