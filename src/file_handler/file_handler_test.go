package file_handler

import (
	"testing"
)

func TestCheckFileTrue(t *testing.T) {
	if !CheckFile("tweet.txt") {
		t.Errorf("checkFile was incorrect")
	}
}

func TestCheckFileFalse(t *testing.T) {
	if !CheckFile("user.txt") {
		t.Errorf("checkFile was incorrect")
	}
}