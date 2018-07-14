package stream_test

import (
	"github.com/kishaningithub/go-infinitestreams/pkg/stream"
	"testing"
	"math/big"
	"github.com/stretchr/testify/assert"
)

func TestStreamEvenReturnsFirst10EvenNos(t *testing.T)  {
	expected := []big.Int{
		*big.NewInt(2), *big.NewInt(4), *big.NewInt(6), *big.NewInt(8),
		*big.NewInt(10), *big.NewInt(12), *big.NewInt(14), *big.NewInt(16),
		*big.NewInt(18), *big.NewInt(20),
	}

	evenNos := stream.Even(big.NewInt(10))

	var actual []big.Int
	for evenNo := range evenNos {
		actual = append(actual, evenNo)
	}

	assert.Equal(t, expected, actual)
}