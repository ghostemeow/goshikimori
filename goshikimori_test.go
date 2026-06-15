package goshikimori

import (
	"os"
	"testing"

	"github.com/ghostemeow/goshikimori/constants"
	"github.com/ghostemeow/goshikimori/internal/utils"
	"github.com/ghostemeow/goshikimori/neko"
)

const app_test = ""

func conf() *Configuration { return SetConfiguration(app_test, "") }

func TestConfiguration(t *testing.T) {
	if app_test != "" {
		t.Logf("Found: %s", app_test)
	} else {
		t.Error("Not found application or token")
		os.Exit(1)
	}
}

func TestUser(t *testing.T) {
	c := conf()
	s, _, _ := c.SearchUser("arctica")

	if s.Id == 181833 && s.Sex == "male" {
		t.Logf("User: %s, Id: %d - found", s.Nickname, s.Id)
	} else {
		t.Error("User: arctica - not found")
	}
}

func TestAnimes(t *testing.T) {
	c := conf()
	o := &Options{
		Page: 1, Limit: 1, Score: 1,
		Censored: false,
	}
	s, _, _ := c.SearchAnimes("Initial D", o)

	for _, v := range s {
		if v.Id == 12725 && v.Status == "released" {
			t.Logf("Anime: %s, Id: %d - found", v.Name, v.Id)
		} else {
			t.Errorf("Anime: %s, Id: %d - not found", v.Name, v.Id)
		}
	}
}

func TestMangas(t *testing.T) {
	c := conf()
	o := &Options{Page: 1, Limit: 1}
	r, _, _ := c.SearchMangas("Initial D", o)

	for _, v := range r {
		if v.Volumes == 48 && v.Chapters == 724 {
			t.Logf("Manga: %s, Id: %d - found", v.Name, v.Id)
		} else {
			t.Errorf("Manga: %s, Id: %d - not found", v.Name, v.Id)
		}
	}
}

func TestClubs(t *testing.T) {
	c := conf()
	o := &Options{Page: 1, Limit: 1}
	clubs, _, _ := c.SearchClubs("milf", o)

	for _, v := range clubs {
		if v.Is_censored == true {
			t.Logf("Best club: %s - found", v.Name)
		} else {
			t.Errorf("Argument: %v or Id: %d - not found", v.Is_censored, v.Id)
		}
	}
}

func TestAchievements(t *testing.T) {
	var s utils.StatusBar
	s.Settings(5, "#", 1)
	s.Run()

	c := conf()
	fast, _, _ := c.FastIdUser("arctica")
	u, _, _ := fast.SearchAchievement()
	neko, _ := neko.Search("Hellsing")

	for _, v := range u {
		if v.Neko_id == neko {
			if v.Progress == 100 {
				t.Logf("Found: %d progress", v.Progress)
			} else {
				t.Error("Progress not found")
			}
		}
	}
}

func TestAnimeVideos(t *testing.T) {
	c := conf()
	fast, _, _ := c.FastIdAnime("initial d first stage")
	a, _, _ := fast.SearchAnimeVideos()

	for _, v := range a {
		if v.Id == 24085 {
			t.Logf("Video: %s", v.Name)
		} else {
			t.Log("Video not found, waiting...")
		}
	}
}

func TestUserUnreadMessages(t *testing.T) {
	var s utils.StatusBar
	s.Settings(5, "#", 1)
	s.Run()

	c := conf()
	fast, _, _ := c.FastIdUser("arctica")
	um, _, _ := fast.UserUnreadMessages()

	if um.News > 0 || um.News == 0 {
		t.Logf("Found: %d news", um.News)
	} else {
		t.Error("News not found")
	}
}

func TestConstantsAnime(t *testing.T) {
	c := conf()
	ca, _, _ := c.SearchConstantsAnime()

	if ca.Kind != nil {
		t.Logf("Found: %s", ca.Kind)
	} else {
		t.Error("Constants not found")
	}
	if ca.Status != nil {
		t.Logf("Found: %s", ca.Status)
	} else {
		t.Error("Constants not found")
	}
}

func TestConstantsManga(t *testing.T) {
	c := conf()
	cm, _, _ := c.SearchConstantsManga()

	if cm.Kind != nil {
		t.Logf("Found: %s", cm.Kind)
	} else {
		t.Error("Constants not found")
	}
	if cm.Status != nil {
		t.Logf("Found: %s", cm.Status)
	} else {
		t.Error("Constants not found")
	}
}

func TestAnimeGraphql(t *testing.T) {
	var s utils.StatusBar
	s.Settings(5, "#", 1)
	s.Run()

	c := conf()
	sch, _ := AnimeSchema(
		ValuesSchema("id", "malId", "name", "rating", "kind", "episodes"),
		"initial d first stage", 1, 1, "", "", "", "", "", "", "", false, nil,
	)
	a, _, _ := c.SearchGraphql(sch)

	for _, v := range a.Data.Animes {
		if v.Id == "185" && v.MalId == "185" && v.Rating == constants.ANIME_RATING_PG_13 &&
			v.Kind == constants.ANIME_KIND_TV && v.Episodes == 26 {
			t.Logf("%s - found", v.Name)
		} else {
			t.Error("AnimeGraphql not found")
		}
	}
}

func TestMangaGraphQL(t *testing.T) {
	c := conf()
	s, _ := MangaSchema(
		ValuesSchema("id", "malId", "name", "kind", "status", "volumes"),
		"initial d", 1, 1, "", "", "", "", "", false, nil,
	)
	m, _, _ := c.SearchGraphql(s)

	for _, v := range m.Data.Mangas {
		if v.Id == "375" && v.MalId == "375" && v.Kind == constants.MANGA_KIND_MANGA &&
			v.Status == constants.MANGA_STATUS_RELEASED && v.Volumes == 48 {
			t.Logf("%s - found", v.Name)
		} else {
			t.Error("MangaGraphql not found")
		}
	}
}

func TestCharacterGraphQL(t *testing.T) {
	c := conf()
	s, _ := CharacterSchema(
		ValuesSchema("id", "malId", "name", "isManga"),
		"Natsuno Yuuki", 1, 1,
	)
	ch, _, _ := c.SearchGraphql(s)

	for _, v := range ch.Data.Characters {
		if v.Id == "7582" && v.MalId == "7582" && v.IsManga {
			t.Logf("%s - found", v.Name)
		} else {
			t.Error("CharacterGraphql not found")
		}
	}
}

func TestPeopleGraphQL(t *testing.T) {
	c := conf()
	s, _ := PeopleSchema(
		ValuesSchema("id", "name", "birthOn{year}"),
		"satsuki", 1, 1, true, false, false,
	)
	p, _, _ := c.SearchGraphql(s)

	for _, v := range p.Data.People {
		if v.Id == "3" && v.BirthOn.Year == 1970 {
			t.Logf("%s - found", v.Name)
		} else {
			t.Error("PeopleGraphql not found")
		}
	}
}

func TestAnimesUsingGenre(t *testing.T) {
	c := conf()
	o := &Options{Page: 1, Limit: 1, Genre_v2: []int{3}}
	s, _, _ := c.SearchAnimes("Initial D", o)

	for _, v := range s {
		if v.Id == 12725 && v.Status == "released" {
			t.Logf("Anime: %s, Id: %d - found", v.Name, v.Id)
		} else {
			t.Errorf("Anime: %s, Id: %d - not found", v.Name, v.Id)
		}
	}
}

func TestMangasUsingGenre(t *testing.T) {
	c := conf()
	o := &Options{Page: 1, Limit: 1, Genre_v2: []int{84}}
	r, _, _ := c.SearchMangas("Initial D", o)

	for _, v := range r {
		if v.Volumes == 48 && v.Chapters == 724 {
			t.Logf("Manga: %s, Id: %d - found", v.Name, v.Id)
		} else {
			t.Errorf("Manga: %s, Id: %d - not found", v.Name, v.Id)
		}
	}
}

func TestLanguageCheck(t *testing.T) {
	if utils.LanguageCheck("Ая Хирано") == "Ая Хирано" {
		t.Log("Cyrillic passed")
	} else {
		t.Error("Cyrillic failed")
	}

	if utils.LanguageCheck("Aya Hirano") == "Aya+Hirano" {
		t.Log("Latin passed")
	} else {
		t.Error("Latin failed")
	}
}

func TestCyrillicPeople(t *testing.T) {
	c := conf()

	fastc, status, _ := c.FastIdPeople("Томокадзу Сэки")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		pc, _, _ := fastc.SearchPeople()

		if pc.Id == 1 && pc.Job_title == "Сэйю" {
			t.Logf("%s - found (Cyrillic alphabet)", pc.Name)
		} else {
			t.Skip()
		}
	}

	fastl, status, _ := c.FastIdPeople("Aya Hirano")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		pl, _, _ := fastl.SearchPeople()

		if pl.Id == 4 && pl.Job_title == "Сэйю" {
			t.Logf("%s - found (Latin alphabet)", pl.Name)
		} else {
			t.Error("People not found (Latin alphabet)")
		}
	}
}

func TestCyrillicCharacter(t *testing.T) {
	c := conf()

	fastc, status, _ := c.FastIdCharacter("Такуми Усуи")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		pc, _, _ := fastc.SearchCharacter()

		if pc.Id == 14523 && pc.Altname == "Perverted Alien" {
			t.Logf("%s - found (Cyrillic alphabet)", pc.Name)
		} else {
			t.Skip()
		}
	}

	fastl, status, _ := c.FastIdCharacter("Takumi Usui")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		pl, _, _ := fastl.SearchCharacter()

		if pl.Id == 14523 && pl.Altname == "Perverted Alien" {
			t.Logf("%s - found (Latin alphabet)", pl.Name)
		} else {
			t.Error("Character not found (Latin alphabet)")
		}
	}
}

func TestCyrillicClub(t *testing.T) {
	var s utils.StatusBar
	s.Settings(5, "#", 1)
	s.Run()

	c := conf()

	fastc, status, _ := c.FastIdClub("Ачивки (достижения)")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		if fastc.Id == 315 {
			t.Logf("Fast club found (Cyrillic alphabet)")
		} else {
			t.Skip()
		}
	}

	fastl, status, _ := c.FastIdClub("Genshin Impact")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		if fastl.Id == 3057 {
			t.Logf("Fast club found (Latin alphabet)")
		} else {
			t.Error("Fast club not found (Latin alphabet)")
		}
	}
}

func TestCyrillicRanobe(t *testing.T) {
	c := conf()

	fastc, status, _ := c.FastIdRanobe("Ангел кровопролития")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		if fastc.Id == 115586 {
			t.Logf("Fast ranobe found (Cyrillic alphabet)")
		} else {
			t.Skip()
		}
	}

	fastl, status, _ := c.FastIdRanobe("Satsuriku no Tenshi")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		if fastl.Id == 115586 {
			t.Logf("Fast ranobe found (Latin alphabet)")
		} else {
			t.Error("Fast ranobe not found (Latin alphabet)")
		}
	}
}

func TestCyrillicManga(t *testing.T) {
	c := conf()

	fastc, status, _ := c.FastIdManga("Тетрадь смерти")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		if fastc.Id == 21 {
			t.Logf("Fast manga found (Cyrillic alphabet)")
		} else {
			t.Skip()
		}
	}

	fastl, status, _ := c.FastIdManga("Death Note")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		if fastl.Id == 21 {
			t.Logf("Fast manga found (Latin alphabet)")
		} else {
			t.Error("Fast manga not found (Latin alphabet)")
		}
	}
}

func TestCyrillicAnime(t *testing.T) {
	var s utils.StatusBar
	s.Settings(5, "#", 1)
	s.Run()

	c := conf()

	fastc, status, _ := c.FastIdAnime("Тетрадь смерти")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		if fastc.Id == 1535 {
			t.Logf("Fast anime found (Cyrillic alphabet)")
		} else {
			t.Skip()
		}
	}

	fastl, status, _ := c.FastIdAnime("Death Note")
	if status == -1 {
		t.Log("timeout from shikimori")
	} else {
		if fastl.Id == 1535 {
			t.Logf("Fast anime found (Latin alphabet)")
		} else {
			t.Error("Fast anime not found (Latin alphabet)")
		}
	}
}
