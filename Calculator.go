package main

import (
	"fmt"
	"time"
)

type Calculator interface {
	add(a, b int) int
}

type TexCal struct {
}

func (tc *TexCal) add(a, b int) int {

	return a + b
}

type Casio struct {
}

func (cc *Casio) add(a, b int) int {
	fmt.Println("Sleeping for 3 seconds...")
	time.Sleep(time.Second * 3)
	return a + b
}
