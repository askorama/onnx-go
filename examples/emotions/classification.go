package main

import (
	"math"
	"sort"
)

func softmax(input []float32) []float32 {
	var sumExp float64
	output := make([]float32, len(input))
	for i := 0; i < len(input); i++ {
		sumExp += math.Exp(float64(input[i]))
	}
	for i := 0; i < len(input); i++ {
		output[i] = float32(math.Exp(float64(input[i]))) / float32(sumExp)
	}
	return output
}

func classify(input []float32) emotions {
	result := make(emotions, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = emotion{
			emotion: emotionTable[i],
			weight:  input[i],
		}
	}
	sort.Sort(sort.Reverse(result))
	return result
}
