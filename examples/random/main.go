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

	a, status, err := c.RandomAnimes(&g.Options{
		Limit: 5, Score: 5, Kind: constants.ANIME_KIND_TV,
		Status: constants.ANIME_STATUS_RELEASED,
		Censored: false, Genre_v2: []int{10, 539},
	})
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range a {
		fmt.Println(
			v.Id, v.Name, v.Score,
			v.Status, v.Released_on, v.Url,
		)
	}

	m, status, err := c.RandomMangas(&g.Options{
		Limit: 5, Score: 8, Kind: constants.MANGA_KIND_MANGA,
		Status: constants.MANGA_STATUS_RELEASED, Censored: false,
	})
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range m {
		fmt.Println(
			v.Id, v.Name, v.Score, v.Chapters,
			v.Volumes, v.Released_on, v.Url,
		)
	}

	r, status, err := c.RandomRanobes(&g.Options{
		Limit: 5, Score: 8, Status: constants.MANGA_STATUS_RELEASED, Censored: false,
	})
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range r {
		fmt.Println(
			v.Id, v.Name, v.Score, v.Chapters,
			v.Volumes, v.Released_on, v.Url,
		)
	}
}
