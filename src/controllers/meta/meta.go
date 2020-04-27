// Package meta implements Instagram meta-search engine logics.
package meta

import (
	"fmt"
	"sync"

	"github.com/joshua-dev/instacrawler/src/controllers/crawler"
	"github.com/joshua-dev/instacrawler/src/core"
)

// Search implements meta-search with search terms of second layer and third layer.
func Search(secondLayer, thirdLayer, secondLayerCache, thirdLayerCache []string) crawler.Response {
	queries := append(secondLayer, thirdLayer...)
	caches := append(secondLayerCache, thirdLayerCache...)
	crawlers := make([]func() ([]core.InstaPost, string, error), len(queries))
	crawlingResults := make([][]core.InstaPost, len(queries))
	endCursors := make([]string, len(caches))

	var syncer sync.WaitGroup

	for idx, query := range queries {
		crawlers[idx] = crawler.Generator(query, caches[idx])
	}

	fmt.Println("successfully generated functions.")

	for idx, crawler := range crawlers {
		syncer.Add(1)
		go func(i int, f func() ([]core.InstaPost, string, error)) {
			defer syncer.Done()
			var success, failure int
			var endCursor string
			var crawlingResult []core.InstaPost
			for success < 3 && failure < 3 {
				posts, pagination, err := f()
				if err != nil {
					failure++
				} else {
					success++
					crawlingResult = append(crawlingResult, posts...)
					endCursor = pagination
				}
			}
			endCursors[i] = endCursor
			crawlingResults[i] = crawlingResult
		}(idx, crawler)
	}
	syncer.Wait()

	fmt.Println("successfully crawled.")

	secondLayerCrawlingResults, thirdLayerCrawlingResults := crawlingResults[:len(secondLayer)], crawlingResults[len(secondLayer):]
	secondLayerCrawlingResultSet, thirdLayerCrawlingResultSet := _OR(secondLayerCrawlingResults), _OR(thirdLayerCrawlingResults)

	fmt.Println("successfully performed OR operation.")

	secondLayerResult, thirdLayerResult := _AND(secondLayerCrawlingResultSet, thirdLayerCrawlingResultSet)

	fmt.Println("successfully performed AND operation.")

	return crawler.Response{
		SecondLayer:      secondLayerResult,
		ThirdLayer:       thirdLayerResult,
		SecondLayerCache: endCursors[:len(secondLayer)],
		ThirdLayerCache:  endCursors[len(secondLayer):],
	}
}

// _OR implements OR operation.
func _OR(posts [][]core.InstaPost) core.PostSet {
	set := make(core.PostSet)
	for _, postArr := range posts {
		for _, post := range postArr {
			if _, exists := set[post.ID]; !exists {
				set[post.ID] = post
			}
		}
	}
	return set
}

// _AND implements AND operation.
func _AND(secondLayerSet, thirdLayerSet core.PostSet) (second, third []core.InstaPost) {
	var syncer sync.WaitGroup

	for id, post := range secondLayerSet {
		syncer.Add(1)
		go func(i string, p core.InstaPost) {
			defer syncer.Done()
			if _, exists := thirdLayerSet[i]; exists {
				third = append(third, p)
			} else {
				second = append(second, p)
			}
		}(id, post)
	}

	syncer.Wait()
	return
}
