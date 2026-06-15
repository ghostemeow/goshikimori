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
	p, status, err := c.SearchPublishers()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(p) == 0 {
		fmt.Println("not found publishers")
		return
	}
	for _, v := range p {
		fmt.Println(v.Id, v.Name)
	}
}
