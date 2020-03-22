package crawler

import (
	"fmt"
)

func ExampleCrawler_Init() {
	c := New()
	if err := c.Init("daily"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c)
	// Output:
	//
}
