package twitter

import (
	"fmt"
)

type Tweet struct {
	Owner string
	Message string
}

func WriteTweet(tweet Tweet) {
	if tweet != nil {
		if tweet.Owner != nil && tweet.Message) != nill {
			fmt.Println(fmt.Sprintf("\t@%s: %s", tweet.Owner, tweet.Message))
		}
	}
}