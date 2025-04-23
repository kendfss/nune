package nune

// prod returns the product of the elements of the slice.
func prod[T Number](s []T) T {
	var prod T = 1
	for i := 0; i < len(s); i++ {
		prod *= s[i]
	}
	return prod
}

func clone[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}
