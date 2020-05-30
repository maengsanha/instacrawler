// Package meta implements Instagram meta-search engine logics.
package meta

import (
	"sync"

	"github.com/joshua-dev/instacrawler/controllers/crawler"
	"github.com/joshua-dev/instacrawler/core/instagram"
)

const (
	maxSuccessCnt int = 3
	maxDeathCnt   int = 3
)

// Search implements meta-search with search terms of second layer and third layer.
func Search(higherLayer, lowerLayer, higherLayerCache, lowerLayerCache []string) crawler.Response {
	queries := append(higherLayer, lowerLayer...)
	caches := append(higherLayerCache, lowerLayerCache...)
	workers := make([]func() ([]instagram.Post, string, error), len(queries))
	crawlingResults := make([][]instagram.Post, len(queries))
	endpoints := make([]string, len(caches))

	for idx, query := range queries {
		workers[idx] = crawler.Generator(query, caches[idx])
	}

	var syncer sync.WaitGroup
	for idx, worker := range workers {
		syncer.Add(1)
		go func(i int, f func() ([]instagram.Post, string, error)) {
			defer syncer.Done()
			var (
				successCnt, deathCnt int
				endpoint             string
				crawlingResult       []instagram.Post
			)
			for successCnt < maxSuccessCnt && deathCnt < maxDeathCnt {
				posts, endCursor, err := f()
				if err != nil {
					deathCnt++
				} else {
					successCnt++
					deathCnt = 0
					crawlingResult = append(crawlingResult, posts...)
					endpoint = endCursor
				}
			}
			endpoints[i] = endpoint
			crawlingResults[i] = crawlingResult
		}(idx, worker)
	}
	syncer.Wait()

	if len(higherLayer) == 0 {
		crawlingResultMap := _OR(crawlingResults)

		var crawlingResult []instagram.Post

		for _, post := range crawlingResultMap {
			crawlingResult = append(crawlingResult, post)
		}
		return crawler.Response{
			HigherLayer:      []instagram.Post{},
			LowerLayer:       crawlingResult,
			HigherLayerCache: endpoints[:len(higherLayerCache)],
			LowerLayerCache:  endpoints[len(higherLayerCache):],
		}
	}

	higherLayerCrawlingResultMap, lowerLayerCrawlingResultMap := _OR(crawlingResults[:len(higherLayer)]), _OR(crawlingResults[len(higherLayer):])

	higherLayerChannel, lowerLayerChannel := make(chan instagram.Post, len(higherLayerCrawlingResultMap)), make(chan instagram.Post, len(higherLayerCrawlingResultMap))

	_AND(higherLayerCrawlingResultMap, lowerLayerCrawlingResultMap, higherLayerChannel, lowerLayerChannel)

	higherLayerResult, lowerLayerResult := make([]instagram.Post, len(higherLayerChannel)), make([]instagram.Post, len(lowerLayerChannel))

	for idx := range higherLayerResult {
		syncer.Add(1)
		go func(i int) {
			defer syncer.Done()
			higherLayerResult[i] = <-higherLayerChannel
		}(idx)
	}
	syncer.Wait()

	for idx := range lowerLayerResult {
		syncer.Add(1)
		go func(i int) {
			defer syncer.Done()
			lowerLayerResult[i] = <-lowerLayerChannel
		}(idx)
	}
	syncer.Wait()

	return crawler.Response{
		HigherLayer:      higherLayerResult,
		LowerLayer:       lowerLayerResult,
		HigherLayerCache: endpoints[:len(higherLayerCache)],
		LowerLayerCache:  endpoints[len(higherLayerCache):],
	}
}

// _OR implements OR operation.
func _OR(posts [][]instagram.Post) instagram.PostMap {
	postMap := make(instagram.PostMap)
	for _, postArr := range posts {
		for _, post := range postArr {
			if _, exists := postMap[post.ID]; !exists {
				postMap[post.ID] = post
			}
		}
	}
	return postMap
}

// _AND implements AND operation.
func _AND(higherLayerPostMap, lowerLayerPostMap instagram.PostMap, higherLayerChannel, lowerLayerChannel chan instagram.Post) {
	var syncer sync.WaitGroup

	defer close(higherLayerChannel)
	defer close(lowerLayerChannel)

	for id, post := range higherLayerPostMap {
		syncer.Add(1)
		go func(id string, post instagram.Post) {
			defer syncer.Done()
			if _, exists := lowerLayerPostMap[id]; exists {
				lowerLayerChannel <- post
			} else {
				higherLayerChannel <- post
			}
		}(id, post)
	}

	syncer.Wait()
}
