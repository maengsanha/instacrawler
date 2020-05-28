// Package meta implements Instagram meta-search engine logics.
package meta

import (
	"sync"

	"github.com/joshua-dev/instacrawler/controllers/crawler"
	"github.com/joshua-dev/instacrawler/core"
)

const (
	maxSuccessCnt int = 3
	maxDeathCnt   int = 3
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
			var (
				successCnt, deathCnt int
				endpoint             string
				crawlingResult       []core.InstaPost
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

	if len(secondLayer) == 0 {
		crawlingResultSet := _OR(crawlingResults)
		var crawlingResult []core.InstaPost
		for _, post := range crawlingResultSet {
			crawlingResult = append(crawlingResult, post)
		}

		return crawler.Response{
			SecondLayer:      nil,
			ThirdLayer:       crawlingResult,
			SecondLayerCache: endpoints[:len(secondLayer)],
			ThirdLayerCache:  endpoints[len(secondLayer):],
		}
	}

	secondLayerCrawlingResultSet, thirdLayerCrawlingResultSet := _OR(crawlingResults[:len(secondLayer)]), _OR(crawlingResults[len(secondLayer):])

	secondLayerChannel, thirdLayerChannel := make(chan core.InstaPost, len(secondLayerCrawlingResultSet)), make(chan core.InstaPost, len(secondLayerCrawlingResultSet))

	_AND(secondLayerCrawlingResultSet, thirdLayerCrawlingResultSet, secondLayerChannel, thirdLayerChannel)

	secondLayerResult, thirdLayerResult := make([]core.InstaPost, len(secondLayerChannel)), make([]core.InstaPost, len(thirdLayerChannel))

	var secondLayerSyncer sync.WaitGroup
	for idx := range secondLayerResult {
		secondLayerSyncer.Add(1)
		go func(i int) {
			defer secondLayerSyncer.Done()
			secondLayerResult[i] = <-secondLayerChannel
		}(idx)
	}
	secondLayerSyncer.Wait()

	var thirdLayerSyncer sync.WaitGroup
	for idx := range thirdLayerResult {
		thirdLayerSyncer.Add(1)
		go func(i int) {
			defer thirdLayerSyncer.Done()
			thirdLayerResult[i] = <-thirdLayerChannel
		}(idx)
	}
	thirdLayerSyncer.Wait()

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
