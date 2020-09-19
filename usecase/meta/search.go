package meta

import (
	"sync"

	"github.com/maengsanha/instacrawler/model/meta"

	"github.com/maengsanha/instacrawler/dataservice/instagram"
	insta "github.com/maengsanha/instacrawler/model/instagram"
)

const (
	MAX_SUCCESS_CNT = 3
	MAX_DEATH_CNT   = 3
)

// Search implements meta-search with search terms of higher layer and lower layer.
func Search(higherLayer, lowerLayer, higherLayerCache, lowerLayerCache []string) meta.Response {
	queries := append(higherLayer, lowerLayer...)
	caches := append(higherLayerCache, lowerLayerCache...)
	workers := make([]func() ([]insta.Post, string, error), len(queries))
	crawlingResults := make([][]insta.Post, len(queries))
	endpoints := make([]string, len(caches))

	for idx, query := range queries {
		workers[idx] = instagram.PageParserGenerator(query, caches[idx])
	}

	var syncer sync.WaitGroup
	for idx, worker := range workers {
		syncer.Add(1)
		go func(i int, f func() ([]insta.Post, string, error)) {
			defer syncer.Done()
			var (
				successCnt, deathCnt int
				endpoint             string
				crawlingResult       []insta.Post
			)
			for successCnt < MAX_SUCCESS_CNT && deathCnt < MAX_DEATH_CNT {
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

		var crawlingResult []insta.Post

		for _, post := range crawlingResultMap {
			crawlingResult = append(crawlingResult, post)
		}
		return meta.Response{
			HigherLayer:      []insta.Post{},
			LowerLayer:       crawlingResult,
			HigherLayerCache: endpoints[:len(higherLayerCache)],
			LowerLayerCache:  endpoints[len(higherLayerCache):],
		}
	}

	higherLayerCrawlingResultMap, lowerLayerCrawlingResultMap := _OR(crawlingResults[:len(higherLayer)]), _OR(crawlingResults[len(higherLayer):])

	higherLayerChannel, lowerLayerChannel := make(chan insta.Post, len(higherLayerCrawlingResultMap)), make(chan insta.Post, len(higherLayerCrawlingResultMap))

	_AND(higherLayerCrawlingResultMap, lowerLayerCrawlingResultMap, higherLayerChannel, lowerLayerChannel)

	higherLayerResult, lowerLayerResult := make([]insta.Post, len(higherLayerChannel)), make([]insta.Post, len(lowerLayerChannel))

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

	return meta.Response{
		HigherLayer:      higherLayerResult,
		LowerLayer:       lowerLayerResult,
		HigherLayerCache: endpoints[:len(higherLayerCache)],
		LowerLayerCache:  endpoints[len(higherLayerCache):],
	}
}

// _OR implements OR operation.
func _OR(posts [][]insta.Post) insta.PostMap {
	postMap := make(insta.PostMap)
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
func _AND(higherLayerPostMap, lowerLayerPostMap insta.PostMap, higherLayerChannel, lowerLayerChannel chan insta.Post) {
	var syncer sync.WaitGroup

	defer close(higherLayerChannel)
	defer close(lowerLayerChannel)

	for id, post := range higherLayerPostMap {
		syncer.Add(1)
		go func(id string, post insta.Post) {
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
