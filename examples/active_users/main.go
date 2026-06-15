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
	result, status, err := c.ActiveUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	if status == 200 {
		fmt.Println(result)
	}
}
