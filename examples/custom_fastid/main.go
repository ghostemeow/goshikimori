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

	// EXAMPLE #1.
	// We can already know our ID via WhoAmi() and
	// without unnecessary searches we get access to additional functions,
	// for example SearchUserFriends().
	me, _, err := c.WhoAmi()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Override the ID.
	custom_fast := c.SetFastId(me.Id)
	// Find the user's friends.
	friends, status, err := custom_fast.SearchUserFriends(&g.Options{Page: 1, Limit: 5})
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range friends {
		fmt.Println(v.Id, v.Nickname, v.Last_online_at)
	}

	// EXAMPLE #2.
	// When searching by NICKNAME go straight to a specific user,
	// you can dilute with all sorts of conditions.
	users, status, err := c.SearchUsers("morr", &g.Options{Page: 1, Limit: 10})
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range users {
		// Old user search filter.
		if v.Id < 5 {
			// Override the ID.
			custom_fast := c.SetFastId(v.Id)
			// Search for completed anime from the user.
			anime_rates, status, err := custom_fast.SearchUserAnimeRates(&g.Options{
				Page: 1, Limit: 50, Status: "completed",
			})
			if status != 200 || err != nil {
				fmt.Println(status, err)
				return
			}
			for _, v := range anime_rates {
				fmt.Println(v.Id, v.Anime.Name, v.Score, v.Text, v.Created_at, v.Updated_at)
			}
		}
	}
}
