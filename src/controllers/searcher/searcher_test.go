// Package searcher implements logics for crawling Instagram top search.

package searcher

import "fmt"

func ExampleSearcher_TopSearch() {
	s := New()
	hashtags, err := s.TopSearch("movie")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashtags))
	// Output:
	// [{"name":"movie","media_count":41311046,"search_result_subtitle":"41.3m"},{"name":"movieclips","media_count":180557,"search_result_subtitle":"180k"},{"name":"moviereviews","media_count":249222,"search_result_subtitle":"249k"},{"name":"movienews","media_count":193945,"search_result_subtitle":"193k"},{"name":"moviestarplanet","media_count":321489,"search_result_subtitle":"321k"},{"name":"moviemaker","media_count":285465,"search_result_subtitle":"285k"},{"name":"movieposters","media_count":169831,"search_result_subtitle":"169k"},{"name":"movies","media_count":22314183,"search_result_subtitle":"22.3m"},{"name":"moviemaking","media_count":629358,"search_result_subtitle":"629k"},{"name":"moviestars","media_count":206563,"search_result_subtitle":"206k"},{"name":"moviefan","media_count":189055,"search_result_subtitle":"189k"},{"name":"moviequote","media_count":247957,"search_result_subtitle":"247k"},{"name":"movieart","media_count":228793,"search_result_subtitle":"228k"},{"name":"moviemagic","media_count":224131,"search_result_subtitle":"224k"},{"name":"moviescene","media_count":303200,"search_result_subtitle":"303k"},{"name":"moviegeek","media_count":199760,"search_result_subtitle":"199k"},{"name":"movielovers","media_count":297843,"search_result_subtitle":"297k"}]
}
