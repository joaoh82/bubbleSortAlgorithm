package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine", "dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"}
	fmt.Println("Original list of words:", words)
	bubbleSort(words)
	fmt.Println("After sorting:", words)
}

func bubbleSort(words []string) {
	N := len(words)
	for i := 0; i < N; i++ {
		fmt.Println("Doing a sweep:", words)
		if !sweep(words, i) {
			return
		}
	}
}

func sweep(words []string, prevPasses int) bool {
	N := len(words)
	var firstIndex int
	secondIndex := 1
	didSwap := false

	for secondIndex < (N - prevPasses) {
		firstWord := words[firstIndex]
		secondWord := words[secondIndex]

		if strings.Compare(firstWord, secondWord) > 0 {
			words[firstIndex] = secondWord
			words[secondIndex] = firstWord
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}
