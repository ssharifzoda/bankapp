package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleDeposit_positive() {
	card := &types.Card{Balance: -10, Active: true}
	Deposit(card, 100)
	fmt.Println(card.Balance)
	//Output:90
}
func ExampleDeposit_inactive() {
	card := &types.Card{Balance: 10, Active: false}
	Deposit(card, 5001)
	fmt.Println(card.Balance)
	//Output:10
}
func ExampleDeposit_unlimit() {
	card := &types.Card{Balance: 10, Active: true}
	Deposit(card, 50_001)
	fmt.Println(card.Balance)
	//Output:10
}
