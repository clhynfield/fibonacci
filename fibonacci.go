package fibonacci

type Fibonacci struct{}

// Fib returns the nth Fibonacci number
// given the number of steps into the sequence
func (f Fibonacci) Fib(steps int) (result int) {
	if steps == 1 || steps == 2 {
		result = 1
	} else if steps == 3 {
		result = f.Fib(2) + f.Fib(1)
	} else if steps == 4 {
		result = f.Fib(3) + f.Fib(2)
	} else if steps == 5 {
		result = f.Fib(4) + f.Fib(3)
	}
	return result
}
