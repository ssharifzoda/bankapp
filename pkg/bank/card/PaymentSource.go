package card

import (
	"fmt"
	"github.com/ssharifzoda/bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	paymentSource := make([]types.PaymentSource, len(cards))
	for i, card := range cards {
		if card.Balance > 0 && card.Active != false {
			paymentSource[i].Balance = card.Balance
			paymentSource[i].Type = "card"
			paymentSource[i].Number = string(card.PAN)
		}
		fmt.Println(paymentSource[i].Number)
		return paymentSource
	}
	return paymentSource
}
