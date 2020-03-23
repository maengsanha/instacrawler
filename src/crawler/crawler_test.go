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
	// Requesting to https://www.instagram.com/explore/tags/daily/...
	// count: 122907884
	// end_cursor: QVFCc1RXRXQyNDlaNVpjZGd5ejNUWEpyUWR4WW9LWlJocnNic2tKY1EyZnRKR3Q2bVRFRVFickZrTVluRXB0VVpDTUtkdDFwV3VXUGtPN0Rtb2RsckFOag==
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
	// Requesting to https://www.instagram.com/explore/tags/daily/...
	// count: 122893632
	// end_cursor: QVFEeFBIRW96NzEzSlJUc25GOE9UbktOUTR2aVFSMVByd1RJRFdqNGxOSFR3eFcxRF9jc3I3Q04tR0JoWklUMm0wV0didVpubzJXX0p5Rl9pUUdWV3ZxVA==
	// has_next_page: true
	//
	// Requesting to https://www.instagram.com/explore/tags/daily/?__a=1&max_id=QVFEeFBIRW96NzEzSlJUc25GOE9UbktOUTR2aVFSMVByd1RJRFdqNGxOSFR3eFcxRF9jc3I3Q04tR0JoWklUMm0wV0didVpubzJXX0p5Rl9pUUdWV3ZxVA==...
	// count: 122893634
	// end_cursor: QVFDZTBKWk96Wm1UVlU1REVWckhYcEdKREd4UTJWTmx0QmJOanp6dVg0NmhVUmljN0VGNlN4NU8wTlo3bkdQSThoeHRuUUxhU1drdzVNZ2tSajF6TmRSNQ==
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
	// Requesting to https://www.instagram.com/explore/tags/gophercon/...
	// Requesting to https://www.instagram.com/explore/tags/gophercon/?__a=1&max_id=QVFETzdZaGxkY3BXN2tzNExrRFlVbUtUdHZ3VE1ENmlLM1BRU1YyMmJNNUlKTzRnVkVLc1FmWUxDNldVemVNYXcycEVRVVFLWE9MZGQ3bXRIRUhnSW9TSQ==...
	// Requesting to https://www.instagram.com/explore/tags/gophercon/?__a=1&max_id=QVFDbVVQVDFqbWtIS2R5ejFZaG85ZzN4cFNOeFJIa2xjQWd1bk5OZlNqNzV2X1JMdk9zcW8xSUhGN1FOZHg0X0pyTFc3RUZqQllNejBEZVdINlN4Q2d5Tw==...
	// Requesting to https://www.instagram.com/explore/tags/gophercon/?__a=1&max_id=QVFBdWtDR3dScXV0LTlnaWxZRGxIb00xRWxqTzVlUzVWdThXNV9TOG9vUkoybWo1NUJJRkgzcEFGSVdoQUtON3A1TGtwa01JRURQbWpfaVlQaUdZVHpKaQ==...
	// Requesting to https://www.instagram.com/explore/tags/gophercon/?__a=1&max_id=QVFDVV9vX0gxSGZCczBpU292XzJzN3V0M19uNEZ4NHB1YzllWXk4SzRjajZRcl80eVBTeXM5VkFRTFQ5QmNZSmdKSjl2MlhpYjZ2cE4zZ3VQUmg5anl0UA==...
	// Requesting to https://www.instagram.com/explore/tags/gophercon/?__a=1&max_id=QVFDTFU0UV82c1h6Ykx3QW41dVdpZ21JVWw5ZHgwNDROUV94QldpT095Q3hpdlpjSUNaMWFTb3lkdVVSRXFPQVJCQWhSTDhqQnB2Nl92Z3U0ZG5aWERfNQ==...
	// Requesting to https://www.instagram.com/explore/tags/gophercon/?__a=1&max_id=QVFCNDZvYTY2LVhHaXZGM2xURENmdzFqZGVHLVdRSUVrNU1vb2xOYWNMcDUzMGM4MDZhOGlrTm1IZUZJUENubE9FVHZialZXek0zRTJDZ1JJd3ZnNk14QQ==...
	// Requesting to https://www.instagram.com/explore/tags/gophercon/?__a=1&max_id=QVFBWldHd0g3WGNrcDVmLWdVYmRwRXFJU2kyUC1RdnpibGNwYUttWE1obFp6dFRjZzlXUHpZMU9mMUNYOHVfRHZEOFRySU1jeEJlcXNHQkJtNHUxajdFSg==...
	// Requesting to https://www.instagram.com/explore/tags/gophercon/?__a=1&max_id=QVFDRFl0WDF3RTVUS0V2SEQtNHhlSXBUbHE0cy1QNHBOMjg2bEVNVVIwX0tTR0hxYWYyaC1sejY4dkV1SnEwN2p1S3lEdHdfN3d1NS00UllfU19ZUkVKUQ==...
	// count: 922
	// end_cursor:
	// has_next_page: false
}
