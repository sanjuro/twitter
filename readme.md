# Requirements

Please write a program to simulate a twitter-like feed. Your program will receive two seven-bit ASCII files. The first file contains a list of users and their followers. The second file contains tweets. Given the users, followers and tweets, the objective is to display a simulated twitter feed for each user to the console. 

The program should be well designed, handle errors and should be of sufficient quality to run on a production system. Indicate all assumptions made in completing the assignment. Please email the source code along with related build files.

Each line of a well-formed user file contains a user, followed by the word ‘follows’ and then a comma separated list of users they follow.  Where there is more than one entry for a user,  consider the union of all these entries to determine the users they follow.

Lines of the tweet file contain a user, followed by greater than, space and then a tweet that may be at most 140 characters in length. The tweets are considered to be posted by the each user in the order they are found in this file.

Your program needs to write console output as follows. For each user / follower (in alphabetical order) output their name on a line. Then for each tweet, emit a line with the following format: <tab>@user: <space>message.

Here is an example. Given user file named user.txt:

```
Ward follows Alan
Alan follows Martin
Ward follows Martin, Alan
```

And given tweet file named tweet.txt:
```
Alan> If you have a procedure with 10 parameters, you probably missed some.
Ward> There are only two hard things in Computer Science: cache invalidation, naming things and off-by-1 errors.
Alan>Random numbers should not be generated with a method chosen at random.
```
Invoking your program with user.txt and tweet.txt as arguments should produce the following console output:

```
Alan
    @Alan: If you have a procedure with 10 parameters, you probably missed some.
    @Alan: Random numbers should not be generated with a method chosen at random.
Martin
Ward
    @Alan: If you have a procedure with 10 parameters, you probably missed some.
    @Ward:There are only two hard things in Computer Science: cache invalidation, naming things and off-by-1 errors.
    @Alan: Random numbers should not be generated with a method chosen at random.
```

# My Assumptions

+ Based on the  file content, output listed in the *Requirements* section and the stated goals the actual tweet file should be as below:
```
Alan> If you have a procedure with 10 parameters, you probably missed some.
Ward> There are only two hard things in Computer Science: cache invalidation, naming things and off-by-1 errors.
Alan> Random numbers should not be generated with a method chosen at random.
```
and the an expected output would look like:
```
Alan
    @Alan: If you have a procedure with 10 parameters, you probably missed some.
    @Alan: Random numbers should not be generated with a method chosen at random.
Martin
Ward
    @Alan: If you have a procedure with 10 parameters, you probably missed some.
    @Ward: There are only two hard things in Computer Science: cache invalidation, naming things and off-by-1 errors.
    @Alan: Random numbers should not be generated with a method chosen at random.
```
+ Users defined in the user file will might have or not have followers.
+ User account names in the system are unique.
+ As the program stores the users and tweets in memory, there will an upper limit based on memory to the users and tweets to process.
+ If there is a tweet from a user that is not present in the user file, it will not be displayed.

# What Software do you reqiure

1. GoLang version go1.10.2 darwin/amd64 and $GO_PATH set correctly
2. Git 1.9.1
3. A linux/MacOS X system.

# How do I build the project

Below are the instructions to checkout and build the assessment project.

1. git clone this repo, this should create a twitte folder.->  
    -> git clone git@github.com:sanjuro/twitter.git /Users/user/Development/Sourcecode/go
2. cd to -> (Go path ) go/src/feed
3. run "go install", there should build the feed file in the bin folder
4. run "go build"

# Running the application

1. To run application type:  feed  -users=user.txt -tweets=tweet.txt
2.  ➜  cd src/feed 
    ➜  feed go test -cover -v  
