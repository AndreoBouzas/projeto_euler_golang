package main

import (
	"fmt"
	"strconv"
)

var soma int

func FibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]

	}

	return f[n]

}

func main() {
	for i := 0; i <= 33; i++ {
		fmt.Print(strconv.Itoa(FibonacciLoop(i)) + " ")
		p := FibonacciLoop(i)
		if p%2 == 0 {
			soma = soma + p
		}
	}
	fmt.Println("\n\nA soma dos termos de valor par, considerando os termos na sequência de Fibonacci cujos valores não excedem quatro milhões é:\n", soma)
	fmt.Println("")

}
