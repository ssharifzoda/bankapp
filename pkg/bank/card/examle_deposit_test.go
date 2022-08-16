package card

import (
	"fmt"
	"github.com/ssharifzoda/bank/pkg/bank/types"
)

func ExampleDeposit_positive() {
	card := &types.Card{Balance: -10, Active: true}
	Deposit(card, 1000)
	fmt.Println(card.Balance)
	//Output:990
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
