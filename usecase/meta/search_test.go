package meta_test

import (
	"fmt"

	"github.com/maengsanha/instacrawler/usecase/meta"
)

func Example_meta_Search() {
	secondLayerQueries := []string{"cafe", "카페", "coffee", "커피", "일상"}
	thirdLayerQueries := []string{"starbucks", "스타벅스", "daily", "korea", "coffee"}
	secondLayerCache := []string{"", "", "", "", ""}
	thirdLayerCache := []string{"", "", "", "", ""}

	resp := meta.Search(secondLayerQueries, thirdLayerQueries, secondLayerCache, thirdLayerCache)

	fmt.Println(len(resp.HigherLayer))
	fmt.Println(len(resp.LowerLayer))
	fmt.Println(len(resp.HigherLayerCache))
	fmt.Println(len(resp.LowerLayerCache))
	// Output:
	// 680
	// 322
	// 5
	// 5
}
