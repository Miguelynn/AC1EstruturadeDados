package main

import (
	"fmt"
)

func hanoi(n int, A string, B string, C string) int {
	if n == 1 {
		fmt.Printf("Mova o disco 1 de %s para %s\n", A, C)
		return 1
	}

	movimentos := hanoi(n-1, A, C, B)

	fmt.Printf("Mova o disco %d de %s para %s\n", n, A, C)

	movimentos += hanoi(n-1, B, A, C)

	return movimentos
}

func main() {
	n := 3 
	movimentos := hanoi(n, "A", "B", "C")
	fmt.Printf("Número de movimentos necessários: %d\n", movimentos)
}
