package main

import (
	"fmt"
)

func main() {

	var dia int
	var mes int
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	if (mes == 1 && dia >= 21 && dia <= 31) || (mes == 2 && dia >= 1 && dia <= 19) {
		fmt.Print("Acuario")
	} else if (mes == 2 && dia >= 20 && dia <= 28) || (mes == 3 && dia >= 1 && dia <= 20) {
		fmt.Print("Piscis")
	} else if (mes == 3 && dia >= 21 && dia <= 31) || (mes == 4 && dia >= 1 && dia <= 20) {
		fmt.Print("Aries")
	} else if (mes == 4 && dia >= 21 && dia <= 30) || (mes == 5 && dia >= 1 && dia <= 21) {
		fmt.Print("Tauro")
	} else if (mes == 5 && dia >= 22 && dia <= 31) || (mes == 6 && dia >= 1 && dia <= 21) {
		fmt.Print("Geminis")
	} else if (mes == 6 && dia >= 22 && dia <= 30) || (mes == 7 && dia >= 1 && dia <= 23) {
		fmt.Print("Cancer")
	} else if (mes == 7 && dia >= 24 && dia <= 31) || (mes == 8 && dia >= 1 && dia <= 23) {
		fmt.Print("Leo")
	} else if (mes == 8 && dia >= 24 && dia <= 31) || (mes == 9 && dia >= 1 && dia <= 23) {
		fmt.Print("Virgo")
	} else if (mes == 9 && dia >= 24 && dia <= 30) || (mes == 10 && dia >= 1 && dia <= 23) {
		fmt.Print("Libra")
	} else if (mes == 10 && dia >= 24 && dia <= 31) || (mes == 11 && dia >= 1 && dia <= 22) {
		fmt.Print("Escorpio")
	} else if (mes == 11 && dia >= 23 && dia <= 31) || (mes == 12 && dia >= 1 && dia <= 21) {
		fmt.Print("Sagitario")
	} else if (mes == 12 && dia >= 22 && dia <= 31) || (mes == 1 && dia >= 1 && dia <= 20) {
		fmt.Print("Capricornio")
	}
}
