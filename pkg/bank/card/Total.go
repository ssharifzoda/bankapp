package card

import "github.com/ssharifzoda/bank/pkg/bank/types"

func Total(cards []types.Card) types.Money {
	var sum int
	for _, card := range cards {
		if int(card.Balance) < 0 && card.Active == false {
			continue
		}
		sum += int(card.Balance)
	}
	return types.Money(sum)
}
