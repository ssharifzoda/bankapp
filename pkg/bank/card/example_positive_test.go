package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	//Output: 1000000
}
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0}, 1)
	fmt.Println(result.Balance)
	//Output: 0
}
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Active: false}, 1)
	fmt.Println(result.Balance)
	//Output: 0
}
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 200000_00}, 20_000_01)
	fmt.Println(result.Balance)
	//Output: 20000000
}
