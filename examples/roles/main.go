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
	// Anime roles.
	fast_anime, status, err := c.FastIdAnime("naruto")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	a, status, err := fast_anime.SearchAnimeRoles()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(a) == 0 {
		fmt.Println("anime not found")
		return
	}
	for _, v := range a {
		fmt.Println(
			v.Roles, v.Roles_Russian,
			v.Character.Id, v.Character.Name,
		)
	}
	// Manga roles.
	fast_manga, status, err := c.FastIdManga("naruto")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	m, status, err := fast_manga.SearchMangaRoles()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(m) == 0 {
		fmt.Println("manga not found")
		return
	}
	for _, v := range m {
		fmt.Println(
			v.Roles, v.Roles_Russian,
			v.Character.Id, v.Character.Name,
		)
	}
	// Ranobe roles.
	fast_ranobe, status, err := c.FastIdRanobe("sword art")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	r, status, err := fast_ranobe.SearchRanobeRoles()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(r) == 0 {
		fmt.Println("ranobe not found")
		return
	}
	for _, v := range r {
		fmt.Println(
			v.Roles, v.Roles_Russian,
			v.Character.Id, v.Character.Name,
		)
	}
}
