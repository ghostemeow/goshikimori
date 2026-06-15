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
	fast, status, err := c.FastIdAnime("initial d first stage")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	w, status, err := fast.SearchAnimeVideos()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(w) == 0 {
		fmt.Println("Video not found")
		return
	}
	for _, v := range w {
		fmt.Println(v.Id, v.Url, v.Image_url, v.Player_url, v.Name, v.Kind, v.Hosting)
	}
}
