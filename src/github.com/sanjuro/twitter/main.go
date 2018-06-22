package main

import (
	"fmt"
	"flag"
	"sort"
	"strings"
)


func main() {
	// Get inputs
	userTextFile := flag.String("users", "", "a file with all the users")
	tweetTextFile := flag.String("tweets", "", "a file with all the tweets")

	flag.Parse()

	var users = make(map[string]TwitterUser)
	var tweets = []Tweet{}
	var order []string
	
	if *userTextFile != "" {

		checkFile(*userTextFile)

		usersLines, err := readLines(*userTextFile)

		if err != nil {
			panic(err)
		}

		for _, usersLine := range usersLines {

			if *tweetTextFile == "" {
				fmt.Println(usersLine)
			}

			usersSplit := strings.Split(usersLine, " follows ")

			var user TwitterUser
			user.name = usersSplit[0]

			user.followers = usersSplit[1]
			following := strings.Split(usersSplit[1], ",")

			for _, follow := range following {
				followName := strings.Trim(follow, " ")
				var follower TwitterUser
				follower.name = followName
				users[follower.name] = follower
			}

			users[user.name] = user
		}

		for _, user := range users {
			order = append(order, user.name)
			sort.Strings(order) 
		}
	}

	if *tweetTextFile != "" {

		checkFile(*tweetTextFile)

		tweetsLines, err := readLines(*tweetTextFile)

		if err != nil {
			panic(err)
		}

		for _, tweetsLine := range tweetsLines {

			if *userTextFile == "" {
				fmt.Println(tweetsLine)
			}
			
			tweetData := strings.Split(tweetsLine, "> ")

			var tweet Tweet
			tweet.owner = tweetData[0]
			tweet.message = tweetData[1]
			tweets = append(tweets, tweet)
		}
	}

	// Create feed
	if *tweetTextFile != "" && *userTextFile != "" {
		for _, key := range order {
			user := users[key]
			writeUser(user)

			for _, tweet := range tweets {
				if user.name == tweet.owner || strings.Contains(user.followers, tweet.owner) {
					writeTweet(tweet)
				}
			}
		}
	}
}
