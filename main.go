package main

import (
	"fmt"
)

func main () {
	arms := []int{10, 20, 30, 40, 50}
	for i, v := range arms {
		fmt.Println("slot", i+1, ":", slot(v))
	}
}
