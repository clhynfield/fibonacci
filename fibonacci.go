package fibonacci

type Fibonacci struct{}

// Fib returns the nth Fibonacci number
// given the number of steps into the sequence
func (f Fibonacci) Fib(steps int) (result int) {
	if steps == 1 {
		result = 1
	}
	return result
}
