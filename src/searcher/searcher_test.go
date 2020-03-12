package searcher

import "fmt"

func ExampleSearcher_TopSearch() {
	s := NewSearcher()
	s.TopSearch("가구")
	fmt.Println(s.Hashtags)
	// Output:
	// #가구: 771k posts(771837)
	// #가구인테리어: 76k posts(76060)
	// #가구디자인: 110k posts(110138)
	// #가구제작: 64.2k posts(64229)
	// #가구스타그램: 186k posts(186802)
	// #가구공방: 68.8k posts(68838)
	// #가구리폼: 13.8k posts(13877)
	// #가구배치: 12.9k posts(12941)
	// #가구디자이너: 10.2k posts(10278)
	// #가구촬영: 8386 posts(8386)
}
