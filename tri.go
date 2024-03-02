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

// IfExecOrVal executes a function or returns a value based on a condition.
// If the condition is true, it executes the provided function `t` and returns its result.
// If the condition is false, it returns the provided value `f`.
func IfExecOrVal[T any](condition bool, t func() T, f T) T {
	if condition {
		return t()
	}
	return f
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
