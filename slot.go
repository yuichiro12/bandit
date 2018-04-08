package main

import (
	"math/rand"
	"time"
)

func init () {
	rand.Seed(time.Now().UnixNano())
}

func slot (n int) int {
	return rand.Intn(n)
}
