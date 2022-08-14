package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 100,
		},
		{
			ID:     2,
			Amount: 250,
		},
		{
			ID:     3,
			Amount: 90,
		},
		{
			ID:     4,
			Amount: 450,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 4
}
