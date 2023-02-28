package main

import "fmt"

// Reduce.Operacoes.Matematicas.EmDiversosNumeros

func main() {
	somaTotal := reduceSoma(10, 10)
	fmt.Println(somaTotal)

	subtracaoTotal := reduceSubtracao(10, 10)
	fmt.Println(subtracaoTotal)

	// multiplicacaoTotal := reduceMultiplicacao(10, 10, 10)
	// fmt.Println(multiplicacaoTotal)
}

// Funcao reduceSoma( pode receber diversos numeros  ) e vai retornar a soma de todos recebidos. 
func reduceSoma(numeros ...int) int{
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

// Funcao reduceSubtracao( pode receber diversos numeros  ) e vai retornar a subtracao de todos recebidos. 
func reduceSubtracao(numeros ...int) int{
	total := 0
	for _, num := range numeros {
		total -= num
	}
	return total
}

// Funcao reduceMultiplicacao( pode receber diversos numeros  ) e vai retornar a multiplicacao de todos recebidos. 
// func reduceMultiplicacao(numeros ...int) int{
// 	total := 0
// 	for _, num := range numeros {
// 		total = total * num
// 	}
// 	return total
// }