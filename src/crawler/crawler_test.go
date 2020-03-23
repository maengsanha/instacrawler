package crawler

import (
	"fmt"
)

func ExampleCrawler_init() {
	crawler := New()
	if err := crawler.init("daily"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("count: %d\n", crawler.Count)
	fmt.Printf("end_cursor: %s\n", crawler.EndCursor)
	fmt.Printf("has_next_page: %t\n", crawler.HasNextPage)
	// Output:
	// count: 122882781
	// end_cursor: QVFDS092UGVTRGhGNTkzbjVzZFZsUk9rcm1CRU80bmlzVnF6YnR3WHRwYTFZeTFkR1dUcmFmaHBGOUJLOEVja28yYTBPdnVxeTZpcW1PRVNxYzFjS1ZJeA==
	// has_next_page: true
}

func ExampleCrawler_next() {
	crawler := New()
	if err := crawler.init("daily"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("count: %d\n", crawler.Count)
	fmt.Printf("end_cursor: %s\n", crawler.EndCursor)
	fmt.Printf("has_next_page: %t\n\n", crawler.HasNextPage)
	if err := crawler.next("daily"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("count: %d\n", crawler.Count)
	fmt.Printf("end_cursor: %s\n", crawler.EndCursor)
	fmt.Printf("has_next_page: %t\n", crawler.HasNextPage)
	// Output:
	// count: 122882914
	// end_cursor: QVFCSVJCcDU4QTVZdndLaGNnQl9XenZwcFltZ00zVEZJMm01QjlkU2tya1YxN2dzUkpNT25zbEFUTWNGdkMydEhtTzFlZU8xdGp6WktyaWktVVVBVXBmWA==
	// has_next_page: true

	// count: 122882909
	// end_cursor: QVFBVTZid1BaampEYThfaE5CYUNjVl8xY0VaTnE3czgxOTlNRTQxY1Itb2xsb21JSWJtSnQyMkpKdGw3QV9WNm1MYm01dEFIUkUyc2hOSmJHdnRGWXdpNw==
	// has_next_page: true
}

func ExampleCrawler_Crawl() {
	crawler := New()
	crawler.Crawl("gophercon")
	fmt.Printf("count: %d\n", crawler.Count)
	fmt.Printf("end_cursor: %s\n", crawler.EndCursor)
	fmt.Printf("has_next_page: %t\n", crawler.HasNextPage)
	// Output:
	// CPU usage: 16
	// count: 922
	// end_cursor:
	// has_next_page: false
}
