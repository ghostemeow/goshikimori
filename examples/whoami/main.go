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
	w, status, err := c.WhoAmi()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fmt.Println(w.Nickname, w.Avatar, w.Locale, w.Last_online_at)
}
