package main

import (
	"fmt"
	"practice/calculator-app/filemanager"
	calculations "practice/calculator-app/prices"
)

func main() {

	// prices := []float64{19.99, 29.99, 4.99, 49.99}
	taxRates := []float64{0.9, 0.08, 0.05, 0.09}

	// result := make(map[float64][]float64)       999

	for _, taxRate := range taxRates {
		// var taxedPrice []float64 = make([]float64, len(prices))
		// for pIndex, price := range prices {
		// 	taxedPrice[pIndex] = price * (1 + val)
		// }
		// result[val] = taxedPrice

		// using calculations package here

		// using filemanager package here
		fm := filemanager.New("calculator-app/prices.txt", fmt.Sprintf("calculator-app/results_%.0f.json", taxRate*100))
		priceJob := calculations.NewTaxedPrices(fm, taxRate)
		priceJob.Process()

	}

	// fmt.Println(result)        999
}
