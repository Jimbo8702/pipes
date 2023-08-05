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

func withParam[T any](operation func(T, T) T, param T) func(T) T {
	return func(n T) T {
		return operation(n, param)
	}
}

