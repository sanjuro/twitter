package main

import (
	"fmt"
)

type TwitterUser struct {
	name string
	followers string
}

func writeUser(user TwitterUser) {
	fmt.Println(user.name)
}

type Tweet struct {
	owner string
	message string
}

func writeTweet(tweet Tweet) {
	fmt.Println(fmt.Sprintf("\t@%s: %s", tweet.owner, tweet.message))
}