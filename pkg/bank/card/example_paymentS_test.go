package card

import (
	"fmt"
	"github.com/ssharifzoda/bank/pkg/bank/types"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 2200,
			Active:  true,
			PAN:     "5058 xxxx xxxx 5750",
		},
	}
	fmt.Print(PaymentSources(cards))
	//Output: 5058 xxxx xxxx 5750
	//[{card 5058 xxxx xxxx 5750 2200}]
}
