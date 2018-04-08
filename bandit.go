package main

import (
	"math"
)

func regret (selection []int) int {
	return sum(selection)
}

func epsilonGreedy (arms []int, epsilon float64, selectCnt int) {
	// initialize
	armsCnt := len(arms)
	exploreCnt := int(math.Ceil(epsilon * float64(selectCnt)))

	// exploration
	for i := 0; i < exploreCnt; i++ {
		slot(arms[i % armsCnt])
	}

	// exploitation
}

func sum (slice []int) int {
	s := 0
	for _, v := range slice {
		s += v
	}
	return s
}

func mean (slice []int) float64 {
	return float64(sum(slice)) / float64(len(slice))
}
