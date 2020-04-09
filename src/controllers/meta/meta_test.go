// Package meta implements meta-search.
package meta

import "fmt"

func Example_meta_Search() {
	secondLayer := []string{
		"옷",
		"clothes",
	}
	thirdLayer := []string{
		"봄",
		"spring",
	}

	output, err := Search(secondLayer, thirdLayer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}
