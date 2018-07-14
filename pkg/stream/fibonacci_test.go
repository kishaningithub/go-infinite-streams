package stream_test

import (
	"github.com/kishaningithub/go-infinitestreams/pkg/stream"
	"testing"
	"math/big"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestStreamFibonacciReturnsFirst101FibonacciNos(t *testing.T)  {
	expected := []big.Int{
		getBigInt("0"),
		getBigInt("1"),
		getBigInt("1"),
		getBigInt("2"),
		getBigInt("3"),
		getBigInt("5"),
		getBigInt("8"),
		getBigInt("13"),
		getBigInt("21"),
		getBigInt("34"),
		getBigInt("55"),
		getBigInt("89"),
		getBigInt("144"),
		getBigInt("233"),
		getBigInt("377"),
		getBigInt("610"),
		getBigInt("987"),
		getBigInt("1597"),
		getBigInt("2584"),
		getBigInt("4181"),
		getBigInt("6765"),
		getBigInt("10946"),
		getBigInt("17711"),
		getBigInt("28657"),
		getBigInt("46368"),
		getBigInt("75025"),
		getBigInt("121393"),
		getBigInt("196418"),
		getBigInt("317811"),
		getBigInt("514229"),
		getBigInt("832040"),
		getBigInt("1346269"),
		getBigInt("2178309"),
		getBigInt("3524578"),
		getBigInt("5702887"),
		getBigInt("9227465"),
		getBigInt("14930352"),
		getBigInt("24157817"),
		getBigInt("39088169"),
		getBigInt("63245986"),
		getBigInt("102334155"),
		getBigInt("165580141"),
		getBigInt("267914296"),
		getBigInt("433494437"),
		getBigInt("701408733"),
		getBigInt("1134903170"),
		getBigInt("1836311903"),
		getBigInt("2971215073"),
		getBigInt("4807526976"),
		getBigInt("7778742049"),
		getBigInt("12586269025"),
		getBigInt("20365011074"),
		getBigInt("32951280099"),
		getBigInt("53316291173"),
		getBigInt("86267571272"),
		getBigInt("139583862445"),
		getBigInt("225851433717"),
		getBigInt("365435296162"),
		getBigInt("591286729879"),
		getBigInt("956722026041"),
		getBigInt("1548008755920"),
		getBigInt("2504730781961"),
		getBigInt("4052739537881"),
		getBigInt("6557470319842"),
		getBigInt("10610209857723"),
		getBigInt("17167680177565"),
		getBigInt("27777890035288"),
		getBigInt("44945570212853"),
		getBigInt("72723460248141"),
		getBigInt("117669030460994"),
		getBigInt("190392490709135"),
		getBigInt("308061521170129"),
		getBigInt("498454011879264"),
		getBigInt("806515533049393"),
		getBigInt("1304969544928657"),
		getBigInt("2111485077978050"),
		getBigInt("3416454622906707"),
		getBigInt("5527939700884757"),
		getBigInt("8944394323791464"),
		getBigInt("14472334024676221"),
		getBigInt("23416728348467685"),
		getBigInt("37889062373143906"),
		getBigInt("61305790721611591"),
		getBigInt("99194853094755497"),
		getBigInt("160500643816367088"),
		getBigInt("259695496911122585"),
		getBigInt("420196140727489673"),
		getBigInt("679891637638612258"),
		getBigInt("1100087778366101931"),
		getBigInt("1779979416004714189"),
		getBigInt("2880067194370816120"),
		getBigInt("4660046610375530309"),
		getBigInt("7540113804746346429"),
		getBigInt("12200160415121876738"),
		getBigInt("19740274219868223167"),
		getBigInt("31940434634990099905"),
		getBigInt("51680708854858323072"),
		getBigInt("83621143489848422977"),
		getBigInt("135301852344706746049"),
		getBigInt("218922995834555169026"),
		getBigInt("354224848179261915075"),
	}
	fibonacciNos := stream.Fibonacci(big.NewInt(101))

	var actual []big.Int
	for fibonacciNo := range fibonacciNos {
		actual = append(actual, fibonacciNo)
	}

	assert.Equal(t, expected, actual)
}

func getBigInt(valueStr string) big.Int  {
    value, ok := new(big.Int).SetString(valueStr, 10)
	if ok {
		return *value
	} else {
		panic(fmt.Sprintf("Invalid integer %s passed", valueStr))
	}
}