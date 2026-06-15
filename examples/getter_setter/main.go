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

	fast, status, err := c.FastIdUser("arctica")
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	brief, status, err := fast.UserBriefInfo()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(brief.Id, brief.Nickname)

	// Getting an id.
	fmt.Println(fast.Id) // fmt.Println(fast.GetFastId())
	// Quick id change.
	new_fast := c.SetFastId(1) // fmt.Println(new_fast.Id)
	new_brief, status, err := new_fast.UserBriefInfo()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(new_brief.Id, new_brief.Nickname)

	// Getting configuration.
	fmt.Println(c.GetConfiguration())
	// Quick configuration change.
	new_config := g.SetConfiguration("Bob", "XXX-XXX-XXX")
	fmt.Println(new_config.Application, new_config.AccessToken)
}
