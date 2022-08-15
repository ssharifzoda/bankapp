package card

import (
	"github.com/ssharifzoda/bank/pkg/bank/types"
	"fmt"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}
func Deposit(card *types.Card, amount types.Money) {
	const depositLimit = 50_000
	if card.Active != false && amount <= depositLimit {
		card.Balance += amount
	}
}
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	const bonusLimit = 5000
	bonus := (int(card.MinBalance) * (percent) * (daysInMonth) / 100) / (daysInYear)
	if bonus > bonusLimit {
		bonus = bonusLimit
	}
	if card.Active != false && card.Balance > 0 {
		card.Balance += types.Money(bonus)
	}
}
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
