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

	secondLayerChannel, thirdLayerChannel := make(chan core.InstaPost, len(secondLayerCrawlingResultSet)), make(chan core.InstaPost, len(secondLayerCrawlingResultSet))

	_AND(secondLayerCrawlingResultSet, thirdLayerCrawlingResultSet, secondLayerChannel, thirdLayerChannel)

	secondLayerResult, thirdLayerResult := make([]core.InstaPost, len(secondLayerChannel)), make([]core.InstaPost, len(thirdLayerChannel))

	for idx := 0; idx < len(secondLayerChannel); idx++ {
		secondLayerResult[idx] = <-secondLayerChannel
	}

	for idx := 0; idx < len(thirdLayerChannel); idx++ {
		thirdLayerResult[idx] = <-thirdLayerChannel
	}

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
func _AND(secondLayerPostSet, thirdLayerPostSet core.PostSet, secondLayerChannel, thirdLayerChannel chan core.InstaPost) {
	var syncer sync.WaitGroup

	defer close(secondLayerChannel)
	defer close(thirdLayerChannel)

	for id, post := range secondLayerPostSet {
		syncer.Add(1)
		go func(id string, post core.InstaPost) {
			defer syncer.Done()
			if _, exists := thirdLayerPostSet[id]; exists {
				thirdLayerChannel <- post
			} else {
				secondLayerChannel <- post
			}
		}(id, post)
	}

	syncer.Wait()
}
