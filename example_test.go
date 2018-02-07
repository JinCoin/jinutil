package jinutil_test

import (
	"fmt"
	"math"

	"github.com/JinCoin/jinutil"
)

func ExampleAmount() {

	a := jinutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = jinutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = jinutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 JIN
	// 100,000,000 Satoshis: 1 JIN
	// 100,000 Satoshis: 0.001 JIN
}

func ExampleNewAmount() {
	amountOne, err := jinutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := jinutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := jinutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := jinutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 JIN
	// 0.01234567 JIN
	// 0 JIN
	// invalid jincoin amount
}

func ExampleAmount_unitConversions() {
	amount := jinutil.Amount(44433322211100)

	fmt.Println("Satoshi to kJIN:", amount.Format(jinutil.AmountKiloBTC))
	fmt.Println("Satoshi to JIN:", amount)
	fmt.Println("Satoshi to MilliJIN:", amount.Format(jinutil.AmountMilliBTC))
	fmt.Println("Satoshi to MicroJIN:", amount.Format(jinutil.AmountMicroBTC))
	fmt.Println("Satoshi to Satoshi:", amount.Format(jinutil.AmountSatoshi))

	// Output:
	// Satoshi to kJIN: 444.333222111 kJIN
	// Satoshi to JIN: 444333.222111 JIN
	// Satoshi to MilliJIN: 444333222.111 mJIN
	// Satoshi to MicroJIN: 444333222111 μJIN
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
