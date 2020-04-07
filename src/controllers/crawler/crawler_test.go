// Package crawler implements Instagram crawling logic.
package crawler

import "fmt"

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
	// count: 123857838
	// end_cursor: QVFDQ1RSbGxtcWROczJlaUlBSVFfOWFIb25mS0Q1SnJIX2VyekpvRDcxR3ZhYlluVm5oYUVJMW9vUEZLZWFrbVdxVlZfRkpzTmRPdnpnLS1aUVlfMkdncA==
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
	// count: 123857867
	// end_cursor: QVFDQTdBWUl4bWtBdTg1MlZXUF9tWlluMjUtUFRobDdNN0o0TWkyRzI2c05mVDQzWXBBR1BmU0xYcmJHQ0RSYVRDRl8tTzB6MXJLODRoLWlPby10ZWZSSA==
	// has_next_page: true
	//
	// count: 123857879
	// end_cursor: QVFBVU5mdl8wbkJicmcxZllsTDlfYkdsZGVFM3FqRkxXMzNWbTItaExralR5ZnREZzhKdGxfUlpzd2dKNDdjNzJERlRPU2xkY2ZyM2l5Q2l2VWNvaERDZg==
	// has_next_page: true
}

func ExampleCrawler_Crawl() {
	crawler := New()
	crawler.Crawl("gophercon")
	fmt.Printf("count: %d\n", crawler.Count)
	fmt.Printf("end_cursor: %s\n", crawler.EndCursor)
	fmt.Printf("has_next_page: %t\n", crawler.HasNextPage)
	// Output:
	// count: 923
	// end_cursor:
	// has_next_page: false
}
