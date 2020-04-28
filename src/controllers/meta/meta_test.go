// Package meta implements Instagram meta-search engine logics.
package meta

import "fmt"

func Example_meta_Search() {
	secondLayerQueries := []string{"cafe", "cafestagram", "cafetour"}
	thirdLayerQueries := []string{"starbucks", "starbuckskorea", "starbuckscoffee"}
	secondLayerCache := []string{"", "", ""}
	thirdLayerCache := []string{"", "", ""}
	resp := Search(secondLayerQueries, thirdLayerQueries, secondLayerCache, thirdLayerCache)
	fmt.Println(len(resp.SecondLayer))
	fmt.Println(len(resp.ThirdLayer))
	fmt.Println(len(resp.SecondLayerCache))
	fmt.Println(len(resp.ThirdLayerCache))
	// Output:
	// 399
	// 1
	// 3
	// 3
}
