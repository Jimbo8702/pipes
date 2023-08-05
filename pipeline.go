package main

type StageFunc[T any] func(<-chan T) <-chan T

type Pipeline[T any] struct {
	stages      []StageFunc[T]
	outputChan  <-chan T
}

func NewPipeline[T any]() *Pipeline[T] {
	return &Pipeline[T]{}
}

func (p *Pipeline[T]) AddStage(stage StageFunc[T]) {
	p.stages = append(p.stages, stage)
}

func (p *Pipeline[T]) RemoveStage(index int) {
	if index >= 0 && index < len(p.stages) {
		p.stages = append(p.stages[:index], p.stages[index+1:]...)
	}
}

func (p *Pipeline[T]) Run(input <-chan T) <-chan T {
	p.outputChan = input

	for _, stage := range p.stages {
		p.outputChan = stage(p.outputChan)
	}

	return p.outputChan
}



// func sliceToChannel(nums []int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for _, n := range nums {
// 			out <- n
// 		}
// 	}()

// 	return out
// }

