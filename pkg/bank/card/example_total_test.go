package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000,
		},
		{
			Balance: 10_000,
		},
	}
	fmt.Println(Total(cards))
	//Output: 20000
}
