package stats

import "github.com/ssharifzoda/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var sum int
	znam := 0
	for _, payment := range payments {
		sum += int(payment.Amount)
		znam++
	}
	return types.Money(sum / znam)
}
