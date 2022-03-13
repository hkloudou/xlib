package xutil

func Uniq[T comparable](a []T) []T {
	u := make([]T, 0, len(a))
	m := make(map[T]bool)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = true
			u = append(u, v)
		}
	}
	return u
}

func Ternary[T any](s bool, t T, f T) T {
	if s {
		return t
	}
	return f
}

func TernaryOp[T any](s bool, t func() T, f func() T) T {
	if s {
		return t()
	}
	return f()
}

func PointerOf[T any](v T) *T {
	return &v
}
