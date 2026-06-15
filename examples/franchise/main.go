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

	// franchise anime.
	fast_anime, status, err := c.FastIdAnime("initial d")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fa, status, err := fast_anime.SearchAnimeFranchise()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(fa.Nodes) == 0 {
		fmt.Println("anime franchise not found")
		return
	}
	for _, v := range fa.Nodes {
		fmt.Println(v.Id, v.Name, v.Kind)
	}

	// franchise manga.
	fast_manga, status, err := c.FastIdManga("naruto")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fm, status, err := fast_manga.SearchMangaFranchise()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(fm.Nodes) == 0 {
		fmt.Println("manga franchise not found")
		return
	}
	for _, v := range fm.Nodes {
		fmt.Println(v.Id, v.Name, v.Kind)
	}

	// franchise ranobe.
	fast_ranobe, status, err := c.FastIdRanobe("sword art")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	fr, status, err := fast_ranobe.SearchRanobeFranchise()
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(fr.Nodes) == 0 {
		fmt.Println("ranobe francise not found")
		return
	}
	for _, v := range fr.Nodes {
		fmt.Println(v.Id, v.Name, v.Kind)
	}
}
