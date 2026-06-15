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
	b, status, err := c.SearchBans()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range b {
		fmt.Println(v.Id, v.User.Id, v.User.Nickname)
	}
}
