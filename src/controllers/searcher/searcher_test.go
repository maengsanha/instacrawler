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
	// movie 41.4m 41454882
	// movieclips 182k 182967
	// moviereviews 252k 252399
	// moviestarplanet 323k 323489
	// moviemaker 286k 286415
	// movieposters 171k 171286
	// movies 22.4m 22416410
	// moviemaking 631k 631359
	// moviestars 207k 207607
	// moviefan 190k 190242
	// moviequote 249k 249740
	// movienews 195k 195215
	// movieart 230k 230841
	// moviemagic 224k 224763
	// moviescene 306k 306339
	// moviegeek 201k 201179
	// movielovers 300k 300531
	// moviecollector 351k 351635
	// movieaddict 341k 341254
}
