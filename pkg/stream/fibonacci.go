package stream

import (
	"math/big"
)

// Even returns a stream of fibonacci numbers
func Fibonacci(limit *big.Int) <-chan big.Int {
	fibonacciNos :=  make(chan big.Int)
	go func() {
		defer close(fibonacciNos)
		fibs := []*big.Int{big.NewInt(0), big.NewInt(1)}
		isInfiniteStream := limit == nil
		for noOfValuesGenerated := big.NewInt(0); isInfiniteStream || noOfValuesGenerated.Cmp(limit) < 0 ; noOfValuesGenerated.Add(noOfValuesGenerated, big.NewInt(1)) {
			nextFib := big.NewInt(0).Add(fibs[0], fibs[1])
			fibs = append(fibs, nextFib)
			var currentFib *big.Int
			currentFib, fibs = fibs[0], fibs[1:]
			fibonacciNos <- *new(big.Int).Set(currentFib)
		}
	}()
	return fibonacciNos
}

