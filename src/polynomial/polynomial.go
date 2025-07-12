package polynomial

import (
	"computorv1/src/expression"
	"computorv1/src/lib"
	"computorv1/src/number"
	"fmt"
)

type Polynomial struct {
	Left 	*expression.Expression
	Right 	*expression.Expression
}

func Create(left, right *expression.Expression) *Polynomial {
	return &Polynomial{
		Left: left, 
		Right: right,
	}
}

func Str(pol Polynomial) string {
	return expression.Str(*pol.Left) + " = " + expression.Str(*pol.Right)
}

func Simplify(pol *Polynomial) {
	isFirstModify := true
	for expression.Simplify(pol.Left) {
		if isFirstModify {
			fmt.Println(" Left:")
			isFirstModify = false
		}
		fmt.Println("  •", expression.Str(*pol.Left))
	}
	isFirstModify = true
	for expression.Simplify(pol.Right) {
		if isFirstModify {
			fmt.Println(" Right:")
			isFirstModify = false
		}
		fmt.Println("  •", expression.Str(*pol.Right))
	}
}

func Regroup(pol *Polynomial) {
	rightValues := expression.GetValuesAsSlice(*pol.Right)

	for _, v := range(rightValues) {
		key, ok := lib.FindKeyByValue(pol.Right.Values, func(other *number.Number) bool {
			return v == other
		})
		if !ok {
			panic("Fatal error in: polynomial.Regroup")
		}
		BasicTransfer(
			pol.Right,
			pol.Left,
			key,
		)
		fmt.Println(" •", Str(*pol))
	}
	fmt.Println("All on left polynome:", Str(*pol))
}

func BasicTransfer(from *expression.Expression, to *expression.Expression, key int) {
	number.Oposite(from.Values[key])
	expression.Append(to, from.Values[key])
	delete(from.Values, key)
}

func DivideTransfer(from *expression.Expression, to *expression.Expression, key int) {
	expression.Append(to, number.Create(from.Values[key].Value, 0))
	from.Values[key].Value = 1
}