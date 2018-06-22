package twitter

import (
	"fmt"
)

type Tweet struct {
	Owner string
	Message string
}

func WriteTweet(tweet Tweet) {
	fmt.Println(fmt.Sprintf("\t@%s: %s", tweet.Owner, tweet.Message))
}