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

	// Similar anime.
	fast_anime, status, err := c.FastIdAnime("vampire knight")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	a, status, err := fast_anime.SearchSimilarAnime()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(a) == 0 {
		fmt.Println("anime not found")
		return
	}
	for _, v := range a {
		fmt.Println(v.Name, v.Id, v.Russian)
	}

	// Similar manga.
	fast_manga, status, err := c.FastIdManga("initial d")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	m, status, err := fast_manga.SearchSimilarManga()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(m) == 0 {
		fmt.Println("manga not found")
		return
	}
	for _, v := range m {
		fmt.Println(v.Name, v.Id, v.Russian)
	}

	// Similar ranobe.
	fast_ranobe, status, err := c.FastIdRanobe("sword art")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	r, status, err := fast_ranobe.SearchSimilarRanobe()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(r) == 0 {
		fmt.Println("ranobe not found")
		return
	}
	for _, v := range r {
		fmt.Println(v.Id, v.Name, v.Score, v.Volumes, v.Chapters)
	}
}
