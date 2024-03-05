package tri

// If is a generic function that returns `t` if `condition` is true, otherwise it returns `f`.
func If[T any](condition bool, t T, f T) T {
	if condition {
		return t
	}
	return f
}

// IfExec executes the given function `t` if the `condition` is true,
// otherwise it executes the function `f`. It returns the result of the executed function.
func IfExec[T any](condition bool, t func() T, f func() T) T {
	if condition {
		return t()
	}
	return f()
}

// IfExecWithErr executes `t` and returns its result and error if the `condition` is true,
// otherwise it executes `f` and returns its result and error.
func IfExecWithErr[T any](condition bool, t func() (T, error), f func() (T, error)) (T, error) {
	if condition {
		return t()
	}
	return f()
}

// IfExecOrVal executes a function or returns a value based on a condition.
// If the condition is true, it executes the provided function `t` and returns its result.
// If the condition is false, it returns the provided value `f`.
func IfExecOrVal[T any](condition bool, t func() T, f T) T {
	if condition {
		return t()
	}
	return f
}

// IfExecOrValWithErr executes `t` and returns its result and a nil error if the `condition` is true,
// otherwise it returns the provided value `f` and a nil error.
func IfExecOrValWithErr[T any](condition bool, t func() (T, error), f T) (T, error) {
	if condition {
		return t()
	}
	return f, nil
}

// IfValOrExec evaluates a condition and returns a value based on the result.
// If the condition is true, it returns the provided value 't'.
// If the condition is false, it executes the provided function 'f' and returns its result.
// The type of the value and the return type of the function must be the same.
func IfValOrExec[T any](condition bool, t T, f func() T) T {
	if condition {
		return t
	}
	return f()
}

// IfValOrExecWithErr returns the provided value `t` and a nil error if the `condition` is true,
// otherwise it executes `f` and returns its result and error.
func IfValOrExecWithErr[T any](condition bool, t T, f func() (T, error)) (T, error) {
	if condition {
		return t, nil
	}
	return f()
}
