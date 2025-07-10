package lib

func MapSlice[T any, R any](input []T, f func(T) R) []R {
    output := make([]R, len(input))
    for i, v := range input {
        output[i] = f(v)
    }
    return output
}