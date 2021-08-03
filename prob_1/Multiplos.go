package main

import "fmt"

func main() {

	//Criando as variáveis
	var (
		soma3ou5 int
		soma3e5  int
		soma5    int
		soma3    int
		i        int
	)

	//Iniciando um laço para calcular os multiplos até 1000
	for i = 1; i < 1000; i++ {

		//coletando o resto de divisão da variável "i" por 3
		a := i % 3

		//coletando o resto de divisão da variável "i" por 5
		b := i % 5

		//Verificando se o resto de divisão da variável "a" ou "b" são iguais a zero!
		//caso sejam iguais a zero isto indica que o número é multiplo do valor de "i"
		if a == 0 || b == 0 {
			//calculando a soma do valor multiplo de 3 ou 5
			soma3ou5 = soma3ou5 + i

			//Verificando se o resto de divisão da variável "a" é igual a zero!
			if a == 0 && b != 0 {
				//calculando a soma do valor multiplo de 3
				soma3 = soma3 + i

				//Verificando se o resto de divisão da variável "b" é igual a zero!
			} else if a != 0 && b == 0 {
				//calculando a soma do valor multiplo de 5
				soma5 = soma5 + i

				//Verificando se o resto de divisão da variável "a" e "b" são iguais a zero!
			} else if a == 0 && b == 0 {
				//calculando a soma do valor multiplo de 3 e 5
				soma3e5 = soma3e5 + i
			}

		}

	}
	//print dos resultados das devidas somas
	fmt.Println("a soma de todos os múltiplos de 3 abaixo de", i, "é", soma3)
	fmt.Println("a soma de todos os múltiplos de 5 abaixo de", i, "é", soma5)
	fmt.Println("a soma de todos os múltiplos de 3 OU 5 abaixo de", i, "é", soma3ou5)
	fmt.Println("a soma de todos os múltiplos de 3 E 5 abaixo de", i, "é", soma3e5)

}
