package main

import "fmt"

func main() {

	runAdder(19, 34, new(TexCal))
	runAdder(18, 87, new(Casio))
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

}

func runAdder(a, b int, calculator Calculator) {

	fmt.Println(calculator.add(a, b))
}
