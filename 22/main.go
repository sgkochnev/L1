package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

func main() {
	//если результат опрераций укладывается в диапазон от -2^63 до 2^63-1, то
	//можно использовать стандартные операции + - * /
	{
		var (
			a int64 = 1 << 21
			b int64 = 1 << 22
		)
		fmt.Println(b + a)
		fmt.Println(a - b)
		fmt.Println(b * a)
		fmt.Println(b / a)
	}

	// используем длинную арифметику
	{
		var (
			a   = big.NewInt(1<<63 - 1)
			b   = big.NewInt(1<<63 - 2)
			res = big.NewInt(0)
		)
		fmt.Println(res.Add(a, b))
		fmt.Println(res.Sub(b.Neg(b), a))
		fmt.Println(res.Mul(a, b))
		fmt.Println(res.Div(b, a))

	}

}
