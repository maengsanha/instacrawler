// Package crawler implements Instagram crawling logic.
package crawler

import (
	"fmt"

	"github.com/joshua-dev/instacrawler/src/core"
)

func ExampleGenerator() {
	crawler := Generator("kakao", "")
	for i := 0; i < 3; i++ {
		posts, endCursor, err := crawler()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(len(posts))
		fmt.Println(endCursor)
	}
	// Output:
	// 70
	// QVFDMnEtTzBSYjlpLXR5aEMxNnhlS2ZncjJvVG12ZmtRRVNyZG1FUy16eHp2NERQRXFWTDNYZF95am5lMTNTV1phVGJYZktKVDAzS2w0MFdfTXJJQ3FqZw==
	// 71
	// QVFBcGlaX2ZGdGpNNVdLdGRwZHNJck9zSG10cmtJSmxPNWlxdHRhUHVDelJmcHBXWVlLb1Q0NC1QRmhlYWYwZFNlcDJhQldIUVlHZkRUZzFSX19zR1JaOA==
	// 70
	// QVFBaE1xNVYwbkJxR25tUGRWeWZrd1JxajlELUg0OWhnUzVhSzNkUVlXSV9wNG1JQS1jWkJDY0J1MXkyZEs4eVhOLThfZ2FERHVqTlJTR1NWXzA0OHJSTQ==
}

func ExampleGenerator2() {
	queries := []string{"kakao", "카카오"}
	caches := make([]string, len(queries))
	crawlers := make([]func() ([]core.InstaPost, string, error), len(queries))
	fmt.Println(crawlers)
	for idx, query := range queries {
		crawlers[idx] = Generator(query, caches[idx])
	}
	fmt.Println(crawlers)
	for i := 0; i < len(crawlers); i++ {
		posts, endCursor, err := crawlers[i]()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(len(posts))
		fmt.Println(endCursor)
	}
	// Output:
	//
}
