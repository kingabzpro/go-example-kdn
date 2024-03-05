package main

import (
    "os"
    "github.com/wcharczuk/go-chart"
)

func main() {
    // Create a new bar chart
    graph := chart.BarChart{
        Title: "Sales Data",
        Background: chart.Style{
            Padding: chart.Box{Top: 40, Left: 10},
        },
        Bars: []chart.Value{
            {Value: 10, Label: "Product A"},
            {Value: 20, Label: "Product B"},
            {Value: 30, Label: "Product C"},
        },
    }

    // Save the chart to an HTML file
    f, _ := os.Create("chart.html")
    defer f.Close()
    graph.Render(chart.SVG, f)
}