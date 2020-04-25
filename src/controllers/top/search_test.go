// Package top implements the Instagram top search crawling logic.
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
	// #kakao 675k posts 675332
	// #kakaofriends 501k posts 501545
	// #kakaolukek 24.8k posts 24822
	// #kakaotalk 146k posts 146284
}
