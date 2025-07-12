package lib

func MapMap[T comparable, D any, R any](input map[T]D, f func(D) R) []R {
    output := make([]R, len(input))
	i := 0
    for _, v := range input {
        output[i] = f(v)
		i++
    }
    return output
}

func FindKeyByValue[T comparable, D any](m map[T]D, isEq func(D) bool) (T, bool) {
    for k, v := range m {
        if isEq(v) {
            return k, true
        }
    }
	var zero T
    return zero, false
}