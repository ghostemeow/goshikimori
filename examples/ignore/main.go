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
	ignore, status, err := fast.AddIgnoreUser()
	//ignore, status, err := fast.RemoveIgnoreUser()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ignore.User_id, ignore.Is_ignored)
}
