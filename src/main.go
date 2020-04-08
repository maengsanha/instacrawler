// Package main runs Instagram crawler api server.
package main

import (
	"fmt"
	"runtime"

	"github.com/joshua-dev/instacrawler/src/controllers/meta"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	secondLayer := []string{
		"옷",
		"clothes",
	}
	thirdLayer := []string{
		"봄",
		"spring",
	}

	secondLayerResult, thirdLayerResult := meta.Search(secondLayer, thirdLayer)
	fmt.Println("-*- second layer result -*-")
	for _, post := range secondLayerResult {
		fmt.Println(post)
	}
	fmt.Println("-*- third layer result -*-")
	for _, post := range thirdLayerResult {
		fmt.Println(post)
	}
}
