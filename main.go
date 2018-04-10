package main

func main () {
	arms := []int{10, 20, 30, 40, 50}
	epsilon := 0.3
	selectCnt := 101
	epsilonGreedy(arms, epsilon, selectCnt)
}
