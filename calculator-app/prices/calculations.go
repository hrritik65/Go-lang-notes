package calculations

import "fmt"

type TaxedPrices struct {
	TaxRate     float64
	InputPrice  []float64
	OutputPrice map[string]float64
}

func (job *TaxedPrices) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrice {
		// format key as string (price with 2 decimals)
		key := fmt.Sprintf("%.2f", price)
		result[key] = price * (1 + job.TaxRate)
	}

	job.OutputPrice = result
}

func NewTaxedPrices(taxRate float64) *TaxedPrices {
	return &TaxedPrices{
		InputPrice: []float64{19.99, 29.99, 4.99, 49.99},
		TaxRate:    taxRate,
	}
}
