package main

import (
	"fmt"

	g "github.com/ghostemeow/goshikimori"
	"github.com/ghostemeow/goshikimori/constants"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func main() {
	c := config()
	// show user messages
	fast, status, err := c.FastIdUser("arctica")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	m, status, err := fast.UserMessages(&g.Options{
		Type: constants.MESSAGE_TYPE_NEWS, Page: 1, Limit: 10,
	})
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(m) == 0 {
		fmt.Println("messages not found")
		return
	}
	for _, v := range m {
		fmt.Println(v.Id, v.Kind, v.HTMLBody, v.Created_at)
	}
	// show user unread messages counts
	um, status, err := fast.UserUnreadMessages()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(um.Messages, um.News, um.Notifications)
}
