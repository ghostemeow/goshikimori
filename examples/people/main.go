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
	// Search peoples.
	o := &g.Options{Kind: constants.PEOPLE_KIND_SEYU}
	sp, status, err := c.SearchPeoples("Aya Hirano", o)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	for _, v := range sp {
		fmt.Println(v.Id, v.Name, v.Russian, v.Image.Original)
	}
	// Search people.
	fast, status, err := c.FastIdPeople("Aya Hirano")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	p, status, err := fast.SearchPeople()
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if p.Id == 0 {
		fmt.Println("people not found")
		return
	}
	fmt.Println(
		p.Id, p.Name, p.Japanese, p.Job_title, p.Website,
		p.Birth_on.Day, p.Birth_on.Month, p.Birth_on.Year,
	)
	for _, v := range p.Groupped_roles {
		fmt.Println(v[0], v[1])
	}
	for _, v := range p.Roles {
		for _, vv := range v.Characters {
			fmt.Println(vv.Id, vv.Name)
		}
	}
	for _, v := range p.Roles {
		for _, vv := range v.Animes {
			fmt.Println(vv.Id, vv.Name, vv.Score)
		}
	}
	for _, v := range p.Works {
		fmt.Println(v.Anime.Id, v.Anime.Name, v.Anime.Score)
	}
}
