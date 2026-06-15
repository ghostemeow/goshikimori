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

// Too many examples, take turns using.
func main() {
	c := config()
	o := &g.Options{Page: 1, Limit: 2}

	// PART 1
	a, status, err := c.SearchClubs("milf", o)
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	if len(a) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range a {
		fmt.Println(v.Id, v.Name, v.Is_censored)
	}

	// PART 2
	fast_spok, status, err := c.FastIdClub("Ahnenerbe")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}

	// Finding anime titles in the club.
	sca, status, err := fast_spok.SearchClubAnimes(o)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(sca) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range sca {
		fmt.Println(v.Id, v.Name, v.Score)
	}

	// Finding manga titles in the club.
	scm, status, err := fast_spok.SearchClubMangas(o)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(scm) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range scm {
		fmt.Println(v.Id, v.Name, v.Score)
	}

	// Finding ranobe titles in the club.
	scr, status, err := fast_spok.SearchClubRanobe(o)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(scr) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range scr {
		fmt.Println(v.Id, v.Name, v.Score)
	}

	// Club members.
	scmem, status, err := fast_spok.SearchClubMembers(o)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(scmem) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range scmem {
		fmt.Println(v.Id, v.Nickname, v.Last_online_at)
	}

	// Club characters.
	scc, status, err := fast_spok.SearchClubCharacters(o)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(scc) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range scc {
		fmt.Println(v.Id, v.Name, v.Russian)
	}

	fast_cude, status, err := c.FastIdClub("Кудере")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}

	// Club clubs.
	scl, status, err := fast_cude.SearchClubClubs(o)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(scl) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range scl {
		fmt.Println(v.Id, v.Name, v.Is_censored)
	}

	// Club images.
	sci, status, err := fast_cude.SearchClubImages(o)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(sci) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range sci {
		fmt.Println(v.Id, v.Original_url, v.Can_destroy, v.User_id)
	}

	// Club collections.
	sccl, status, err := fast_cude.SearchClubCollections(o)
	if status != 200 || err != nil {
		fmt.Println(err)
		return
	}
	if len(sccl) == 0 {
		fmt.Println("Club not found")
		return
	}
	for _, v := range sccl {
		fmt.Println(v.Body, v.Created_at, v.Comments_count)
	}

	// PART 3
	// Join/leave from club.
	fast_faq, status, err := c.FastIdClub("FAQ - Часто задаваемые вопросы")
	if status != 200 || err != nil {
		fmt.Println(status, err)
		return
	}
	cc, err := fast_faq.ClubJoin()
	//cc, err := fast_faq.ClubLeave()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cc)
}
