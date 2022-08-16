package card

import (
	"github.com/ssharifzoda/bank/pkg/bank/types"
)

func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Active == true && amount <= card.Balance && amount > 0 && amount <= 2000000 {
		card.Balance = card.Balance - amount
	}

	return card
}
