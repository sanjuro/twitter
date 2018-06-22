package file_handler

import (
	"testing"
)

func TestCheckFileTrue(t *testing.T) {
	if !CheckFile("tweet.txt") {
		t.Errorf("CheckFile was incorrect")
	}
}

func TestCheckFileFalse(t *testing.T) {
	if !CheckFile("user.txt") {
		t.Errorf("CheckFile was incorrect")
	}
}

func TestReadLinesFalse(t *testing.T) {
	tweetsLines, err := ReadLines("empty.txt")
	
	if err != nil {
		t.Errorf("ReadLines was incorrect")
	}

	if tweetsLines != nil {
		t.Errorf("ReadLines was incorrect")
	}
}