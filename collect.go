package collect

// ToMapOfBool collects elements of a slice into a map of bool. Slice
// must consist of comparable types, which are valid map keys.
func ToMapOfBool[T comparable](s []T) (m map[T]bool) {
	m = make(map[T]bool, len(s))

	for _, v := range s {
		m[v] = true
	}

	return
}

// ToMapOfBoolFunc collects elements of a slice into a map of bool.
// The second argument is a key function which converts an element into a
// comparable key.
func ToMapOfBoolFunc[T any, F comparable](s []T, f func(T) F) (m map[F]bool) {
	m = make(map[F]bool, len(s))

	for _, v := range s {
		m[f(v)] = true
	}

	return
}

// ToMapOfEmptyStruct collects elements of a slice into a map of struct{}. Slice
// must consist of comparable types, which are valid map keys.
func ToMapOfEmptyStruct[T comparable](s []T) (m map[T]struct{}) {
	m = make(map[T]struct{}, len(s))

	for _, v := range s {
		m[v] = struct{}{}
	}

	return
}

// ToMapOfEmptyStructFunc collects elements of a slice into a map of struct{}.
// The second argument is a key function which converts an element into a
// comparable key.
func ToMapOfEmptyStructFunc[T any, F comparable](s []T, f func(T) F) (m map[F]struct{}) {
	m = make(map[F]struct{}, len(s))

	for _, v := range s {
		m[f(v)] = struct{}{}
	}

	return
}

// ToMap collects elements of a slice into a maps using a key-value function.
func ToMap[T any, K comparable, V any](s []T, kvFunc func(T) (K, V)) (m map[K]V) {
	m = make(map[K]V, len(s))

	for _, e := range s {
		k, val := kvFunc(e)

		m[k] = val
	}

	return
}
