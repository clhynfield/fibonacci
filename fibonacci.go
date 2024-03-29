package fibonacci

type Fibonacci struct{}

// Fib returns the nth Fibonacci number
// given the number of steps into the sequence
func (f Fibonacci) Fib(steps int) (result int) {
	if steps == 1 || steps == 2 {
		result = 1
	} else if steps > 2 {
		result = f.Fib(steps-1) + f.Fib(steps-2)
	}
	return result
}
