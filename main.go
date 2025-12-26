package main

import "fmt"

func main() {

	runAdder(19, 34, new(TexCal))
	runAdder(18, 87, new(Casio))

}

func runAdder(a, b int, calculator Calculator) {

	fmt.Println(calculator.add(a, b))
}
