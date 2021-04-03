package stats

import (
	"fmt"

	"github.com/manucher051299/bank/v2/pkg/types"
)

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 10,
		},
		{
			Amount: 15,
		},
		{
			Amount: 5,
		},
	}))
	//Output
	//10
}

func ExampleTotalInCategory() {
	fmt.Println(TotalInCategory([]types.Payment{
		{
			Amount:   10,
			Category: "s",
		},
		{
			Amount:   15,
			Category: "w",
		},
		{
			Amount:   5,
			Category: "s",
		},
	}, "s"))
	//Output
	//15
}
