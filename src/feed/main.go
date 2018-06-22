package main

import (
	"fmt"
	"flag"
	"sort"
	"strings"

	"file_handler"
	"twitter"
)


func main() {
	// Get inputs
	userTextFile := flag.String("users", "", "a file with all the users")
	tweetTextFile := flag.String("tweets", "", "a file with all the tweets")

	flag.Parse()

	var users = make(map[string]twitter.TwitterUser)
	var tweets = []twitter.Tweet{}
	var order []string
	
	if *userTextFile != "" {

		file_handler.CheckFile(*userTextFile)

		usersLines, err := file_handler.ReadLines(*userTextFile)

		if err != nil {
			panic(err)
		}

		for _, usersLine := range usersLines {

			if *tweetTextFile == "" {
				fmt.Println(usersLine)
			}

			usersSplit := strings.Split(usersLine, " follows ")

			var user twitter.TwitterUser
			user.Name = usersSplit[0]

			user.Followers = usersSplit[1]
			following := strings.Split(usersSplit[1], ",")

			for _, follow := range following {
				followName := strings.Trim(follow, " ")
				var follower twitter.TwitterUser
				follower.Name = followName
				users[follower.Name] = follower
			}

			users[user.Name] = user
		}

		for _, user := range users {
			order = append(order, user.Name)
			sort.Strings(order) 
		}
	}

	if *tweetTextFile != "" {

		file_handler.CheckFile(*tweetTextFile)

		tweetsLines, err := file_handler.ReadLines(*tweetTextFile)

		if err != nil {
			panic(err)
		}

		for _, tweetsLine := range tweetsLines {

			if *userTextFile == "" {
				fmt.Println(tweetsLine)
			}
			
			tweetData := strings.Split(tweetsLine, "> ")

			var tweet twitter.Tweet
			tweet.Owner = tweetData[0]
			tweet.Message = tweetData[1]
			tweets = append(tweets, tweet)
		}
	}

	// Create feed
	if *tweetTextFile != "" && *userTextFile != "" {
		for _, key := range order {
			var user twitter.TwitterUser
			user = users[key]
			twitter.WriteUser(user)

			for _, tweet := range tweets {
				if user.Name == tweet.Owner || strings.Contains(user.Followers, tweet.Owner) {
					twitter.WriteTweet(tweet)
				}
			}
		}
	}
}
