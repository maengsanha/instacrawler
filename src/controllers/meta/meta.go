// Package meta implements Instagram meta-search engine logics.
package meta

import (
	"sync"

	"github.com/joshua-dev/instacrawler/src/controllers/crawler"
	"github.com/joshua-dev/instacrawler/src/core"
)

const (
	maxSuccessCount int = 3
	maxFailureCount int = 3
)

// Search implements meta-search with search terms of second layer and third layer.
func Search(secondLayer, thirdLayer, secondLayerCache, thirdLayerCache []string) crawler.Response {
	queries := append(secondLayer, thirdLayer...)
	caches := append(secondLayerCache, thirdLayerCache...)
	workers := make([]func() ([]core.InstaPost, string, error), len(queries))
	crawlingResults := make([][]core.InstaPost, len(queries))
	endpoints := make([]string, len(caches))

	for idx, query := range queries {
		workers[idx] = crawler.Generator(query, caches[idx])
	}

	var syncer sync.WaitGroup
	for idx, worker := range workers {
		syncer.Add(1)
		go func(i int, f func() ([]core.InstaPost, string, error)) {
			defer syncer.Done()
			var success, failure int
			var endpoint string
			var crawlingResult []core.InstaPost
			for success < maxSuccessCount && failure < maxFailureCount {
				posts, endCursor, err := f()
				if err != nil {
					failure++
				} else {
					success++
					crawlingResult = append(crawlingResult, posts...)
					endpoint = endCursor
				}
			}
			endpoints[i] = endpoint
			crawlingResults[i] = crawlingResult
		}(idx, worker)
	}
	syncer.Wait()

	secondLayerCrawlingResultSet, thirdLayerCrawlingResultSet := _OR(crawlingResults[:len(secondLayer)]), _OR(crawlingResults[len(secondLayer):])
	secondLayerResult, thirdLayerResult := _AND(secondLayerCrawlingResultSet, thirdLayerCrawlingResultSet)

	return crawler.Response{
		SecondLayer:      secondLayerResult,
		ThirdLayer:       thirdLayerResult,
		SecondLayerCache: endpoints[:len(secondLayer)],
		ThirdLayerCache:  endpoints[len(secondLayer):],
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
func _AND(secondLayerPostSet, thirdLayerPostSet core.PostSet) (secondLayerPosts, thirdLayerPosts []core.InstaPost) {
	var syncer sync.WaitGroup

	for id, post := range secondLayerPostSet {
		syncer.Add(1)
		go func(i string, p core.InstaPost) {
			defer syncer.Done()
			if _, exists := thirdLayerPostSet[i]; exists {
				thirdLayerPosts = append(thirdLayerPosts, p)
			} else {
				secondLayerPosts = append(secondLayerPosts, p)
			}
		}(id, post)
	}

	syncer.Wait()
	return
}
