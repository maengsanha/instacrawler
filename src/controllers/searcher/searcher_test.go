// Package searcher implements logics for crawling Instagram top search.

package searcher

import "fmt"

func ExampleSearcher_TopSearch() {
	s := New()
	s.TopSearch("movie")
	fmt.Println(s.Hashtags)
	// Output:
	// #movie: 41.2m posts(41232722)
	// #movieclips: 179k posts(179343)
	// #moviequote: 246k posts(246793)
	// #moviereviews: 247k posts(247418)
	// #movienews: 193k posts(193396)
	// #moviestarplanet: 320k posts(320463)
	// #moviemaker: 284k posts(284819)
	// #movieposters: 169k posts(169022)
	// #movies: 22.2m posts(22256907)
	// #moviemaking: 627k posts(627936)
	// #moviestars: 206k posts(206072)
	// #movieart: 228k posts(228137)
	// #moviemagic: 223k posts(223736)
	// #moviefan: 191k posts(191057)
	// #movieaddict: 336k posts(336163)
	// #moviescene: 301k posts(301807)
	// #movielovers: 296k posts(296353)
	// #moviecollector: 347k posts(347445)
	// #moviereview: 725k posts(725158)
	// #movielover: 624k posts(624300)
	// #moviepremiere: 214k posts(214260)
	// #moviequotes: 973k posts(973445)
	// #moviecollection: 441k posts(441206)
	// #moviebuff: 637k posts(637040)
}
