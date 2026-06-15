package main

import (
	"fmt"

	g "github.com/ghostemeow/goshikimori"
	"github.com/ghostemeow/goshikimori/constants"
	"github.com/ghostemeow/goshikimori/genres"
)

func config() *g.Configuration {
	return g.SetConfiguration(
		"APPLICATION_NAME",
		"PRIVATE_KEY",
	)
}

func main() {
	c := config()
	gnrs, status, err := c.SearchGenres(constants.GENRES_ANIME)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(gnrs) == 0 {
		fmt.Println("not found genres")
		return
	}
	for _, v := range gnrs {
		fmt.Println(v.Id, v.Name, v.Russian, v.Kind, v.Entry_type)
	}
	// A small map helper.
	m := genres.GenerateGenres(constants.GENERATE_GENRES_ANIME, gnrs)
	fmt.Println(m)
}
