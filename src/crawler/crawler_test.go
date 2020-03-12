package crawler

import (
	"fmt"
	"testing"
)

func TestCrawlNumCoreSpriteHashtag(t *testing.T) {
	if CrawlNumCoreSpriteHashtag("daily") != 122204072 {
		t.Errorf("Expected 122204072, but got %d\n", CrawlNumCoreSpriteHashtag("daily"))
	} else {
		t.Log("Test Succeded.")
	}
}

func ExampleParsePostURL() {
	postURLs := ParsePostURL("daily")

	fmt.Printf("Got %d postings.\n", len(postURLs))

	for _, url := range postURLs {
		fmt.Println(url)
	}
	// Output:
	// Got 81 postings.
	// https://instagram.com/p/B9pAFuolqAn
	// https://instagram.com/p/B9pAFKmH679
	// https://instagram.com/p/B9pAE4fAVTA
	// https://instagram.com/p/B9pAE0bF8Ok
	// https://instagram.com/p/B9pAEjahnRj
	// https://instagram.com/p/B9pAEG_gdQv
	// https://instagram.com/p/B9pAD_4l9gj
	// https://instagram.com/p/B9pADfJgvPS
	// https://instagram.com/p/B9pAC_cpfP3
	// https://instagram.com/p/B9pACqaBUrr
	// https://instagram.com/p/B9pACQpBhr5
	// https://instagram.com/p/B9pAB9RFki6
	// https://instagram.com/p/B9pABP1BbPR
	// https://instagram.com/p/B9o__K1leGb
	// https://instagram.com/p/B9o__J1l1gy
	// https://instagram.com/p/B9o_-ptF5De
	// https://instagram.com/p/B9o_9fEnIgh
	// https://instagram.com/p/B9o_wypBAPa
	// https://instagram.com/p/B9o_Y1nhSbS
	// https://instagram.com/p/B9o_Lq-F__m
	// https://instagram.com/p/B9o_FpRlFZO
	// https://instagram.com/p/B9o_CmYJMCF
	// https://instagram.com/p/B9o_A8vFdFR
	// https://instagram.com/p/B9o-4uLhQSI
	// https://instagram.com/p/B9o-jVuhvOY
	// https://instagram.com/p/B9o-lflj2L4
	// https://instagram.com/p/B9o-cywHCo7
	// https://instagram.com/p/B9o-bQ0FhX8
	// https://instagram.com/p/B9o-VNDD882
	// https://instagram.com/p/B9o-IVoF48G
	// https://instagram.com/p/B9o9qU9h674
	// https://instagram.com/p/B9o9XEmA4hZ
	// https://instagram.com/p/B9o8jLnHtin
	// https://instagram.com/p/B9o7lUSFZ7i
	// https://instagram.com/p/B9o7hRSlZOw
	// https://instagram.com/p/B9o7GaPAdoY
	// https://instagram.com/p/B9o6MGpH9JP
	// https://instagram.com/p/B9o53j_ga70
	// https://instagram.com/p/B9o5fEqlixE
	// https://instagram.com/p/B9o5CR3lLE2
	// https://instagram.com/p/B9o4984lmqJ
	// https://instagram.com/p/B9o35UjgffA
	// https://instagram.com/p/B9o2xMflQZn
	// https://instagram.com/p/B9o2quJhtd5
	// https://instagram.com/p/B9o2h6rnhhl
	// https://instagram.com/p/B9o0mZNAe3c
	// https://instagram.com/p/B9ozRU2lon2
	// https://instagram.com/p/B9oyB8lBJv-
	// https://instagram.com/p/B9oxyAYFPpz
	// https://instagram.com/p/B9oxKC_huxd
	// https://instagram.com/p/B9owBhkDWxE
	// https://instagram.com/p/B9ouSNRA7Zu
	// https://instagram.com/p/B9oqRdqAmt3
	// https://instagram.com/p/B9omx3zFEmC
	// https://instagram.com/p/B9olqwSlb5E
	// https://instagram.com/p/B9oh8pcFBqq
	// https://instagram.com/p/B9oe4wZhLIF
	// https://instagram.com/p/B9oY-5XFlm0
	// https://instagram.com/p/B9oVl7PF1eC
	// https://instagram.com/p/B9oNetyFzKn
	// https://instagram.com/p/B9n9C7JlxIO
	// https://instagram.com/p/B9n2gMvp8kR
	// https://instagram.com/p/B9npu6QlbDo
	// https://instagram.com/p/B9mCrsLhhFP
	// https://instagram.com/p/B9k1GjrpVdH
	// https://instagram.com/p/B9kIj3KJvLp
	// https://instagram.com/p/B9jtG_NhGrO
	// https://instagram.com/p/B9eAu9xhymK
	// https://instagram.com/p/B9c86yEhYxR
	// https://instagram.com/p/B9PpxEQAAhQ
	// https://instagram.com/p/B7lawbvjo-6
	// https://instagram.com/p/B7ggTlgjAhd
	// https://instagram.com/p/B9oukjqhQjO
	// https://instagram.com/p/B9oyv09lttR
	// https://instagram.com/p/B9oyJ66BA4q
	// https://instagram.com/p/B9ov1sMlbn2
	// https://instagram.com/p/B9owDfbFA6C
	// https://instagram.com/p/B9o0XVRHgDd
	// https://instagram.com/p/B9o0VwDjdz4
	// https://instagram.com/p/B9oypejnr5M
	// https://instagram.com/p/B9oyJ6MgGpf
}

func ExampleCoreSpriteHashtag() {
	CoreSpriteHashtag("https://www.instagram.com/p/B9orl1xHH9c/")
	// Output:
	//
}
