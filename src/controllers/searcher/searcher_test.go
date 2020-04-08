// Package searcher implements logics for crawling Instagram top search.

package searcher

import "fmt"

func ExampleSearcher_TopSearch() {
	s := New()
	hashtags, err := s.TopSearch("movie")
	if err != nil {
		panic(err)
	}

	for _, hashtag := range hashtags {
		fmt.Println(hashtag)
	}
	// Output:
	// movie 41.2m 41297855
	// movieclips 180k 180372
	// moviereviews 248k 248883
	// movienews 193k 193830
	// moviestarplanet 321k 321367
	// moviemaker 285k 285365
	// movieposters 169k 169671
	// movies 22.3m 22304530
	// moviemaking 629k 629117
	// moviestars 206k 206510
	// moviefan 190k 190315
	// moviequote 247k 247754
	// movieart 228k 228741
	// moviemagic 224k 224067
	// movieaddict 337k 337616
	// moviescene 302k 302932
	// moviegeek 200k 200139
	// movielovers 297k 297598
	// moviecollector 347k 347970
	// moviebuff 638k 638187
	// moviereview 728k 728530
	// movielover 626k 626544
}
