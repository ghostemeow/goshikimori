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
	s, status, err := c.SearchStudios()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(s) == 0 {
		fmt.Println("not found studios")
		return
	}
	for _, v := range s {
		fmt.Println(v.Id, v.Name, v.Filtered_name, v.Real)
	}
}
