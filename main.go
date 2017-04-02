package main

import "fmt"

func main() {
	numbers := []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	fmt.Println("Original list of numbers:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)
}

func bubbleSort(numbers []int) {
	N := len(numbers)
	for i := 0; i < N; i++ {
		fmt.Println("Doing a sweep:", numbers)
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, prevPasses int) bool {
	N := len(numbers)
	var firstIndex int
	secondIndex := 1
	didSwap := false

	for secondIndex < (N - prevPasses) {
		firstNumber := numbers[firstIndex]
		secondNumber := numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}
