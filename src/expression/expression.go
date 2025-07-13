package expression

import (
	"computorv1/src/lib"
	"computorv1/src/number"
	"fmt"
	"maps"
	"sort"
	"strings"
)

type Operation string

// Constantes possibles
const (
    ADD 		Operation = "+"
    MULTIPLY 	Operation = "*"
    DIVIDE 		Operation = "/"
)

type Expression struct {
	Values 		map[int]*number.Number
	AddedValues []*number.Number
	Operation	Operation
}

func Create(operationType Operation) *Expression {
	if operationType != ADD && operationType != MULTIPLY && operationType != DIVIDE {
		panic("Operation can only be +, * or /")
	}
	return &Expression{
		Values: make(map[int]*number.Number),
		AddedValues: make([]*number.Number, 0, 3),
		Operation: operationType,
	}
}

func Str(exp Expression) string {
	values := GetValuesAsSlice(exp)

	var mappedSlice []string = lib.MapSlice(values, func(n *number.Number) string {
		return number.Str(*n)
	})
	var mappedAddedValues []string = lib.MapSlice(exp.AddedValues, func(n *number.Number) string {
		return number.Str(*n)
	})

	mappedSlice = append(mappedSlice, mappedAddedValues...)
	if len(mappedSlice) == 0 {
		return "0"
	}

	separator := fmt.Sprintf(" %s ", exp.Operation)
	return strings.Join(mappedSlice, separator)
}

func GetValuesAsSlice(exp Expression) []*number.Number {
	gettedValues := maps.Values(exp.Values)
	values := make([]*number.Number, 0, 3)
	for v := range gettedValues {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].Degree > values[j].Degree
	})

	return values
}

func GetHightestDegree(exp Expression) int {
	values := GetValuesAsSlice(exp)
	if len(values) == 0 {
		return 0
	}

	return values[0].Degree
}

func Append(exp *Expression, value *number.Number) {
	_, ok := exp.Values[value.Degree]
	if ok {
		exp.AddedValues = append(exp.AddedValues, value)
	} else {
		exp.Values[value.Degree] = value
	}
}

func Simplify(exp *Expression) bool {
	if len(exp.AddedValues) == 0 {
		return false
	}

	var value *number.Number = lib.PopFront(&exp.AddedValues)

	v, ok := exp.Values[value.Degree]
	if ok {
		switch exp.Operation {
		case "+":
			number.Add(v, value)
		case "/":
			number.Divide(v, value)
		case "*":
			number.Multiply(v, value)
		}
		if v.Value == 0 {
			delete(exp.Values, v.Degree)
		}
	} else {
		exp.Values[value.Degree] = value
	}

	return true
}