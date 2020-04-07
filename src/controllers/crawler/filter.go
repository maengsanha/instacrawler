// Package crawler implements Instagram crawling logic.
package crawler

import (
	"sync"

	"github.com/joshua-dev/instacrawler/src/core"
)

// or implements OR operation.
func or(layer []*Crawler) (set core.PostSet) {
	for _, c := range layer {
		for _, post := range c.InstaPosts {
			if _, exists := set[post]; !exists {
				set[post] = struct{}{}
			}
		}
	}

	return set
}

// and implements AND operation.
func and(secondLayer, thirdLayer core.PostSet) (secondLayerResult *core.InstaPosts, thirdLayerResult *core.InstaPosts) {
	var syncer sync.WaitGroup

	for secondLayerPost := range secondLayer {
		for thirdLayerPost := range thirdLayer {
			syncer.Add(1)
			go func(secondLayerPost, thirdLayerPost core.InstaPost, secondLayer, thirdLayer *core.InstaPosts) {
				defer syncer.Done()
				if secondLayerPost.ID == thirdLayerPost.ID {
					*thirdLayer = append([]core.InstaPost(*thirdLayer), thirdLayerPost)
				} else {
					*secondLayer = append([]core.InstaPost(*secondLayer), secondLayerPost)
				}
			}(secondLayerPost, thirdLayerPost, secondLayerResult, thirdLayerResult)
		}
	}
	syncer.Wait()

	return secondLayerResult, thirdLayerResult
}

// Meta implements meta-search with search terms of layer 2 and layer 3.
func Meta(secondLayer, thirdLayer []*Crawler) (secondLayerResult *core.InstaPosts, thirdLayerResult *core.InstaPosts) {
	secondLayerPostSet := or(secondLayer)
	thirdLayerPostSet := or(thirdLayer)

	return and(secondLayerPostSet, thirdLayerPostSet)
}
