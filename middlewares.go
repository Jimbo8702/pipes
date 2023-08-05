package main

func makeStageFunc[T any](operation func(T) T) StageFunc[T] {
	return func(in <-chan T) <-chan T {
		out := make(chan T)
		go func() {
			for n := range in {
				out <- operation(n)
			}
			close(out)
		}()
		return out
	}
}