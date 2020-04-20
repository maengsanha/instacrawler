// Package meta implements meta-search.
package meta

import "fmt"

func Example_meta_Search() {
	secondLayer := []string{
		"golang",
	}
	thirdLayer := []string{
		"gophercon",
	}

	output := Search(secondLayer, thirdLayer)
	fmt.Printf("Second Layer: %s\n", output.SecondLayer)
	fmt.Printf("Third Layer: %s\n", output.ThirdLayer)
	// Output:
	//
}
