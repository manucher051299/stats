package stats

import (
	"github.com/manucher051299/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	ans := types.Money(0)
	cnt := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			ans += payment.Amount
			cnt++
		}
	}
	return ans / types.Money(cnt)
}
func TotalInCategory(payments []types.Payment, category string) types.Money {
	ans := types.Money(0)
	for _, payment := range payments {
		if string(payment.Category) == category && payment.Status != types.StatusFail {
			ans += payment.Amount
		}
	}
	return ans
}
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := make(map[types.Category]types.Money)
	categoryLength := make(map[types.Category]types.Money)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			categories[payment.Category] += payment.Amount
			categoryLength[payment.Category]++
		}
	}
	for key, category := range categoryLength {
		categories[key] /= category
	}
	return categories
}
func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	ans := make(map[types.Category]types.Money)
	for key := range first {
		ans[key] = second[key] - first[key]
	}
	for key := range second {
		ans[key] = second[key] - first[key]
	}
	return ans
}
