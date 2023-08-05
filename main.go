package main

import (
	"fmt"
)

func fib(n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()

	return out
}

func main() {
	//input0
	// nums := []int{2, 3, 4, 7, 1}

	// Create the pipeline
	pipeline := NewPipeline[int]()

	// Add stages to the pipeline
	pipeline.AddStage(makeStageFunc(sq))
	pipeline.AddStage(makeStageFunc(plusTwo))

	// //Start a input stream
	inputStream := fib(100000)

	// //run
	finalChannel := pipeline.Run(inputStream)

	// // Process the final channel
	for n := range finalChannel {
		fmt.Println(n)
	}
}

