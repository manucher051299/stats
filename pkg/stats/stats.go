package stats

import (
	"github.com/manucher051299/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	ans := types.Money(0)
	for _, payment := range payments {
		ans += payment.Amount
	}
	return ans / types.Money(len(payments))
}
func TotalInCategory(payments []types.Payment, category string) types.Money {
	ans := types.Money(0)
	for _, payment := range payments {
		if string(payment.Category) == category {
			ans += payment.Amount
		}
	}
	return ans
}
