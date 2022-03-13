package xutil

func ForEach[T any](a []T, f func(T) bool) {
	for _, e := range a {
		if !f(e) {
			break
		}
	}
}

func Filter[T any](a []T, f func(T) bool) []T {
	var n []T
	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
