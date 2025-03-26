package monad

type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{Value: value}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func (r Result[T]) Bind(fn func(T) Result[T]) Result[T] {
	if r.Err != nil {
		return r
	}
	return fn(r.Value)
}
