package card

import (
	"fmt"
	"github.com/ssharifzoda/bank/pkg/bank/types"
)

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10000,
		},
		{
			Balance: 10_000,
		},
	}
	fmt.Println(Total(cards))
	//Output: 20000
}
