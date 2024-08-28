package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(div(2, 0))
	fmt.Println(safeDiv(2, 0))
	fmt.Println(safeDiv(7, 2))
}

func safeDiv(a, b int) (q int, err error) {
	// q & err are local variables in safeDiv (just like a & b)
	defer func() {
		if e := recover(); e != nil {
			log.Println("Error: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
