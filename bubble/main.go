package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := generateSlice(59)
	fmt.Println("our list of numbers is: ", numbers)
	bubbleSort(numbers)

	fmt.Println("after sort: ",  numbers)

}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99) - rand.Intn(99)
	}
	return slice
}

func bubbleSort(numbers []int) {
	var N int = len(numbers)
	for i := 0; i < N; i++ {
		fmt.Println("sorting: ", numbers)
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, prevPasses int) bool {
	var (
	N int = len(numbers)
	firstIndex int = 0
	secondIndex int = 1
	didSwap bool = false
	)

	for secondIndex < (N - prevPasses) {
		var (
		firstNumber int = numbers[firstIndex]
		secondtNumber int = numbers[secondIndex]
		)

		if firstNumber > secondtNumber {
			 numbers[firstIndex] = secondtNumber
			 numbers[secondIndex] = firstNumber
			 didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}
