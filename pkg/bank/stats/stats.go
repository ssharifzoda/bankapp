package stats

import "github.com/ssharifzoda/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var sum int
	for _, payment := range payments {
		sum += int(payment.Amount)
	}
	return types.Money(sum / len(payments))
}
