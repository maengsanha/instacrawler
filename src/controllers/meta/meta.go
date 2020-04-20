// Package meta implements meta-search.
package meta

import (
	"sync"

	"github.com/joshua-dev/instacrawler/src/controllers/crawler"
	"github.com/joshua-dev/instacrawler/src/core"
)

// _AND implements AND operation.
func _AND(secondLayer, thirdLayer core.PostSet) (secondLayerResult []core.InstaPost, thirdLayerResult []core.InstaPost) {
	var syncer sync.WaitGroup

	for secondLayerPost := range secondLayer {
		for thirdLayerPost := range thirdLayer {
			syncer.Add(1)
			go func(secondLayerPost, thirdLayerPost core.InstaPost, secondLayer, thirdLayer []core.InstaPost) {
				defer syncer.Done()
				if secondLayerPost.ID == thirdLayerPost.ID {
					thirdLayerResult = append(thirdLayerResult, thirdLayerPost)
				} else {
					secondLayerResult = append(secondLayerResult, secondLayerPost)
				}
			}(secondLayerPost, thirdLayerPost, secondLayerResult, thirdLayerResult)
		}
	}
	syncer.Wait()

	return secondLayerResult, thirdLayerResult
}

// _OR implements OR operation.
func _OR(layer []*crawler.Crawler) core.PostSet {
	set := make(core.PostSet)
	for _, c := range layer {
		for _, post := range c.InstaPosts {
			if _, exists := set[post]; !exists {
				set[post] = struct{}{}
			}
		}
	}

	return set
}

// categorize implements categorization of meta-search.
func categorize(secondLayer, thirdLayer []*crawler.Crawler) ([]core.InstaPost, []core.InstaPost) {
	secondLayerPostSet := _OR(secondLayer)
	thirdLayerPostSet := _OR(thirdLayer)

	return _AND(secondLayerPostSet, thirdLayerPostSet)
}

// Search implements meta-search with search terms of second layer and third layer.
func Search(secondLayer, thirdLayer []string) crawler.Response {
	var queries []string = append(secondLayer, thirdLayer...)
	var workers []*crawler.Crawler
	var syncer sync.WaitGroup

	for i := 0; i < len(queries); i++ {
		workers = append(workers, crawler.New())
	}

	for index, worker := range workers {
		syncer.Add(1)
		go func(idx int, c *crawler.Crawler) {
			defer syncer.Done()
			c.Crawl(queries[idx])
		}(index, worker)
	}
	syncer.Wait()

	secondLayerResult, thirdLayerResult := categorize(workers[:len(secondLayer)], workers[len(secondLayer):])

	output := crawler.Response{
		SecondLayer: secondLayerResult,
		ThirdLayer:  thirdLayerResult,
	}

	return output
}
