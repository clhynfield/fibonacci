package fibonacci_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/clhynfield/fibonacci"
)

var _ = Describe("Fibonacci", func() {
	var fibonacci Fibonacci

	DescribeTable("Calculating Fibonacci numbers",
		func(steps int, expected int) {
			Expect(fibonacci.Fib(steps)).To(Equal(expected))
		},
		Entry("fib(0)", 0, 0),
		Entry("fib(1)", 1, 1),
	)
})
