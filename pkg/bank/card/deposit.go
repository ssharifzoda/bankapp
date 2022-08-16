package card

import "github.com/ssharifzoda/bank/pkg/bank/types"

func Deposit(card *types.Card, amount types.Money) {
	const depositLimit = 50_000
	if card.Active != false && amount <= depositLimit {
		card.Balance += amount
	}
}
