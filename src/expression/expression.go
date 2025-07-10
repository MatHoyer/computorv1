package expression

import (
	"computorv1/src/lib"
	"computorv1/src/number"
	"fmt"
	"strings"
)

type Operation string

// Constantes possibles
const (
    ADD 		Operation = "+"
    MULTIPLY 	Operation = "*"
    DIVIDE 		Operation = "/"
)

type expression struct {
	values 		[]*number.Number
	operation	Operation
}

func Create(values []*number.Number, operationType Operation) *expression {
	if operationType != ADD && operationType != MULTIPLY && operationType != DIVIDE {
		panic("Operation can only be +, * or /")
	}
	return &expression{values, operationType}
}

func Str(exp *expression) string {
	mappedSlice := lib.MapSlice(exp.values, func(n *number.Number) string {
		return number.Str(*n)
	})

	separator := fmt.Sprintf(" %s ", exp.operation)
	return strings.Join(mappedSlice, separator)
}