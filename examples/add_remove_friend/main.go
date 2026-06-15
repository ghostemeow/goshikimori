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
	fast, status, err := c.FastIdUser("morr")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fr, status, err := fast.AddFriend()
	//fr, status, err := fast.RemoveFriend()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fr.Notice)
}
