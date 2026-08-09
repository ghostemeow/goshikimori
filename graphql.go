package goshikimori

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/ghostemeow/goshikimori/genres"
	"github.com/ghostemeow/goshikimori/internal/concatenation"
)

// Available anime options:
//   - id malId name russian licenseNameRu english japanese synonyms kind rating score status episodes episodesAired duration airedOn{year month day date} releasedOn{year month day date} url season
//   - poster{id originalUrl mainUrl}
//   - fansubbers fandubbers licensors createdAt updatedAt nextEpisodeAt isCensored
//   - genres{id name russian kind}
//   - studios{id name imageUrl}
//   - personRoles{id rolesRu rolesEn person{id name poster{id}}}
//   - characterRoles{id rolesRu rolesEn character{id name poster{id}}}
//   - related{id anime{id name} manga{id name} relationKind relationText}
//   - videos{id url name kind}
//   - screenshots{id originalUrl x166Url x332Url}
//   - scoresStats{score count}
//   - statusesStats{status count}
//   - description descriptionHtml descriptionSource
//   - chronology{...} // All the parameters described above can be added to the timeline.
//
// Available manga options:
//   - id malId name russian licenseNameRu english japanese synonyms kind score status volumes chapters airedOn{year month day date} releasedOn{year month day date} url
//   - poster{id originalUrl mainUrl}
//   - licensors createdAt updatedAt isCensored
//   - genres{id name russian kind}
//   - publishers{id name}
//   - personRoles{id rolesRu rolesEn person{id name poster{id}}}
//   - characterRoles{id rolesRu rolesEn character{id name poster{id}}}
//   - related{id anime{id name} manga{id name} relationKind relationText}
//   - scoresStats{score count}
//   - statusesStats{status count}
//   - description descriptionHtml descriptionSource
//   - chronology{...} // All the parameters described above can be added to the timeline.
//
// Available character options:
//   - id malId name russian japanese synonyms url createdAt updatedAt isAnime isManga isRanobe
//   - poster{id originalUrl mainUrl}
//   - description descriptionHtml descriptionSource
//
// Available people options:
//   - id malId name russian japanese synonyms url isSeyu isMangaka isProducer website createdAt updatedAt
//   - birthOn{year month day date} deceasedOn{year month day date}
//   - poster{id originalUrl mainUrl}
//
// Available userrate options:
//   - id text createdAt updatedAt rewatches score status episodes chapters volumes
//   - anime{...} // All parameters from the "Available anime options:"
//   - manga{...} // All parameters from the "Available manga options:"
func ValuesSchema(input ...string) string {
	var res bytes.Buffer

	// We always return 1, even if the slice is empty.
	// But we are not allowed to return an empty value,
	// so an additional check is present and returns id.
	if len(input) == 1 && len(strings.TrimSpace(input[0])) == 0 {
		return "id"
	}

	// TODO (ghostemeow): add keyword checks and the keyword is "full" to add everything.
	for i := 0; i < len(input); i++ {
		res.WriteString(input[i])
		res.WriteString(" ")
	}

	return res.String()
}

// AnimeOptions are the options for AnimeSchema.
//
// The zero value of a field means the option is skipped (except Censored,
// which is always written).
type AnimeOptions struct {
	Page     int
	Limit    int
	Score    int
	Order    string
	Kind     string
	Status   string
	Season   string
	Duration string
	Rating   string
	Mylist   string
	Censored bool
	GenreV2  []int
}

// MangaOptions are the options for MangaSchema.
//
// The zero value of a field means the option is skipped (except Censored,
// which is always written).
type MangaOptions struct {
	Page     int
	Limit    int
	Score    int
	Order    string
	Kind     string
	Status   string
	Season   string
	Mylist   string
	Censored bool
	GenreV2  []int
}

// CharacterOptions are the options for CharacterSchema.
//
// The zero value of a field means the option is skipped.
type CharacterOptions struct {
	Page  int
	Limit int
}

// PeopleOptions are the options for PeopleSchema.
type PeopleOptions struct {
	Page       int
	Limit      int
	IsSeyu     bool
	IsMangaka  bool
	IsProducer bool
}

// UserRatesOptions are the options for UserRatesSchema.
//
// The zero value of a field means the option is skipped.
type UserRatesOptions struct {
	Page       int
	Limit      int
	Status     string
	TargetType string
	Order      string
}

// Values: parameters we want to receive from the server.
//
// Name: anime name.
//
// If you use the 'order' parameter, you don't need to enter the name of the anime.
//
// Exclamation mark(!) indicates ignore.
//
// Options: an AnimeOptions struct. The zero value of a field means the option is skipped
// (except Censored, which is always written).
//
// 'Options' settings:
//
//   - Page: >= 1;
//
//   - Limit: <= 50;
//
//   - Order:
//
//     > ANIME_ORDER_ID, ANIME_ORDER_RANKED, ANIME_ORDER_KIND,
//     ANIME_ORDER_POPULARITY, ANIME_ORDER_NAME, ANIME_ORDER_AIRED_ON,
//     ANIME_ORDER_EPISODES, ANIME_ORDER_STATUS;
//
//   - Kind:
//
//     > ANIME_KIND_TV, ANIME_KIND_MOVIE, ANIME_KIND_OVA, ANIME_KIND_ONA,
//     ANIME_KIND_SPECIAL, ANIME_KIND_MUSIC, ANIME_KIND_TV_13, ANIME_KIND_TV_24,
//     ANIME_KIND_TV_48, ANIME_KIND_TV_NOT_EQUAL, ANIME_KIND_MOVIE_NOT_EQUAL,
//     ANIME_KIND_OVA_NOT_EQUAL, ANIME_KIND_ONA_NOT_EQUAL, ANIME_KIND_SPECIAL_NOT_EQUAL,
//     ANIME_KIND_MUSIC_NOT_EQUAL, ANIME_KIND_TV_13_NOT_EQUAL,
//     ANIME_KIND_TV_24_NOT_EQUAL, ANIME_KIND_TV_48_NOT_EQUAL;
//
//   - Status:
//
//     > ANIME_STATUS_ANONS, ANIME_STATUS_ONGOING, ANIME_STATUS_RELEASED,
//     ANIME_STATUS_ANONS_NOT_EQUAL, ANIME_STATUS_ONGOING_NOT_EQUAL,
//     ANIME_STATUS_RELEASED_NOT_EQUAL;
//
//   - Season:
//
//     > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//     SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//     SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//     SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//     SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//   - Duration:
//
//     > ANIME_DURATION_S, ANIME_DURATION_D, ANIME_DURATION_F,
//     ANIME_DURATION_S_NOT_EQUAL, ANIME_DURATION_D_NOT_EQUAL,
//     ANIME_DURATION_F_NOT_EQUAL;
//
//   - Rating:
//
//     > ANIME_RATING_NONE, ANIME_RATING_G, ANIME_RATING_PG,
//     ANIME_RATING_PG_13, ANIME_RATING_R, ANIME_RATING_R_PLUS, ANIME_RATING_RX,
//     ANIME_RATING_G_NOT_EQUAL, ANIME_RATING_PG_NOT_EQUAL,
//     ANIME_RATING_PG_13_NOT_EQUAL, ANIME_RATING_R_NOT_EQUAL,
//     ANIME_RATING_R_PLUS_NOT_EQUAL, ANIME_RATING_RX_NOT_EQUAL;
//
//   - Mylist:
//
//     > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//     MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
//   - Score: 1-9 maximum;
//
//   - Censored: true, false;
//
//   - Genre_v2: id search. Below is a list of all available genres by id:
//
//     > 1 (Action); 2 (Adventure); 3 (Cars); 4 (Comedy); 5 (Dementia); 6 (Demons); 7 (Mystery);
//     8 (Drama); 9 (Ecchi); 10 (Fantasy); 11 (Game); 12 (Hentai); 13 (Historical); 14 (Horror);
//     15 (Kids); 16 (Magic); 17 (Martial Arts); 18 (Mecha); 19 (Music); 20 (Parody); 21 (Samurai);
//     22 (Romance); 23 (School); 24 (Sci-Fi); 25 (Shoujo); 26 (Shoujo Ai); 27 (Shounen); 28 (Shounen Ai);
//     29 (Space); 30 (Sports); 31 (Super Power); 32 (Vampire); 33 (Yaoi); 34 (Yuri); 35 (Harem);
//     36 (Slice of Life); 37 (Supernatural); 38 (Military); 39 (Police); 40 (Psychological);
//     41 (Thriller); 42 (Seinen); 43 (Josei); 539 (Erotica); 541 (Work Life); 543 (Gourmet);
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL.md
func AnimeSchema(values string, name string, options AnimeOptions) string {
	var parameterOptions bytes.Buffer

	if options.Page >= 1 {
		parameterOptions.WriteString(", page: ")
		parameterOptions.WriteString(strconv.Itoa(options.Page))
	}
	if options.Limit >= 1 && options.Limit <= 50 {
		parameterOptions.WriteString(", limit: ")
		parameterOptions.WriteString(strconv.Itoa(options.Limit))
	}
	if options.Score >= 1 && options.Score <= 9 {
		parameterOptions.WriteString(", score: ")
		parameterOptions.WriteString(strconv.Itoa(options.Score))
	}
	if options.Order != "" {
		parameterOptions.WriteString(", order: ")
		parameterOptions.WriteString(options.Order)
	}
	if options.Kind != "" {
		parameterOptions.WriteString(", kind: \"")
		parameterOptions.WriteString(options.Kind)
		parameterOptions.WriteString("\"")
	}
	if options.Status != "" {
		parameterOptions.WriteString(", status: \"")
		parameterOptions.WriteString(options.Status)
		parameterOptions.WriteString("\"")
	}
	if options.Season != "" {
		parameterOptions.WriteString(", season: \"")
		parameterOptions.WriteString(options.Season)
		parameterOptions.WriteString("\"")
	}
	if options.Duration != "" {
		parameterOptions.WriteString(", duration: \"")
		parameterOptions.WriteString(options.Duration)
		parameterOptions.WriteString("\"")
	}
	if options.Rating != "" {
		parameterOptions.WriteString(", rating: \"")
		parameterOptions.WriteString(options.Rating)
		parameterOptions.WriteString("\"")
	}
	if options.Mylist != "" {
		parameterOptions.WriteString(", mylist: \"")
		parameterOptions.WriteString(options.Mylist)
		parameterOptions.WriteString("\"")
	}
	parameterOptions.WriteString(", censored: ")
	parameterOptions.WriteString(strconv.FormatBool(options.Censored))
	if genre := genres.MapGenresAnime(options.GenreV2); genre != "" {
		parameterOptions.WriteString(", genre: \"")
		parameterOptions.WriteString(genre)
		parameterOptions.WriteString("\"")
	}

	// 22({animes(search: " "){}}) + ?(name) + ?(parameterOptions) + ?(values)
	return concatenation.Url(22+len(name)+len(parameterOptions.String())+len(values), []string{
		"{animes(search: \"", name, "\"",
		parameterOptions.String(), ")",
		"{", values, "}}",
	})
}

// Values: parameters we want to receive from the server.
//
// Name: manga name.
//
// If you use the 'order' parameter, you don't need to enter the name of the anime.
//
// Exclamation mark(!) indicates ignore.
//
// Options: an MangaOptions struct. The zero value of a field means the option is skipped
// (except Censored, which is always written).
//
// 'Options' settings:
//
//   - Page: >= 1;
//
//   - Limit: <= 50;
//
//   - Order:
//
//     > MANGA_ORDER_ID, MANGA_ORDER_RANKED, MANGA_ORDER_KIND, MANGA_ORDER_POPULARITY,
//     MANGA_ORDER_NAME, MANGA_ORDER_AIRED_ON, MANGA_ORDER_VOLUMES,
//     MANGA_ORDER_CHAPTERS, MANGA_ORDER_STATUS;
//
//   - Kind:
//
//     > MANGA_KIND_MANGA, MANGA_KIND_MANHWA, MANGA_KIND_MANHUA, MANGA_KIND_LIGHT_NOVEL,
//     MANGA_KIND_NOVEL, MANGA_KIND_ONE_SHOT, MANGA_KIND_DOUJIN, MANGA_KIND_MANGA_NOT_EQUAL,
//     MANGA_KIND_MANHWA_NOT_EQUAL, MANGA_KIND_MANHUA_NOT_EQUAL, MANGA_KIND_LIGHT_NOVEL_NOT_EQUAL,
//     MANGA_KIND_NOVEL_NOT_EQUAL, MANGA_KIND_ONE_SHOT_NOT_EQUAL, MANGA_KIND_DOUJIN_NOT_EQUAL;
//
//   - Status:
//
//     > MANGA_STATUS_ANONS, MANGA_STATUS_ONGOING, MANGA_STATUS_RELEASED, MANGA_STATUS_PAUSED,
//     MANGA_STATUS_DISCONTINUED, MANGA_STATUS_ANONS_NOT_EQUAL, MANGA_STATUS_ONGOING_NOT_EQUAL,
//     MANGA_STATUS_RELEASED_NOT_EQUAL, MANGA_STATUS_PAUSED_NOT_EQUAL, MANGA_STATUS_DISCONTINUED_NOT_EQUAL;
//
//   - Season:
//
//     > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//     SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//     SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//     SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//     SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//   - Mylist:
//
//     > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//     MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
//   - Score: 1-9 maximum;
//
//   - Censored: true, false;
//
//   - Genre_v2: id search. Below is a list of all available genres by id:
//
//     > 46 (Mystery); 47 (Shounen); 48 (Supernatural);
//     49 (Comedy); 50 (Drama); 51 (Ecchi); 52 (Seinen); 53 (Sci-Fi); 54 (Slice of Life); 55 (Shounen Ai);
//     56 (Action); 57 (Fantasy); 58 (Magic); 59 (Hentai); 60 (School); 61 (Doujinshi); 62 (Romance);
//     63 (Shoujo); 64 (Vampire); 65 (Yaoi); 66 (Martial Arts); 67 (Psychological); 68 (Adventure);
//     69 (Historical); 70 (Military); 71 (Harem); 72 (Demons); 73 (Shoujo Ai); 74 (Gender Bender);
//     75 (Yuri); 76 (Sports); 77 (Kids); 78 (Music); 79 (Game); 80 (Horror); 81 (Thriller);
//     82 (Super Power); 83 (Mecha); 84 (Cars); 85 (Space); 86 (Parody); 87 (Josei); 88 (Samurai);
//     89 (Police); 90 (Dementia); 540 (Erotica); 542 (Work Life); 544 (Gourmet);
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL.md
func MangaSchema(values string, name string, options MangaOptions) string {
	var parameterOptions bytes.Buffer

	if options.Page >= 1 {
		parameterOptions.WriteString(", page: ")
		parameterOptions.WriteString(strconv.Itoa(options.Page))
	}
	if options.Limit >= 1 && options.Limit <= 50 {
		parameterOptions.WriteString(", limit: ")
		parameterOptions.WriteString(strconv.Itoa(options.Limit))
	}
	if options.Score >= 1 && options.Score <= 9 {
		parameterOptions.WriteString(", score: ")
		parameterOptions.WriteString(strconv.Itoa(options.Score))
	}
	if options.Order != "" {
		parameterOptions.WriteString(", order: ")
		parameterOptions.WriteString(options.Order)
	}
	if options.Kind != "" {
		parameterOptions.WriteString(", kind: \"")
		parameterOptions.WriteString(options.Kind)
		parameterOptions.WriteString("\"")
	}
	if options.Status != "" {
		parameterOptions.WriteString(", status: \"")
		parameterOptions.WriteString(options.Status)
		parameterOptions.WriteString("\"")
	}
	if options.Season != "" {
		parameterOptions.WriteString(", season: \"")
		parameterOptions.WriteString(options.Season)
		parameterOptions.WriteString("\"")
	}
	if options.Mylist != "" {
		parameterOptions.WriteString(", mylist: \"")
		parameterOptions.WriteString(options.Mylist)
		parameterOptions.WriteString("\"")
	}
	parameterOptions.WriteString(", censored: ")
	parameterOptions.WriteString(strconv.FormatBool(options.Censored))
	if genre := genres.MapGenresManga(options.GenreV2); genre != "" {
		parameterOptions.WriteString(", genre: \"")
		parameterOptions.WriteString(genre)
		parameterOptions.WriteString("\"")
	}

	// 22({mangas(search: " "){}}) + ?(name) + ?(parameterOptions) + ?(values)
	return concatenation.Url(22+len(name)+len(parameterOptions.String())+len(values), []string{
		"{mangas(search: \"", name, "\"",
		parameterOptions.String(), ")",
		"{", values, "}}",
	})
}

// Values: parameters we want to receive from the server.
//
// Name: character name.
//
// Options: a CharacterOptions struct. The zero value of a field means the option is skipped.
//
// 'Options' settings:
//   - Page: >= 1;
//   - Limit: <= 50;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL.md
func CharacterSchema(values string, name string, options CharacterOptions) string {
	var parameterOptions bytes.Buffer

	if options.Page >= 1 {
		parameterOptions.WriteString(", page: ")
		parameterOptions.WriteString(strconv.Itoa(options.Page))
	}
	if options.Limit >= 1 && options.Limit <= 50 {
		parameterOptions.WriteString(", limit: ")
		parameterOptions.WriteString(strconv.Itoa(options.Limit))
	}

	// 26({characters(search: " "){}}) + ?(name) + ?(parameterOptions) + ?(values)
	return concatenation.Url(26+len(name)+len(parameterOptions.String())+len(values), []string{
		"{characters(search: \"", name, "\"",
		parameterOptions.String(), ")",
		"{", values, "}}",
	})
}

// Values: parameters we want to receive from the server.
//
// Name: people name.
//
// Options: a PeopleOptions struct. The zero value of a field means the option is skipped.
//
// 'Options' settings:
//   - Page: >= 1;
//   - Limit: <= 50;
//   - isSeyu: true, false;
//   - isMangaka: true, false;
//   - isProducer: true, false;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL.md
func PeopleSchema(values string, name string, options PeopleOptions) string {
	var parameterOptions bytes.Buffer

	if options.Page >= 1 {
		parameterOptions.WriteString(", page: ")
		parameterOptions.WriteString(strconv.Itoa(options.Page))
	}
	if options.Limit >= 1 && options.Limit <= 50 {
		parameterOptions.WriteString(", limit: ")
		parameterOptions.WriteString(strconv.Itoa(options.Limit))
	}
	parameterOptions.WriteString(", isSeyu: ")
	parameterOptions.WriteString(strconv.FormatBool(options.IsSeyu))
	parameterOptions.WriteString(", isMangaka: ")
	parameterOptions.WriteString(strconv.FormatBool(options.IsMangaka))
	parameterOptions.WriteString(", isProducer: ")
	parameterOptions.WriteString(strconv.FormatBool(options.IsProducer))

	// 22({people(search: " "){}}) + ?(name) + ?(parameterOptions) + ?(values)
	return concatenation.Url(22+len(name)+len(parameterOptions.String())+len(values), []string{
		"{people(search: \"", name, "\"",
		parameterOptions.String(), ")",
		"{", values, "}}",
	})
}

// Auxiliary function for UserRatesSchema().
//
//   - Filed:
//
//     > GRAPHQL_ORDER_FIELD_ID, GRAPHQL_ORDER_FIELD_UPDATED_AT;
//
//   - Order:
//
//     > GRAPHQL_ORDER_ORDER_ASC, GRAPHQL_ORDER_ORDER_DESC;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL.md
func UserRatesOrder(field, order string) string {
	var res bytes.Buffer

	if field != "" {
		res.WriteString(", order: { field: ")
		res.WriteString(field)
		if order != "" {
			res.WriteString(", order: ")
			res.WriteString(order)
		}
		res.WriteString(" }")
	} else {
		if order != "" {
			res.WriteString(", order: { order: ")
			res.WriteString(order)
			res.WriteString(" }")
		}
	}

	return res.String()
}

// Values: parameters we want to receive from the server.
//
// UserId: user id.
//
// Options: an UserRatesOptions struct. The zero value of a field means the option is skipped.
//
// 'Options' settings:
//
//   - Page: >= 1;
//
//   - Limit: <= 50;
//
//   - Status:
//
//     > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//     MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
//   - TargetType:
//
//     > TARGET_TYPE_ANIME, TARGET_TYPE_MANGA;
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL.md
func UserRatesSchema(values string, userId int, options UserRatesOptions) string {
	var parameterOptions bytes.Buffer

	id := strconv.Itoa(userId)

	if options.Page >= 1 {
		parameterOptions.WriteString(", page: ")
		parameterOptions.WriteString(strconv.Itoa(options.Page))
	}
	if options.Limit >= 1 && options.Limit <= 50 {
		parameterOptions.WriteString(", limit: ")
		parameterOptions.WriteString(strconv.Itoa(options.Limit))
	}
	if options.Status != "" {
		parameterOptions.WriteString(", status: ")
		parameterOptions.WriteString(options.Status)
	}
	if options.TargetType != "" {
		parameterOptions.WriteString(", targetType: ")
		parameterOptions.WriteString(options.TargetType)
	}
	if options.Order != "" {
		parameterOptions.WriteString(options.Order)
	}

	// 24({userRates(userId: 0){}}) + ?(id) + ?(parameterOptions) + ?(values)
	return concatenation.Url(24+len(id)+len(parameterOptions.String())+len(values), []string{
		"{userRates(userId: ", id,
		parameterOptions.String(), ")",
		"{", values, "}}",
	})
}

// TODO (ghostemeow): create query with variables.
