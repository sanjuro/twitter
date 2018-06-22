package twitter

import (
	"fmt"
)

type TwitterUser struct {
	Name string
	Followers string
}

func WriteUser(user TwitterUser) {
	if user != nil && user.Name != nill{
		fmt.Println(user.Name)
	}
}