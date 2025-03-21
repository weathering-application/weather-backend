package monad

type IO[T any] struct {
	Run func() (T, error)
}

func Pure[T any](value T) IO[T] {
	return IO[T]{Run: func() (T, error) { return value, nil }}
}

// Map applies a pure function to the result of the computation
func (io IO[T]) Map(f func(T) T) IO[T] {
	return IO[T]{Run: func() (T, error) {
		result, err := io.Run()
		if err != nil {
			return result, err
		}
		return f(result), nil
	}}
}

// FlatMap (aka Bind) chains computations, ensuring errors are propagated
func (io IO[T]) FlatMap(f func(T) IO[T]) IO[T] {
	return IO[T]{Run: func() (T, error) {
		result, err := io.Run()
		if err != nil {
			return result, err
		}
		return f(result).Run()
	}}
}
