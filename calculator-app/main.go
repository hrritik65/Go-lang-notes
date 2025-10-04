package main

import "fmt"

func main() {

	prices := []float64{19.99, 29.99, 4.99, 49.99}
	taxRate := []float64{0.07, 0.08, 0.05, 0.09}

	result := make(map[float64][]float64)

	for _, val := range taxRate {
		var taxedPrice []float64 = make([]float64, len(prices))
		for pi, price := range prices {
			taxedPrice[pi] = price * (1 + val)
		}
		result[val] = taxedPrice

	}

	fmt.Println(result)
}
