// Package meta implements meta-search.
package meta

import (
	"fmt"
)

func Example_meta_Search() {
	secondLayer := []string{
		"옷",
		"clothes",
	}
	thirdLayer := []string{
		"봄",
		"spring",
	}

	secondLayerResult, thirdLayerResult := Search(secondLayer, thirdLayer)
	fmt.Println("-*- second layer result -*-")
	for _, post := range secondLayerResult {
		fmt.Println(post)
	}
	fmt.Println("-*- third layer result -*-")
	for _, post := range thirdLayerResult {
		fmt.Println(post)
	}
}
