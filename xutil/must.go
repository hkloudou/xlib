package xutil

func Mustb[T any](v T, ok bool) T {
	if !ok {
		panic("not ok")
	}
	return v
}

func Mustb2[T1 any, T2 any](v1 T1, v2 T2, ok bool) (T1, T2) {
	if !ok {
		panic("not ok")
	}
	return v1, v2
}

func Mustb3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, ok bool) (T1, T2, T3) {
	if !ok {
		panic("not ok")
	}
	return v1, v2, v3
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return v
}

func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err.Error())
	}
	return v1, v2
}

func Must3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err.Error())
	}
	return v1, v2, v3
}
