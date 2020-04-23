// Package top implements the top search crawling logic on Instagram.
package top

import "fmt"

func ExampleSearch() {
	hashtags, err := Search("kakao")
	if err != nil {
		panic(err)
	}
	for _, hashtag := range hashtags {
		fmt.Println(hashtag)
	}
	// Output:
	// #kakao 674k posts 674895
	// #kakaofriends 501k posts 501300
}
