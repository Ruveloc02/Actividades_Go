package main

import "fmt"

func main() {

	var n float64
	var i float64
	var e float64
	e = 0
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		var aux float64 = fact(i)
		e = e + (1 / aux)
	}
	fmt.Print(e)
}

func fact(i float64) float64 {
	var factorial float64 = i
	if factorial > 1 {
		factorial = i * fact(i-1)
		return factorial
	} else {
		return 1
	}
}
