package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"os"
)

func main() {
	// Loading the CSV file
	f, err := os.Open("adult.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	// Filter the data: individuals with education level "HS-grad"
	hsGrad := df.Filter(dataframe.F{Colname: "education", Comparator: series.Eq, Comparando: "HS-grad"})
	fmt.Println("\nFiltered DataFrame (HS-grad):")
	fmt.Println(hsGrad)

	// Summarize the data: Average age of individuals in the dataset
	avgAge := df.Col("age").Mean()
	fmt.Printf("\nAverage age: %.2f\n", avgAge)

	// Describe the data
	fmt.Println("\nGenerate descriptive statistics:")
	description := df.Describe()
	fmt.Println(description)

}
