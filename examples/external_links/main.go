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
	// Anime external links.
	fast_anime, status, err := c.FastIdAnime("initial d")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fra, status, err := fast_anime.SearchAnimeExternalLinks()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(fra) == 0 {
		fmt.Println("external links not found")
		return
	}
	for _, v := range fra {
		fmt.Println(v.Id, v.Kind, v.Url, v.Source, v.Entry_type)
	}
	// Manga external links.
	fast_manga, status, err := c.FastIdManga("initial d")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	frm, status, err := fast_manga.SearchMangaExternalLinks()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(frm) == 0 {
		fmt.Println("external links not found")
		return
	}
	for _, v := range frm {
		fmt.Println(v.Id, v.Kind, v.Url, v.Source, v.Entry_type)
	}
	// Ranobe external links.
	fast_ranobe, status, err := c.FastIdRanobe("Ookami to Koushinryou")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	frr, status, err := fast_ranobe.SearchRanobeExternalLinks()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(frr) == 0 {
		fmt.Println("external links not found")
		return
	}
	for _, v := range frr {
		fmt.Println(v.Id, v.Kind, v.Url, v.Source, v.Entry_type)
	}
}
