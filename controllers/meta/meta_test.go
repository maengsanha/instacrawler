// Package meta implements Instagram meta-search engine logics.
package meta

import "fmt"

func Example_meta_Search() {
	secondLayerQueries := []string{"cafe", "카페", "coffee", "커피", "일상"}
	thirdLayerQueries := []string{"starbucks", "스타벅스", "daily", "korea", "coffee"}
	secondLayerCache := []string{"", "", "", "", ""}
	thirdLayerCache := []string{"", "", "", "", ""}

	resp := Search(secondLayerQueries, thirdLayerQueries, secondLayerCache, thirdLayerCache)

	fmt.Println(len(resp.SecondLayer))
	fmt.Println(len(resp.ThirdLayer))
	fmt.Println(len(resp.SecondLayerCache))
	fmt.Println(len(resp.ThirdLayerCache))
	// Output:
	// 606
	// 319
	// 5
	// 5
}
