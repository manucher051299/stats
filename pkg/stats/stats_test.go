package stats

import (
	"reflect"
	"testing"

	"github.com/manucher051299/bank/v2/pkg/types"
)

func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment
	result := CategoriesAvg(payments)
	if len(result) != 0 {
		t.Error("result length != 0")
	}
}
func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}
	result := CategoriesAvg(payments)
	if len(result) != 0 {
		t.Error("result length != 0")
	}
}
func TestCategoriesAvg_notFound(t *testing.T) {
	payments := []types.Payment{
		{Amount: 10, Category: "auto"},
		{Amount: 15, Category: "food"},
		{Amount: 20, Category: "auto"},
		{Amount: 30, Category: "auto"},
		{Amount: 15, Category: "food"},
		{Amount: 35, Category: "fun"},
	}
	expected := map[types.Category]types.Money{
		"auto": 20,
		"fun":  35,
		"food": 15,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result expected %v result %v", expected, result)
	}
}
