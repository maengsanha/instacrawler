package crawler

const requestBaseURL = "https://instagram.com/explore/tags/"

// Crawl returns the URL of each post, cover photo URL, tag information, and number of likes found in the tag search by query in the Posting slice.
func Crawl(query string) []Posting {

}

// CrawlNumCoreSpriteHashtag returns the number of posts in the tag search url.
func CrawlNumCoreSpriteHashtag(query string) int {
	const accessorClassName = ".g47SY"
}

// CoreSpriteHashtag accesses to the url, returns a Posting structure containing cover photo URL, tag information, and number of likes
func CoreSpriteHashtag(url string) Posting {
	const accessorClassName = ".v1Nh3"
}

// CoreSpriteHashtagSrc returns cover photo URL.
func CoreSpriteHashtagSrc(url string) string {
	const accessorClassName = "._9AhH0"
}

// CoreSpriteHashtagTagsInfo returns hashtag informations.
func CoreSpriteHashtagTagsInfo(url string) string {
	const accessorClassName = ".xil3i"
}

// CoreSpriteHashtagLike returns number of likes.
func CoreSpriteHashtagLike(url string) int {
	const accessorClassName = ".sqdOP"
}
