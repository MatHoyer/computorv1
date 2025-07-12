package number

import (
	"fmt"
)

type Number struct {
    Value 	float32
    Degree  int
	IsRoot	bool
}

func Create(value float32, degree int) *Number {
	return &Number{
		Value: value,
		Degree: degree,
		IsRoot: false,
	}
}

func CreateRoot(value float32, degree int) *Number {
	return &Number{
		Value: value,
		Degree: degree,
		IsRoot: true,
	}
}

func Add(num *Number, other *Number) {
	if num.Degree != other.Degree {
		panic("Can't add number with different degrees")
	}
	num.Value += other.Value
}

func Multiply(num *Number, other *Number) {
	num.Value *= other.Value
	num.Degree += other.Degree
}

func Divide(num *Number, other *Number) {
	if other.Value == 0 {
		panic("Can't divide by 0")
	}
	num.Value /= other.Value
	num.Degree -= other.Degree
}

func Eq(num Number, other Number) bool {
	if num.Value != other.Value {
		return false
	}
	if num.Degree != other.Degree {
		return false
	}
	if num.IsRoot != other.IsRoot {
		return false
	}
	return true
}

func Str(num Number) string {
	if num.Value == 0 {
		return "0"
	}

	prefix := ""
	xDegree := ""

	if num.Degree == 0 || (num.Value != 1 && num.Value != -1) {
		prefix = fmt.Sprintf("%g", num.Value)
	} else if num.Value == -1 {
		prefix = "-"
	}

	if num.Degree == 1 {
		xDegree = "X"
	} else if num.Degree != 0 {
		xDegree = fmt.Sprintf("X^%d", num.Degree)
	}

	return prefix + xDegree
}

func Oposite(num *Number) {
	num.Value = -num.Value
}