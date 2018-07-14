package stream

import (
	"math/big"
)

// Even returns a stream of even numbers
func Even(limit *big.Int) <-chan big.Int {
	evenNos :=  make(chan big.Int)
	go func() {
		defer close(evenNos)
		isInfiniteStream := limit == nil
		noOfValuesGenerated := big.NewInt(0)
		for i := big.NewInt(2); isInfiniteStream || noOfValuesGenerated.Cmp(limit) < 0; i.Add(i, big.NewInt(2)) {
			evenNos <- *new(big.Int).Set(i)
			noOfValuesGenerated.Add(noOfValuesGenerated, big.NewInt(1))
		}
	}()
	return evenNos
}

