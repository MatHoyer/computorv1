package lib

func MapSlice[T any, R any](input []T, f func(T) R) []R {
    output := make([]R, len(input))
    for i, v := range input {
        output[i] = f(v)
    }
    return output
}

func PopFront[T any](slice *[]T) T {
	return Pop(slice, 0)
}

func Pop[T any](slice *[]T, index int) T {
	popedValue := (*slice)[index]
	*slice = append((*slice)[:index], (*slice)[index+1:]...)

	return popedValue
}