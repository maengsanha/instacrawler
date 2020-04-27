// Package meta implements Instagram meta-search engine logics.
package meta

import "fmt"

func Example_meta_Search() {
	secondLayerQueries := []string{"카페", "cafe", "카페추천", "카페투어"}
	thirdLayerQueries := []string{"스타벅스", "스벅", "starbucks"}
	secondLayerCache := []string{"", "", "", ""}
	thirdLayerCache := []string{"", "", ""}
	resp := Search(secondLayerQueries, thirdLayerQueries, secondLayerCache, thirdLayerCache)
	fmt.Println(len(resp.SecondLayer))
	fmt.Println(len(resp.ThirdLayer))
	fmt.Println(len(resp.SecondLayerCache))
	fmt.Println(len(resp.ThirdLayerCache))
	// Output:
	//
}
