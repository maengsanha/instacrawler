package searcher

import "fmt"

func ExampleSearcher_TopSearch() {
	s := New()
	s.TopSearch("movie")
	fmt.Println(s.Hashtags)
	// Output:
	// #movie: 40.8m posts(40842605)
	// #movieclips: 172k posts(172026)
	// #movieposters: 164k posts(164510)
	// #moviereviews: 238k posts(238775)
	// #movienews: 190k posts(190426)
	// #moviestarplanet: 315k posts(315819)
	// #moviequote: 241k posts(241067)
	// #moviestars: 202k posts(202440)
	// #movies: 21.9m posts(21965623)
	// #moviemaking: 621k posts(621773)
	// #moviemaker: 280k posts(280731)
	// #moviefan: 187k posts(187957)
	// #movieart: 223k posts(223944)
	// #moviemagic: 221k posts(221918)
	// #movieaddict: 326k posts(326475)
	// #moviescene: 293k posts(293624)
	// #movielovers: 289k posts(289088)
	// #moviecollector: 339k posts(339683)
	// #moviereview: 703k posts(703842)
	// #movielover: 610k posts(610177)
	// #moviepremiere: 212k posts(212496)
	// #moviequotes: 944k posts(944253)
	// #moviescenes: 809k posts(809582)
	// #moviecollection: 432k posts(432388)
	// #moviebuff: 623k posts(623503)
}
