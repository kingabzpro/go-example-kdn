package main

import (
	"fmt"
	"github.com/go-gota/gota/series"
)

func main() {
	fmt.Println(series.New([]string{"z", "y", "d", "e"}, series.String, "col"))
	
	a := map[string]series.Type{
		"A": series.String,
		"D": series.Bool,
	}

	fmt.Println(a)
}