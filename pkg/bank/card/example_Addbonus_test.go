package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleAddBonus_positive() {
	card := &types.Card{Balance: 10, Active: true, MinBalance: 10000}
	AddBonus(card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 34
}
func ExampleAddBonus_inactive() {
	card := &types.Card{Balance: 10, Active: false, MinBalance: 10000}
	AddBonus(card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 10
}
func ExampleAddBonus_noBalance() {
	card := &types.Card{Balance: -10, Active: true, MinBalance: 10000}
	AddBonus(card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: -10
}
func ExampleAddBonus_unlimit() {
	card := &types.Card{Balance: 10, Active: true, MinBalance: 150_000000}
	AddBonus(card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 5010
}
