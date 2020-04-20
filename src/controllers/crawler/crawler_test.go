// Package crawler implements Instagram crawling logic.
package crawler

import "fmt"

func ExampleCrawler_init() {
	crawler := New()
	crawler.init("kakao")
	fmt.Println(crawler)
	fmt.Println(len(crawler.InstaPosts))
	// Output:
	// count: 64956348
	// end_cursor: QVFBRC1fby1jQUZIMHlfRm14VGZlUDhSaV9nQ2FoQjdrQWMtRjctVlN0UTJieWl6Vjc1Q2QyUmhzbW9wSmYtRzgwb2t4dXVPWGRncmx0bU1seVlyd0tWTQ==
	// has_next_page: true
	// 49
}

func ExampleCrawler_next() {
	crawler := New()
	crawler.init("kakao")
	fmt.Println(crawler)
	fmt.Println(len(crawler.InstaPosts))
	crawler.next("kakao")
	fmt.Println(crawler)
	fmt.Println(len(crawler.InstaPosts))
	// Output:
	// count: 64956351
	// end_cursor: QVFCdzN5N2lUbmYxVW9yTHY0OE9XS0JmRGZlZG00Q0lkcUVLRHpkbjJidlRUVzVtUGR0dG8tR3FiWFFFQjRvVk96VF9WRVdldUhYZ013cVZQc2tMOEV0Qw==
	// has_next_page: true
	// 64
	// count: 64956350
	// end_cursor: QVFBdnY2X2FFcXQ2eDZaSWIzc2tUdl9UQVZDZ05kSDA3SXhiZGV3eVBaZFBnSlVxWmp4U0h6c2trS2NPaWdQX1Q3RUpzS3RRdEZtOUplb2ZaNUJoVWEwLQ==
	// has_next_page: true
	// 68
}

func ExampleCrawler_Crawl() {
	crawler := New()
	crawler.Crawl("kakao")
	fmt.Println(crawler)
	fmt.Println(len(crawler.InstaPosts))
	// Output:
	// count: 0
	// end_cursor:
	// has_next_page: false
	// 352
}
