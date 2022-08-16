package card

import "github.com/ssharifzoda/bank/pkg/bank/types"

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
