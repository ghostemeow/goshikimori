package main

import (
	"fmt"

	g "github.com/ghostemeow/goshikimori"
	"github.com/ghostemeow/goshikimori/neko"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func main() {
	c := config()
	fast, status, err := c.FastIdUser("incarnati0n")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	a, status, err := fast.SearchAchievement()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(a) == 0 {
		fmt.Println("achievements not found")
		return
	}
	neko, err := neko.Search("Hellsing")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range a {
		if v.Neko_id == neko {
			fmt.Printf("level: %d - progress %d\n", v.Level, v.Progress)
			fmt.Printf("created: %v - updated: %v\n", v.Created_at, v.Updated_at)
		}
	}
}
