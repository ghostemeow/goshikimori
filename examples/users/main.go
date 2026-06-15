package main

import (
	"fmt"

	g "github.com/ghostemeow/goshikimori"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func main() {
	c := config()
	o := &g.Options{Page: 1, Limit: 10}
	u, status, err := c.SearchUsers("angel", o)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(u) == 0 {
		fmt.Println("users not found")
		return
	}
	for _, v := range u {
		fmt.Println(v.Id, v.Nickname, v.Last_online_at)
	}
}
