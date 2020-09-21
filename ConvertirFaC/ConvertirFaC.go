package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64
	var celsius float64
	fmt.Print("Ingrese los grados fahrenheit: ")
	fmt.Scan(&fahrenheit)
	celsius = ((fahrenheit - 32.0) * (5.0 / 9.0))
	fmt.Println("Los grados en Celsius son: ", celsius)
}
