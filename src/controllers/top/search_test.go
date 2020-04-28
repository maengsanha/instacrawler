// Package top implements the Instagram top search crawling logic.
package top

import "fmt"

func ExampleSearch() {
	hashtags, err := Search("cafe")
	if err != nil {
		panic(err)
	}
	for _, hashtag := range hashtags {
		fmt.Println(hashtag)
	}
	// Output:
	// #café 7m posts 7080757
	// #cafe 60.5m posts 60537500
	// #cafè 205k posts 205626
	// #cafebkk 200k posts 200862
	// #cafeteller 267k posts 267979
	// #cafeinterior 137k posts 137350
	// #cafegourmet 131k posts 131814
	// #cafehopper 184k posts 184403
	// #caferacers 510k posts 510920
	// #cafetour 289k posts 289339
	// #cafeteria 1.6m posts 1644158
	// #cafelover 326k posts 326270
	// #cafejakarta 267k posts 267971
	// #cafegram 204k posts 204367
	// #cafesurabaya 181k posts 181465
	// #cafédatarde 130k posts 130106
	// #cafehop 187k posts 187801
	// #cafebandung 338k posts 338175
	// #cafedesign 227k posts 227712
	// #cafehopmy 299k posts 299089
}
