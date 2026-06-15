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

	// related anime
	fast_anime, status, err := c.FastIdAnime("initial d second stage")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	a, status, err := fast_anime.SearchRelatedAnime()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(a) == 0 {
		fmt.Println("anime not found")
		return
	}
	for _, v := range a {
		fmt.Println(v.Relation, v.Relation_Russian, v.Anime.Score)
	}

	// related manga
	fast_manga, status, err := c.FastIdManga("vampire knight")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	m, status, err := fast_manga.SearchRelatedManga()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(m) == 0 {
		fmt.Println("Manga not found")
		return
	}
	for _, v := range m {
		fmt.Println(
			v.Relation, v.Relation_Russian,
			v.Manga.Score, v.Manga.Status,
		)
	}

	// related ranobe
	fast_ranobe, status, err := c.FastIdRanobe("sword art")
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	r, status, err := fast_ranobe.SearchRelatedRanobe()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(r) == 0 {
		fmt.Println("Ranobe not found")
		return
	}
	for _, v := range r {
		fmt.Println(
			v.Relation, v.Relation_Russian,
			v.Manga.Score, v.Manga.Status,
		)
	}
}
