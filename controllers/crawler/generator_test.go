// Package crawler implements Instagram crawling logic.
package crawler

import "fmt"

func ExampleGenerator() {
	crawler := Generator("cafe", "")
	for i := 0; i < 3; i++ {
		posts, endCursor, err := crawler()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(len(posts))
		fmt.Println(endCursor)
	}
	// Output:
	// 71
	// QVFDdnlxSHl3RDZiNGRya08zbElfM3NzUGwxbFhVM2N4Z05fTHd1bG90cVM2OTZaSTBIa1Z0VzlGM3Bya3RqUG11Tk91RFA5WlE0RDIyUEJpOTYtdVRIZg==
	// 72
	// QVFDS29fY19VZG8tTVY2cFNSTkZfOTczTnctTHZzd0dNcFdKVERXUy1MV0pjRHp5ZkdVa1VyUWFyNmhpNDZKcjRXSE1EMU1xd3lwaDFITndLYnJXV2tyQg==
	// 72
	// QVFDZjFtTTJKM2dELVdpZnVHbzlMRmxpRXZkNVhPajZXX3FOajF3b2JENDFlN3JkWW1kOWQ4SloyNkw1OXI1cHRZei1OUEE5SnRuRkpPWWIwSkl4N2lrWQ==
}
