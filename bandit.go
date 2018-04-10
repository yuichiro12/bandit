package main

import (
	"fmt"
	"math"
)

func regret (selection []int) int {
	return sum(selection)
}

func epsilonGreedy (arms []int, epsilon float64, selectCnt int) {
	// initialize
	armsCnt := len(arms)
	armsTest := make([][]int, armsCnt)
	for i := range armsTest {
		armsTest[i] = make([]int, 0)
	}

	// exploration
	exploreCnt := int(math.Ceil(epsilon * float64(selectCnt)))
	for i := 0; i < exploreCnt; i++ {
		currentArm := i % armsCnt
		armsTest[currentArm] = append(armsTest[currentArm], slot(arms[currentArm]))
	}
	armMeans := make([]float64, armsCnt)
	for i, testi := range armsTest {
		armMeans[i] = mean(testi)
	}
	fmt.Println(armMeans)
	bestArm := arms[findMaxIndex(armMeans)]

	// exploitation
	exploitCnt := selectCnt - exploreCnt
	armsExploit := make([]int, exploitCnt)
	for i := 0; i < exploitCnt; i++ {
		armsExploit[i] = slot(bestArm)
	}
	fmt.Println(armsTest)
	fmt.Println(armsExploit)
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

func findMaxIndex (slice []float64) int {
	maxv := 0.0
	maxi := 0
	for i, v := range slice {
		if v > maxv {
			 maxi = i
		}
	}
	return maxi
}
