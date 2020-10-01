package main

import "fmt"

func main() {
	var iteraciones int
	var euler float64
	var aux float64

	fmt.Scan(&iteraciones)

	for i := 0; i < iteraciones+1; i++ {
		aux = 1
		for j := i; j >= 1; j-- {
			aux = aux * float64(j)
		}
		euler = euler + (1 / aux)
	}
	fmt.Print(euler)
}
