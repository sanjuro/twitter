package twitter

import (
	"testing"
)

func TestCheckFileTrue(t *testing.T) {
	if !CheckFile("tweet.txt") {
		t.Errorf("CheckFile was incorrect")
	}
}