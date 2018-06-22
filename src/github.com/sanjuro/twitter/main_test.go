package main

import (
	"testing"
)

func TestCheckFileTrue(t *testing.T) {
	if !checkFile("tweet.txt") {
		t.Errorf("checkFile was incorrect")
	}
}

func TestCheckFileFalse(t *testing.T) {
	if !checkFile("user.txt") {
		t.Errorf("checkFile was incorrect")
	}
}