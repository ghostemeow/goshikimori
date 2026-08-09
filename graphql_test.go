package goshikimori

import "testing"

// Constants are not available in tests, so they are implemented manually.

func TestAnimeSchema(t *testing.T) {
	pass_normal := `{animes(search: "initial d", page: 1, limit: 1, score: 8, order: id, kind: "tv", status: "!anons", season: "199x", duration: "F", rating: "!rx", mylist: "completed", censored: false, genre: "3-Cars,4-Comedy"){id name russian english japanese score airedOn{year month day date} }}`
	normal := AnimeSchema(
		ValuesSchema("id", "name", "russian", "english", "japanese", "score", "airedOn{year month day date}"),
		"initial d",
		AnimeOptions{Page: 1, Limit: 1, Score: 8, Order: "id", Kind: "tv", Status: "!anons", Season: "199x", Duration: "F", Rating: "!rx", Mylist: "completed", Censored: false, GenreV2: []int{3, 4, 4, 3}},
	)
	if normal == pass_normal {
		t.Log("Normal AnimeSchema passed")
	} else {
		t.Error("Normal AnimeSchema failed")
	}

	pass_empty := `{animes(search: "initial d", page: 1, limit: 1, score: 1, censored: false){id}}`
	empty := AnimeSchema(
		ValuesSchema(""),
		"initial d",
		AnimeOptions{Page: 1, Limit: 1, Score: 1},
	)
	if empty == pass_empty {
		t.Log("Empty AnimeSchema passed")
	} else {
		t.Error("Empty AnimeSchema failed")
	}
}

func TestMangaSchema(t *testing.T) {
	pass_normal := `{mangas(search: "angel", page: 1, limit: 3, score: 8, order: ranked, kind: "manga", status: "released", mylist: "planned", censored: false, genre: "50-Drama,64-Vampire"){id name russian volumes chapters releasedOn{year month day date} url }}`
	normal := MangaSchema(
		ValuesSchema("id", "name", "russian", "volumes", "chapters", "releasedOn{year month day date}", "url"),
		"angel",
		MangaOptions{Page: 1, Limit: 3, Score: 8, Order: "ranked", Kind: "manga", Status: "released", Mylist: "planned", Censored: false, GenreV2: []int{50, 64, 64, 50}},
	)
	if normal == pass_normal {
		t.Log("Normal MangaSchema passed")
	} else {
		t.Error("Normal MangaSchema failed")
	}

	pass_empty := `{mangas(search: "initial d", page: 1, limit: 1, score: 1, censored: false){id}}`
	empty := MangaSchema(
		ValuesSchema(""),
		"initial d",
		MangaOptions{Page: 1, Limit: 1, Score: 1},
	)
	if empty == pass_empty {
		t.Log("Empty MangaSchema passed")
	} else {
		t.Error("Empty MangaSchema failed")
	}
}

func TestCharacterSchema(t *testing.T) {
	pass := `{characters(search: "onizuka", page: 1, limit: 1){id name russian poster{originalUrl} description }}`
	normal := CharacterSchema(
		ValuesSchema("id", "name", "russian", "poster{originalUrl}", "description"),
		"onizuka",
		CharacterOptions{Page: 1, Limit: 1},
	)
	if normal == pass {
		t.Log("Normal CharacterSchema passed")
	} else {
		t.Error("Normal CharacterSchema failed")
	}
}

func TestPeopleSchema(t *testing.T) {
	pass := `{people(search: "satsuki", page: 1, limit: 1, isSeyu: true, isMangaka: false, isProducer: false){id name russian url website birthOn{year month day date} }}`
	normal := PeopleSchema(
		ValuesSchema("id", "name", "russian", "url", "website", "birthOn{year month day date}"),
		"satsuki",
		PeopleOptions{Page: 1, Limit: 1, IsSeyu: true},
	)
	if normal == pass {
		t.Log("Normal PeopleSchema passed")
	} else {
		t.Error("Normal PeopleSchema failed")
	}
}

func TestUserRatesSchema(t *testing.T) {
	pass := `{userRates(userId: 181833, page: 1, limit: 10, status: completed, targetType: Anime, order: { field: id, order: desc }){id text score createdAt anime {name} }}`
	normal := UserRatesSchema(
		ValuesSchema("id", "text", "score", "createdAt", "anime {name}"),
		181833,
		UserRatesOptions{Page: 1, Limit: 10, Status: "completed", TargetType: "Anime", Order: UserRatesOrder("id", "desc")},
	)
	if normal == pass {
		t.Log("Normal UserRatesSchema passed")
	} else {
		t.Error("Normal UserRatesSchema failed")
	}

	pass_empty := `{userRates(userId: 181833, page: 1, limit: 1){id}}`
	empty := UserRatesSchema(
		ValuesSchema(""),
		181833,
		UserRatesOptions{Page: 1, Limit: 1},
	)
	if empty == pass_empty {
		t.Log("Empty UserRatesSchema passed")
	} else {
		t.Error("Empty UserRatesSchema failed")
	}
}
