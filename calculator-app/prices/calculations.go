package calculations

import (
	// to read from file and deal with io
	"fmt"
	"practice/calculator-app/conversion"
	"practice/calculator-app/filemanager"
)

type TaxedPricesJob struct {
	IOManager   filemanager.FileManager `json:"-"`
	TaxRate     float64                 `json:"tax_rate"`
	InputPrice  []float64               `json:"-"`
	OutputPrice map[string]string       `json:"prices"`
}

// fucntion to read prices from file
func (Job *TaxedPricesJob) ReadPricesFromFile() {

	lines, err := Job.IOManager.ReadLines("calculator-app/prices.txt")

	if err != nil {
		fmt.Println("Error reading file step 1:", err)
		return
	}

	// this below code is used as a package

	// // to convert the price from string to float
	// prices := make([]float64, len(lines))

	// for i, line := range lines {
	// 	floatPrice, err := strconv.ParseFloat(line, 64)
	// 	if err != nil {
	// 		fmt.Println("Error converting string to float:", err)
	// 		return
	// 	}
	// 	prices[i] = floatPrice
	// }

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println("Error reading file step 2:", err)
		return
	}

	Job.InputPrice = prices
}

func (job *TaxedPricesJob) Process() {
	job.ReadPricesFromFile()
	result := make(map[string]string)

	for _, price := range job.InputPrice {
		// format key as string (price with 2 decimals)
		taxedPrices := price * (1 + job.TaxRate)
		FinalKey := fmt.Sprintf("%.2f", price)
		FinalValue := fmt.Sprintf("%.2f", taxedPrices)
		result[FinalKey] = FinalValue
	}

	// job.OutputPrice = result
	// fmt.Println(result)

	job.OutputPrice = result

	job.IOManager.WriteJson(job)

}

func NewTaxedPrices(fm filemanager.FileManager, taxRate float64) *TaxedPricesJob {
	return &TaxedPricesJob{
		IOManager: fm,
		// InputPrice: []float64{19.99, 29.99, 4.99, 49.99},
		TaxRate: taxRate,
	}
}
