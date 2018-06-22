package twitter

import (
	"fmt"
)

type TwitterUser struct {
	Name string
	Followers string
}

func WriteUser(user TwitterUser) {
	fmt.Println(user.Name)
}