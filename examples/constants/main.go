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
	ca, status, err := c.SearchConstantsAnime()
	if err != nil {
		fmt.Println(err)
		return
	}
	if status == 200 {
		fmt.Println(ca.Kind, ca.Status)
	}

	cm, status, err := c.SearchConstantsManga()
	if err != nil {
		fmt.Println(err)
		return
	}
	if status == 200 {
		fmt.Println(cm.Kind, cm.Status)
	}

	ur, status, err := c.SearchConstantsUserRate()
	if err != nil {
		fmt.Println(err)
		return
	}
	if status == 200 {
		fmt.Println(ur.Status)
	}

	cc, status, err := c.SearchConstantsClub()
	if err != nil {
		fmt.Println(err)
		return
	}
	if status == 200 {
		fmt.Println(cc.Join_policy, cc.Comment_policy, cc.Image_upload_policy)
	}

	cs, status, err := c.SearchConstantsSmileys()
	if err != nil {
		fmt.Println(err)
		return
	}
	if status == 200 {
		for _, v := range cs {
			fmt.Println(v.Bbcode, v.Path)
		}
	}
}
