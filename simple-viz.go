package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	// Sample data: replace this CSV string with the path to your actual data file or another data source.
	f, err := os.Open("adult.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	
	// Read the data into a DataFrame.
	df := dataframe.ReadCSV(f)

	// Extract the 'age' column and convert it to a slice of float64s for plotting.
	ages := df.Col("age").Float()

	// Create a new plot.
	p:= plot.New()

	p.Title.Text = "Age Distribution"
	p.X.Label.Text = "Age"
	p.Y.Label.Text = "Frequency"

	// Create a histogram of the 'age' column.
	h, err := plotter.NewHist(plotter.Values(ages), 16) // 16 bins.
	if err != nil {
		log.Fatal(err)
	}
	h.FillColor = color.RGBA{R: 255, A: 255}

	p.Add(h)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "age_distribution.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Histogram saved as age_distribution.png")
}
