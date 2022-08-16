package stats

import (
	"fmt"
	"github.com/ssharifzoda/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 1000,
		},
		{
			Amount: 200,
		},
	}
	fmt.Println(Avg(payments))
	//Output:600
}
