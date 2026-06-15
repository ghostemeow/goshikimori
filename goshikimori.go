// Copyright (C) 2026 ghostemeow <andreyisback@yandex.ru>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Comments are made in the style of "godoc" syntax support.
//
// More information can be found in the [examples] folder.
//
// [examples]: https://github.com/ghostemeow/goshikimori/blob/master/examples/
package goshikimori

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	"github.com/ghostemeow/goshikimori/internal/concatination"
	"github.com/ghostemeow/goshikimori/constants"
	"github.com/ghostemeow/goshikimori/genres"
	"github.com/ghostemeow/goshikimori/internal/models"
	"github.com/ghostemeow/goshikimori/internal/request"
)

// Only the application needs to be specified in SetConfiguration().
//
// Name: user name.
//
// Search by user is case sensitive.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (c *Configuration) SearchUser(name string) (models.Users, int, error) {
	var u models.Users

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 6(users/) + ?(name)
		concatination.Url(32+len(name), []string{constants.SITE, "users/", url.QueryEscape(name)}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return u, status, err
	}

	if err := json.Unmarshal(data, &u); err != nil {
		return u, status, err
	}

	return u, status, nil
}

// FIXME (ghostemeow): Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// Name: user name.
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 100 maximum;
//
// Don't use Stats.Statuses.Anime and Stats.Statuses.Manga: empty slice.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/users
func (c *Configuration) SearchUsers(name string, r Result) ([]models.Users, int, error) {
	var u []models.Users

	opt := r.OptionsOnlyPageLimitV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 13(users?search=) + ?(name) + 1(&) + ?(Result)
		concatination.Url(40+len(name)+len(opt), []string{
			constants.SITE, "users?search=", url.QueryEscape(name), "&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &u); err != nil {
		return u, status, err
	}

	return u, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserFriends(r Result) ([]models.UserFriends, int, error) {
	var uf []models.UserFriends

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(users/) + ?(id) + 9(/friends?) + ?(Result)
		concatination.Url(41+len(str_id)+len(opt), []string{
			constants.SITE, "users/", str_id, "/friends?" + opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &uf); err != nil {
		return nil, status, err
	}

	return uf, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserClubs() ([]models.Clubs, int, error) {
	var uc []models.Clubs

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(users/) + ?(id) + 6(/clubs)
		concatination.Url(38+len(str_id), []string{
			constants.SITE, "users/", str_id, "/clubs",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &uc); err != nil {
		return nil, status, err
	}

	return uc, status, nil
}

// FIXME (ghostemeow): Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//
//   - Page: 100000 maximum;
//
//   - Limit: 5000 maximum;
//
//   - Censored: true, false;
//
//   - Status:
//
//     > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//     MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;
//
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserAnimeRates(r Result) ([]models.UserAnimeRates, int, error) {
	var ar []models.UserAnimeRates

	opt := r.OptionsAnimeRatesV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(users/) + ?(id) + 13(/anime_rates?) + ?(Result)
		concatination.Url(45+len(str_id)+len(opt), []string{
			constants.SITE, "users/", str_id, "/anime_rates?" + opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &ar); err != nil {
		return nil, status, err
	}

	return ar, status, nil
}

// FIXME (ghostemeow): Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' Settings:
//   - Page: 100000 maximum;
//   - Limit: 5000 maximum;
//   - Censored: true, false;
//
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserMangaRates(r Result) ([]models.UserMangaRates, int, error) {
	var mr []models.UserMangaRates

	opt := r.OptionsMangaRatesV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(users/) + ?(id) + 13(/manga_rates?) + ?(Result)
		concatination.Url(45+len(str_id)+len(opt), []string{
			constants.SITE, "users/", str_id, "/manga_rates?" + opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &mr); err != nil {
		return nil, status, err
	}

	return mr, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserFavourites() (models.UserFavourites, int, error) {
	var uf models.UserFavourites

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(users/) + ?(id) + 11(/favourites)
		concatination.Url(43+len(str_id), []string{
			constants.SITE, "users/", str_id, "/favourites",
		}),constants.MAX_EXPECTATION,
	)
	if err != nil {
		return uf, status, err
	}

	if err := json.Unmarshal(data, &uf); err != nil {
		return uf, status, err
	}

	return uf, status, nil
}

// FIXME (ghostemeow): Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//
//   - Page: 100000 maximum;
//
//   - Limit: 100 maximum;
//
//   - Target_id: id anime/manga/ranobe;
//
//   - Target_type:
//
//     > TARGET_TYPE_ANIME, TARGET_TYPE_MANGA;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserHistory(r Result) ([]models.UserHistory, int, error) {
	var uh []models.UserHistory

	opt := r.OptionsUserHistoryV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(users/) + ?(id) + 9(/history?) + ?(Result)
		concatination.Url(41+len(str_id)+len(opt), []string{
			constants.SITE, "users/", str_id, "/history?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &uh); err != nil {
		return nil, status, err
	}

	return uh, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (f *FastId) SearchUserBans() ([]models.Bans, int, error) {
	var b []models.Bans

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(users/) + ?(id) + 5(/bans)
		concatination.Url(37+len(str_id), []string{
			constants.SITE, "users/", str_id, "/bans",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &b); err != nil {
		return nil, status, err
	}

	return b, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/whoami
func (c *Configuration) WhoAmi() (models.Who, int, error) {
	var w models.Who

	data, status, err := request.NewGetRequestWithCancelAndBearer(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 12(users/whoami)
		concatination.Url(38, []string{constants.SITE, "users/whoami"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return w, status, err
	}

	if err := json.Unmarshal(data, &w); err != nil {
		return w, status, err
	}

	return w, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchAnime() (models.Anime, int, error) {
	var a models.Anime

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(animes/) + ?(id)
		concatination.Url(33+len(str_id), []string{
			constants.SITE, "animes/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return a, status, err
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return a, status, err
	}

	return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: anime name.
//
// 'Options' settings:
//
//   - Page: 100000 maximum;
//
//   - Limit: 50 maximum;
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
// Set to true to discard hentai, yaoi and yuri.
//
//   - Type: "deprecated";
//   - Studio: not supported;
//   - Franchise: not supported;
//   - Ids: not supported;
//   - Exclude_ids: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchAnimes(name string, r Result) ([]models.Animes, int, error) {
	var a []models.Animes

	opt := r.OptionsAnimeV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 14(animes?search=) + ?(name) + 1(&) + ?(Result)
		concatination.Url(41+len(name)+len(opt), []string{
			constants.SITE, "animes?search=", url.QueryEscape(name), "&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return nil, status, err
	}

	return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchManga() (models.Manga, int, error) {
	var m models.Manga

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(mangas/) + ?(id)
		concatination.Url(33+len(str_id), []string{
			constants.SITE, "mangas/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return m, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return m, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: manga name.
//
// 'Options' settings:
//
//   - Page: 100000 maximum;
//
//   - Limit: 50 maximum;
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
// Set to true to discard hentai, yaoi and yuri.
//
//   - Type: "deprecated";
//   - Publisher: not supported;
//   - Franchise: not supported;
//   - Ids: not supported;
//   - Exclude_ids: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchMangas(name string, r Result) ([]models.Mangas, int, error) {
	var m []models.Mangas

	opt := r.OptionsMangaV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 14(mangas?search=) + ?(name) + 1(&) + ?(Result)
		concatination.Url(41+len(name)+len(opt), []string{
			constants.SITE, "mangas?search=", url.QueryEscape(name), "&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/anime_manga_ranobe
func (f *FastId) SearchRanobe() (models.Manga, int, error) {
	var m models.Manga

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(ranobe/) + ?(id)
		concatination.Url(33+len(str_id), []string{
			constants.SITE, "ranobe/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return m, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return m, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: ranobe name.
//
// 'Options' settings:
//
//   - Page: 100000 maximum;
//
//   - Limit: 50 maximum;
//
//   - Order:
//
//     > MANGA_ORDER_ID, MANGA_ORDER_RANKED, MANGA_ORDER_KIND, MANGA_ORDER_POPULARITY,
//     MANGA_ORDER_NAME, MANGA_ORDER_AIRED_ON, MANGA_ORDER_VOLUMES,
//     MANGA_ORDER_CHAPTERS, MANGA_ORDER_STATUS;
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
// Set to true to discard hentai, yaoi and yuri.
//
//   - Publisher: not supported;
//   - Franchise: not supported;
//   - Ids: not supported;
//   - Exclude_ids: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/anime_manga_ranobe
func (c *Configuration) SearchRanobes(name string, r Result) ([]models.Mangas, int, error) {
	var m []models.Mangas

	opt := r.OptionsRanobeV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 14(ranobe?search=) + ?(name) + 1(&) + ?(Result)
		concatination.Url(41+len(name)+len(opt), []string{
			constants.SITE, "ranobe?search=", url.QueryEscape(name), "&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: user name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdUser(name string) (*FastId, int, error) {
	var u models.Users

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 6(users/) + ?(name)
		concatination.Url(32+len(name), []string{
			constants.SITE, "users/", url.QueryEscape(name)}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &u); err != nil {
		return &FastId{Id: 0, Conf: *c, Err: err}, status, err
	}

	return &FastId{Id: u.Id, Conf: *c, Err: err}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: anime name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdAnime(name string) (*FastId, int, error) {
	var a []models.Animes

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 14(animes?search=) + ?(name)
		concatination.Url(40+len(name), []string{
			constants.SITE, "animes?search=", url.QueryEscape(name)}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return &FastId{Id: 0, Conf: *c, Err: err}, status, err
	}

	// if len == 0; we get panic: runtime error.
	// To avoid a crash, process the error here.
	//
	// There is no point in processing the error. there is no place to catch it.
	if len(a) == 0 {
		return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil
	}

	return &FastId{Id: a[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: manga name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdManga(name string) (*FastId, int, error) {
	var m []models.Mangas

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 14(mangas?search=) + ?(name)
		concatination.Url(40+len(name), []string{
			constants.SITE, "mangas?search=", url.QueryEscape(name)}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return &FastId{Id: 0, Conf: *c, Err: err}, status, err
	}

	// if len == 0; we get panic: runtime error.
	// To avoid a crash, process the error here.
	//
	// There is no point in processing the error. there is no place to catch it.
	if len(m) == 0 {
		return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil
	}

	return &FastId{Id: m[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: ranobe name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdRanobe(name string) (*FastId, int, error) {
	var m []models.Mangas

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 14(ranobe?search=) + ?(name)
		concatination.Url(40+len(name), []string{
			constants.SITE, "ranobe?search=", url.QueryEscape(name)}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return &FastId{Id: 0, Conf: *c, Err: err}, status, err
	}

	// if len == 0; we get panic: runtime error.
	// To avoid a crash, process the error here.
	//
	// There is no point in processing the error. there is no place to catch it.
	if len(m) == 0 {
		return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil
	}

	return &FastId{Id: m[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: club name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdClub(name string) (*FastId, int, error) {
	var cl []models.Clubs

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 13(clubs?search=) + ?(name)
		concatination.Url(39+len(name), []string{
			constants.SITE, "clubs?search=", url.QueryEscape(name)}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &cl); err != nil {
		return &FastId{Id: 0, Conf: *c, Err: err}, status, err
	}

	// if len == 0; we get panic: runtime error.
	// To avoid a crash, process the error here.
	//
	// There is no point in processing the error. there is no place to catch it.
	if len(cl) == 0 {
		return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil
	}

	return &FastId{Id: cl[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: character name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdCharacter(name string) (*FastId, int, error) {
	var ch []models.CharacterInfo

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 25(characters/search?search=) + ?(name)
		concatination.Url(51+len(name), []string{
			constants.SITE, "characters/search?search=", url.QueryEscape(name),
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &ch); err != nil {
		return &FastId{Id: 0, Conf: *c, Err: err}, status, err
	}

	// if len == 0; we get panic: runtime error.
	// To avoid a crash, process the error here.
	//
	// There is no point in processing the error. there is no place to catch it.
	if len(ch) == 0 {
		return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil
	}

	return &FastId{Id: ch[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: people name.
//
// Search by user is case sensitive.
func (c *Configuration) FastIdPeople(name string) (*FastId, int, error) {
	var ap []models.AllPeople
	// testing

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 21(people/search?search=) + ?(name)
		concatination.Url(47+len(name), []string{
			constants.SITE, "people/search?search=", url.QueryEscape(name),
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &ap); err != nil {
		return &FastId{Id: 0, Conf: *c, Err: err}, status, err
	}

	// if len == 0; we get panic: runtime error.
	// To avoid a crash, process the error here.
	//
	// There is no point in processing the error. there is no place to catch it.
	if len(ap) == 0 {
		return &FastId{Id: 0, Conf: *c, Err: nil}, status, nil
	}

	return &FastId{Id: ap[0].Id, Conf: *c, Err: nil}, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/anime_screenshots
func (f *FastId) SearchAnimeScreenshots() ([]models.AnimeScreenshots, int, error) {
	var s []models.AnimeScreenshots

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(animes/) + ?(id) + 12(/screenshots)
		concatination.Url(45+len(str_id), []string{
			constants.SITE, "animes/", str_id, "/screenshots",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return nil, status, err
	}

	return s, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchAnimeFranchise() (models.Franchise, int, error) {
	var ff models.Franchise

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(animes/) + ?(id) + 10(/franchise)
		concatination.Url(43+len(str_id), []string{
			constants.SITE, "animes/", str_id, "/franchise",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ff, status, err
	}

	if err := json.Unmarshal(data, &ff); err != nil {
		return ff, status, err
	}

	return ff, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchMangaFranchise() (models.Franchise, int, error) {
	var ff models.Franchise

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(mangas/) + ?(id) + 10(/franchise)
		concatination.Url(43+len(str_id), []string{
			constants.SITE, "mangas/", str_id, "/franchise",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ff, status, err
	}

	if err := json.Unmarshal(data, &ff); err != nil {
		return ff, status, err
	}

	return ff, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/franchise
func (f *FastId) SearchRanobeFranchise() (models.Franchise, int, error) {
	var ff models.Franchise

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(ranobe/) + ?(id) + 10(/franchise)
		concatination.Url(43+len(str_id), []string{
			constants.SITE, "ranobe/", str_id, "/franchise",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ff, status, err
	}

	if err := json.Unmarshal(data, &ff); err != nil {
		return ff, status, err
	}

	return ff, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchAnimeExternalLinks() ([]models.ExternalLinks, int, error) {
	var el []models.ExternalLinks

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(animes/) + ?(id) + 15(/external_links)
		concatination.Url(48+len(str_id), []string{
			constants.SITE, "animes/", str_id, "/external_links",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &el); err != nil {
		return nil, status, err
	}

	return el, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchMangaExternalLinks() ([]models.ExternalLinks, int, error) {
	var el []models.ExternalLinks

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(mangas/) + ?(id) + 15(/external_links)
		concatination.Url(48+len(str_id), []string{
			constants.SITE, "mangas/", str_id, "/external_links",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &el); err != nil {
		return nil, status, err
	}

	return el, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/external_links
func (f *FastId) SearchRanobeExternalLinks() ([]models.ExternalLinks, int, error) {
	var el []models.ExternalLinks

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(ranobe/) + ?(id) + 15(/external_links)
		concatination.Url(48+len(str_id), []string{
			constants.SITE, "ranobe/", str_id, "/external_links",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &el); err != nil {
		return nil, status, err
	}

	return el, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarAnime() ([]models.Animes, int, error) {
	var a []models.Animes

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(animes/) + ?(id) + 8(/similar)
		concatination.Url(41+len(str_id), []string{
			constants.SITE, "animes/", str_id, "/similar",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return nil, status, err
	}

	return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarManga() ([]models.Mangas, int, error) {
	var m []models.Mangas

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(mangas/) + ?(id) + 8(/similar)
		concatination.Url(41+len(str_id), []string{
			constants.SITE, "mangas/", str_id, "/similar",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/similar
func (f *FastId) SearchSimilarRanobe() ([]models.Mangas, int, error) {
	var m []models.Mangas

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(ranobe/) + ?(id) + 8(/similar)
		concatination.Url(41+len(str_id), []string{
			constants.SITE, "ranobe/", str_id, "/similar",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedAnime() ([]models.RelatedAnimes, int, error) {
	var a []models.RelatedAnimes

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(animes/) + ?(id) + 8(/related)
		concatination.Url(41+len(str_id), []string{
			constants.SITE, "animes/", str_id, "/related",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return nil, status, err
	}

	return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedManga() ([]models.RelatedMangas, int, error) {
	var m []models.RelatedMangas

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(mangas/) + ?(id) + 8(/related)
		concatination.Url(41+len(str_id), []string{
			constants.SITE, "mangas/", str_id, "/related",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/related
func (f *FastId) SearchRelatedRanobe() ([]models.RelatedMangas, int, error) {
	var m []models.RelatedMangas

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(ranobe/) + ?(id) + 8(/related)
		concatination.Url(41+len(str_id), []string{
			constants.SITE, "ranobe/", str_id, "/related",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// FIXME (ghostemeow): Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// Name: club name.
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (c *Configuration) SearchClubs(name string, r Result) ([]models.Clubs, int, error) {
	var cl []models.Clubs

	opt := r.OptionsOnlyPageLimitV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 13(clubs?search=) + ?(name) + 1(&) + ?(Result)
		concatination.Url(40+len(name)+len(opt), []string{
			constants.SITE, "clubs?search=", url.QueryEscape(name), "&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &cl); err != nil {
		return nil, status, err
	}

	return cl, status, nil
}

// FIXME (ghostemeow): The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubAnimes(r Result) ([]models.Animes, int, error) {
	var a []models.Animes

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 8(/animes?) + ?(Result)
		concatination.Url(40+len(str_id)+len(opt), []string{
			constants.SITE, "clubs/", str_id, "/animes?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return nil, status, err
	}

	return a, status, nil
}

// FIXME (ghostemeow): The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubMangas(r Result) ([]models.Mangas, int, error) {
	var m []models.Mangas

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 8(/mangas?) + ?(Result)
		concatination.Url(40+len(str_id)+len(opt), []string{
			constants.SITE, "clubs/", str_id, "/mangas?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// FIXME (ghostemeow): The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubRanobe(r Result) ([]models.Mangas, int, error) {
	var m []models.Mangas

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 8(/ranobe?) + ?(Result)
		concatination.Url(40+len(str_id)+len(opt), []string{
			constants.SITE, "clubs/", str_id, "/ranobe?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// FIXME (ghostemeow): The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 20 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubCharacters(r Result) ([]models.CharacterInfo, int, error) {
	var ci []models.CharacterInfo

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 12(/characters?) + ?(Result)
		concatination.Url(44+len(str_id)+len(opt), []string{
			constants.SITE, "clubs/", str_id, "/characters?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &ci); err != nil {
		return nil, status, err
	}

	return ci, status, nil
}

// FIXME (ghostemeow): The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubClubs(r Result) ([]models.Clubs, int, error) {
	var cc []models.Clubs

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 7(/clubs?) + ?(Result)
		concatination.Url(39+len(str_id)+len(opt), []string{
			constants.SITE, "clubs/", str_id, "/clubs?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &cc); err != nil {
		return nil, status, err
	}

	return cc, status, nil
}

// FIXME (ghostemeow): The limit does not work and always gives the maximum amount.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Page: 4 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubCollections(r Result) ([]models.ClubCollections, int, error) {
	var cc []models.ClubCollections

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 13(/collections?) + ?(Result)
		concatination.Url(45+len(str_id)+len(opt), []string{
			constants.SITE, "clubs/", str_id, "/collections?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &cc); err != nil {
		return nil, status, err
	}

	return cc, status, nil
}

// FIXME (ghostemeow): Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubMembers(r Result) ([]models.UserFriends, int, error) {
	var uf []models.UserFriends

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 9(/members?) + ?(Result)
		concatination.Url(41+len(str_id)+len(opt), []string{
			constants.SITE, "clubs/", str_id, "/members?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &uf); err != nil {
		return nil, status, err
	}

	return uf, status, nil
}

// FIXME (ghostemeow): Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 100 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) SearchClubImages(r Result) ([]models.ClubImages, int, error) {
	var cm []models.ClubImages

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 8(/images?) + ?(Result)
		concatination.Url(40+len(str_id)+len(opt), []string{
			constants.SITE, "clubs/", str_id, "/images?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &cm); err != nil {
		return nil, status, err
	}

	return cm, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) ClubJoin() (int, error) {
	str_id := strconv.Itoa(f.Id)

	_, status, err := request.NewPostRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 5(/join)
		concatination.Url(37+len(str_id), []string{
			constants.SITE, "clubs/", str_id, "/join",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return status, err
	}

	return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/club
func (f *FastId) ClubLeave() (int, error) {
	str_id := strconv.Itoa(f.Id)

	_, status, err := request.NewPostRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 6(clubs/) + ?(id) + 6(/leave)
		concatination.Url(38+len(str_id), []string{
			constants.SITE, "clubs/", str_id, "/leave",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return status, err
	}

	return status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// As a result, we return a complete list of all achievements.
//
// Next comes the filtering through "NekoSearch" and the error about obtaining
// specific achievements is already being processed there.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/achievements
func (f *FastId) SearchAchievement() ([]models.Achievements, int, error) {
	var a []models.Achievements

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 21(achievements?user_id=) + ?(id)
		concatination.Url(47+len(str_id), []string{
			constants.SITE, "achievements?user_id=", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return nil, status, err
	}

	return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/video
func (f *FastId) SearchAnimeVideos() ([]models.AnimeVideos, int, error) {
	var v []models.AnimeVideos

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		concatination.Url(40+len(str_id), []string{
			constants.SITE, "animes/", str_id, "/videos",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return nil, status, err
	}

	return v, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/roles
func (f *FastId) SearchAnimeRoles() ([]models.Roles, int, error) {
	var r []models.Roles

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(animes/) + ?(id) + 6(/roles)
		concatination.Url(39+len(str_id), []string{
			constants.SITE, "animes/", str_id, "/roles",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &r); err != nil {
		return nil, status, err
	}

	return r, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/roles
func (f *FastId) SearchMangaRoles() ([]models.Roles, int, error) {
	var r []models.Roles

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(mangas/) + ?(id) + 6(/roles)
		concatination.Url(39+len(str_id), []string{
			constants.SITE, "mangas/", str_id, "/roles",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &r); err != nil {
		return nil, status, err
	}

	return r, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/roles
func (f *FastId) SearchRanobeRoles() ([]models.Roles, int, error) {
	var r []models.Roles

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(ranobe/) + ?(id) + 6(/roles)
		concatination.Url(39+len(str_id), []string{
			constants.SITE, "ranobe/", str_id, "/roles",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &r); err != nil {
		return nil, status, err
	}

	return r, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/bans
func (c *Configuration) SearchBans() ([]models.Bans, int, error) {
	var b []models.Bans

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 4(bans)
		concatination.Url(30, []string{constants.SITE, "bans"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &b); err != nil {
		return nil, status, err
	}

	return b, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Censored: true, false;
//
// Set to false to allow hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/calendar
func (c *Configuration) SearchCalendar(r Result) ([]models.Calendar, int, error) {
	var ca []models.Calendar

	opt := r.OptionsCalendar()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 9(calendar?) + ?(Result)
		concatination.Url(35+len(opt), []string{constants.SITE, "calendar?", opt}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &ca); err != nil {
		return nil, status, err
	}

	return ca, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name:
//
// > GENRES_ANIME, GENRES_MANGA;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/genres
func (c *Configuration) SearchGenres(name string) ([]genres.Genres, int, error) {
	var g []genres.Genres

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 12(genres?kind=) + ?(name)
		concatination.Url(38+len(name), []string{constants.SITE, "genres?kind=", name}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &g); err != nil {
		return nil, status, err
	}

	return g, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/studios
func (c *Configuration) SearchStudios() ([]models.Studios, int, error) {
	var s []models.Studios

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 7(studios)
		concatination.Url(33, []string{constants.SITE, "studios"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return nil, status, err
	}

	return s, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/publishers
func (c *Configuration) SearchPublishers() ([]models.Publishers, int, error) {
	var p []models.Publishers

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 10(publishers)
		concatination.Url(36, []string{constants.SITE, "publishers"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &p); err != nil {
		return nil, status, err
	}

	return p, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/forums
func (c *Configuration) SearchForums() ([]models.Forums, int, error) {
	var f []models.Forums

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 6(forums)
		concatination.Url(32, []string{constants.SITE, "forums"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &f); err != nil {
		return nil, status, err
	}

	return f, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/add_remove_friend
func (f *FastId) AddFriend() (models.FriendRequest, int, error) {
	var ff models.FriendRequest

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewPostRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 8(friends/) + ?(id)
		concatination.Url(34+len(str_id), []string{
			constants.SITE, "friends/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ff, status, err
	}

	if err := json.Unmarshal(data, &ff); err != nil {
		return ff, status, err
	}

	return ff, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/add_remove_friend
func (f *FastId) RemoveFriend() (models.FriendRequest, int, error) {
	var ff models.FriendRequest

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewDeleteRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 8(friends/) + ?(id)
		concatination.Url(34+len(str_id), []string{
			constants.SITE, "friends/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ff, status, err
	}

	if err := json.Unmarshal(data, &ff); err != nil {
		return ff, status, err
	}

	return ff, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Show current user unread messages counts.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/messages
func (f *FastId) UserUnreadMessages() (models.UnreadMessages, int, error) {
	var um models.UnreadMessages

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancelAndBearer(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 6(users/) + ?(id) + 16(/unread_messages)
		concatination.Url(48+len(str_id), []string{
			constants.SITE, "users/", str_id, "/unread_messages",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return um, status, err
	}

	if err := json.Unmarshal(data, &um); err != nil {
		return um, status, err
	}

	return um, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// 'Options' settings:
//
//   - Page: 100000 maximum;
//
//   - Limit: 100 maximum;
//
//   - Type:
//
//     > MESSAGE_TYPE_INBOX, MESSAGE_TYPE_PRIVATE, MESSAGE_TYPE_SENT,
//     MESSAGE_TYPE_NEWS, MESSAGE_TYPE_NOTIFICATIONS;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/messages
func (f *FastId) UserMessages(r Result) ([]models.Messages, int, error) {
	var m []models.Messages

	opt := r.OptionsMessagesV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancelAndBearer(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 6(users/) + ?(id) + 10(/messages?) + ?(Result)
		concatination.Url(42+len(str_id)+len(opt), []string{
			constants.SITE, "users/", str_id, "/messages?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsAnime() (models.Constants, int, error) {
	var ca models.Constants

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 15(constants/anime)
		concatination.Url(41, []string{constants.SITE, "constants/anime"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ca, status, err
	}

	if err := json.Unmarshal(data, &ca); err != nil {
		return ca, status, err
	}

	return ca, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsManga() (models.Constants, int, error) {
	var cm models.Constants

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 15(constants/manga)
		concatination.Url(41, []string{constants.SITE, "constants/manga"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return cm, status, err
	}

	if err := json.Unmarshal(data, &cm); err != nil {
		return cm, status, err
	}

	return cm, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsUserRate() (models.ConstantsUserRate, int, error) {
	var ur models.ConstantsUserRate

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 19(constants/user_rate)
		concatination.Url(45, []string{constants.SITE, "constants/user_rate"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ur, status, err
	}

	if err := json.Unmarshal(data, &ur); err != nil {
		return ur, status, err
	}

	return ur, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsClub() (models.ConstantsClub, int, error) {
	var cc models.ConstantsClub

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 14(constants/club)
		concatination.Url(40, []string{constants.SITE, "constants/club"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return cc, status, err
	}

	if err := json.Unmarshal(data, &cc); err != nil {
		return cc, status, err
	}

	return cc, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/constants
func (c *Configuration) SearchConstantsSmileys() ([]models.ConstantsSmileys, int, error) {
	var cs []models.ConstantsSmileys

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 17(constants/smileys)
		concatination.Url(43, []string{constants.SITE, "constants/smileys"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &cs); err != nil {
		return nil, status, err
	}

	return cs, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//
//   - Limit: 50 maximum;
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
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/random
func (c *Configuration) RandomAnimes(r Result) ([]models.Animes, int, error) {
	var a []models.Animes

	opt := r.OptionsAnimeV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 20(animes?order=random&) + ?(Result)
		concatination.Url(46+len(opt), []string{
			constants.SITE, "animes?order=random&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return nil, status, err
	}

	return a, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//
//   - Limit: 50 maximum;
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
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/random
func (c *Configuration) RandomMangas(r Result) ([]models.Mangas, int, error) {
	var m []models.Mangas

	opt := r.OptionsMangaV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 20(mangas?order=random&) + ?(Result)
		concatination.Url(46+len(opt), []string{
			constants.SITE, "mangas?order=random&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//  - Limit: 50 maximum;
//  - Status:
//
//  > MANGA_STATUS_ANONS, MANGA_STATUS_ONGOING, MANGA_STATUS_RELEASED, MANGA_STATUS_PAUSED,
//  MANGA_STATUS_DISCONTINUED, MANGA_STATUS_ANONS_NOT_EQUAL, MANGA_STATUS_ONGOING_NOT_EQUAL,
//  MANGA_STATUS_RELEASED_NOT_EQUAL, MANGA_STATUS_PAUSED_NOT_EQUAL, MANGA_STATUS_DISCONTINUED_NOT_EQUAL;
//
//  - Season:
//
//  > SEASON_198x, SEASON_199x, SEASON_2000_2010, SEASON_2010_2014,
//  SEASON_2015_2019, SEASON_2020_2021, SEASON_2022, SEASON_2023,
//  SEASON_198x_NOT_EQUAL, SEASON_199x_NOT_EQUAL, SEASON_2000_2010_NOT_EQUAL,
//  SEASON_2010_2014_NOT_EQUAL, SEASON_2015_2019_NOT_EQUAL,
//  SEASON_2020_2021_NOT_EQUAL, SEASON_2022_NOT_EQUAL, SEASON_2023_NOT_EQUAL;
//
//  - Mylist:
//
//  > MY_LIST_PLANNED, MY_LIST_WATCHING, MY_LIST_REWATCHING,
//  MY_LIST_COMPLETED, MY_LIST_ON_HOLD, MY_LIST_DROPPED;

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
// Set to true to discard hentai, yaoi and yuri.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/random
func (c *Configuration) RandomRanobes(r Result) ([]models.Mangas, int, error) {
	var m []models.Mangas

	opt := r.OptionsRanobeV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 20(ranobe?order=random&) + ?(Result)
		concatination.Url(46+len(opt), []string{
			constants.SITE, "ranobe?order=random&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, status, err
	}

	return m, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/character
func (f *FastId) SearchCharacter() (models.Character, int, error) {
	var ch models.Character

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 11(characters/) + ?(id)
		concatination.Url(37+len(str_id), []string{
			constants.SITE, "characters/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ch, status, err
	}

	if err := json.Unmarshal(data, &ch); err != nil {
		return ch, status, err
	}

	return ch, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Name: character name.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/character
func (c *Configuration) SearchCharacters(name string) ([]models.CharacterInfo, int, error) {
	var ci []models.CharacterInfo

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 25(characters/search?search=) + ?(name)
		concatination.Url(51+len(name), []string{constants.SITE,
			"characters/search?search=", url.QueryEscape(name)}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &ci); err != nil {
		return nil, status, err
	}

	return ci, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/people
func (f *FastId) SearchPeople() (models.People, int, error) {
	var p models.People

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(people/) + ?(id)
		concatination.Url(33+len(str_id), []string{
			constants.SITE, "people/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return p, status, err
	}

	if err := json.Unmarshal(data, &p); err != nil {
		return p, status, err
	}

	return p, status, nil
}

// FIXME (ghostemeow): Page and limit not supprted, idk why. Check later.
//
// Only the application needs to be specified in SetConfiguration().
//
// Name: people name.
//
// 'Options' settings:
//
//   - Kind:
//
//     > PEOPLE_KIND_SEYU, PEOPLE_KIND_MANGAKA, PEOPLE_KIND_PRODUCER;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/people
func (c *Configuration) SearchPeoples(name string, r Result) ([]models.AllPeople, int, error) {
	var ap []models.AllPeople

	opt := r.OptionsPeopleV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		concatination.Url(48+len(name)+len(opt), []string{
			constants.SITE, "people/search?search=", url.QueryEscape(name), "&", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &ap); err != nil {
		return nil, status, err
	}

	return ap, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Linked_type:
//
// > FAVORITES_LINKED_TYPE_ANIME, FAVORITES_LINKED_TYPE_MANGA,
// FAVORITES_LINKED_TYPE_RANOBE, FAVORITES_LINKED_TYPE_PERSON,
// FAVORITES_LINKED_TYPE_CHARACTER;
//
// Kind(required when Linked_type is Person):
//
// > FAVORITES_KIND_COMMON, FAVORITES_KIND_SEYU, FAVORITES_KIND_MANGAKA,
// FAVORITES_KIND_PRODUCER, FAVORITES_KIND_PERSON;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesCreate(linked_type, kind string) (models.Favorites, int, error) {
	var fa models.Favorites

	if linked_type != constants.FAVORITES_LINKED_TYPE_PERSON {
		kind = ""
	}

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewPostRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 10(favorites/) + ?(linked_type) + 1(/) + ?(id) + 1(/) + ?(kind)
		concatination.Url(38+len(linked_type)+len(str_id)+len(kind), []string{
			constants.SITE, "favorites/", linked_type, "/", str_id, "/", kind,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return fa, status, err
	}

	if err := json.Unmarshal(data, &fa); err != nil {
		return fa, status, err
	}

	return fa, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Linked_type:
//
// > FAVORITES_LINKED_TYPE_ANIME, FAVORITES_LINKED_TYPE_MANGA,
// FAVORITES_LINKED_TYPE_RANOBE, FAVORITES_LINKED_TYPE_PERSON,
// FAVORITES_LINKED_TYPE_CHARACTER;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesDelete(linked_type string) (models.Favorites, int, error) {
	var ff models.Favorites

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewDeleteRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 10(favorites/) + ?(linked_type) + 1(/) + ?(id)
		concatination.Url(37+len(linked_type)+len(str_id), []string{
			constants.SITE, "favorites/", linked_type, "/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return ff, status, err
	}

	if err := json.Unmarshal(data, &ff); err != nil {
		return ff, status, err
	}

	return ff, status, nil
}

// FIXME (ghostemeow): https://github.com/ghostemeow/goshikimori/issues/14
//
// In SetConfiguration(), you must specify the application and the token.
//
// Position: a new position on the list, it starts from 0.
//
// You can only get a StatusCode.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/favorites
func (f *FastId) FavoritesReorder(position int) (int, error) {
	str_id := strconv.Itoa(f.Id)

	_, status, err := request.NewReorderPostRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 10(favorites/) + ?(id) + 8(/reorder)
		concatination.Url(44+len(str_id), []string{
			constants.SITE, "favorites/", str_id, "/reorder",
		}), position, constants.MAX_EXPECTATION,
	)
	if err != nil {
		return status, err
	}

	return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/ignore
func (f *FastId) AddIgnoreUser() (models.IgnoreUser, int, error) {
	var i models.IgnoreUser

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewPostRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 9(v2/users/) + ?(id) + 7(/ignore)
		concatination.Url(42+len(str_id), []string{
			constants.SITE, "v2/users/", str_id, "/ignore",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return i, status, err
	}

	if err := json.Unmarshal(data, &i); err != nil {
		return i, status, err
	}

	return i, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/ignore
func (f *FastId) RemoveIgnoreUser() (models.IgnoreUser, int, error) {
	var i models.IgnoreUser

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewDeleteRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 9(v2/users/) + ?(id) + 7(/ignore)
		concatination.Url(42+len(str_id), []string{
			constants.SITE, "v2/users/", str_id, "/ignore",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return i, status, err
	}

	if err := json.Unmarshal(data, &i); err != nil {
		return i, status, err
	}

	return i, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/dialogs
func (c *Configuration) Dialogs() ([]models.Dialogs, int, error) {
	var d []models.Dialogs

	data, status, err := request.NewGetRequestWithCancelAndBearer(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 7(dialogs)
		concatination.Url(33, []string{constants.SITE, "dialogs"}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &d); err != nil {
		return nil, status, err
	}

	return d, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// When using FastIdUser()/SetFastId(), specify the user's nickname (not your own).
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/dialogs
func (f *FastId) SearchDialogs() ([]models.SearchDialogs, int, error) {
	var sd []models.SearchDialogs

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancelAndBearer(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 8(dialogs/) + ?(id)
		concatination.Url(34+len(str_id), []string{
			constants.SITE, "dialogs/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &sd); err != nil {
		return nil, status, err
	}

	return sd, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// When using FastIdUser()/SetFastId(), specify the user's nickname (not your own).
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/dialogs
func (f *FastId) DeleteDialogs() (models.FriendRequest, int, error) {
	var fr models.FriendRequest

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewDeleteRequestWithCancel(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 8(dialogs/) + ?(id)
		concatination.Url(34+len(str_id), []string{
			constants.SITE, "dialogs/", str_id,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return fr, status, err
	}

	if err := json.Unmarshal(data, &fr); err != nil {
		// Original error message from models/v1.
		return fr, status, errors.New("не найдено ни одного сообщения для удаления")
	}

	return fr, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/user
func (f *FastId) UserBriefInfo() (models.Info, int, error) {
	var i models.Info

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 6(users/) + ?(id) + 5(/info)
		concatination.Url(37+len(str_id), []string{
			constants.SITE, "users/", str_id, "/info",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return i, status, err
	}

	if err := json.Unmarshal(data, &i); err != nil {
		return i, status, err
	}

	return i, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// https://github.com/ghostemeow/goshikimori/issues/26
func (c *Configuration) SignOut() ([]byte, int, error) {
	data, status, err := request.NewPostRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 14(users/sign_out)
		concatination.Url(40, []string{
			constants.SITE, "users/sign_out",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return data, status, err
	}

	return data, status, nil
}

// If we get a json unmarshal type error, it is a server-side error, namely:
// ["PG::DiskFull: ERROR:  could not resize shared memory segment \"/PostgreSQL.1559179908\" to 16777216 bytes: No space left on device\n"]
//
// Only the application needs to be specified in SetConfiguration().
//
// Users having at least 1 completed animes and active during last month.
//
// Time to complete request increased to 40 seconds. Too big request.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/active_users
func (c *Configuration) ActiveUsers() ([]int, int, error) {
	ids := make([]int, 0)

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 18(stats/active_users)
		concatination.Url(44, []string{
			constants.SITE, "stats/active_users",
		}), constants.CUSTOM_MAX_EXPECTATION_ACTIVE_USERS,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &ids); err != nil {
		return nil, status, err
	}

	return ids, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsAnime(r Result) ([]models.Topics, int, error) {
	var t []models.Topics

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(animes/) + ?(id) + 8(/topics?) + ?(Result)
		concatination.Url(41+len(str_id)+len(opt), []string{
			constants.SITE, "animes/", str_id, "/topics?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return nil, status, err
	}

	return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsManga(r Result) ([]models.Topics, int, error) {
	var t []models.Topics

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(mangas/) + ?(id) + 8(/topics?) + ?(Result)
		concatination.Url(41+len(str_id)+len(opt), []string{
			constants.SITE, "mangas/", str_id, "/topics?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return nil, status, err
	}

	return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (f *FastId) SearchTopicsRanobe(r Result) ([]models.Topics, int, error) {
	var t []models.Topics

	opt := r.OptionsOnlyPageLimitV2()
	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancel(
		f.Conf.Application,
		// 26(constants.SITE) + 7(ranobe/) + ?(id) + 8(/topics?) + ?(Result)
		concatination.Url(41+len(str_id)+len(opt), []string{
			constants.SITE, "ranobe/", str_id, "/topics?", opt,
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return nil, status, err
	}

	return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//
//   - Page: 100000 maximum;
//
//   - Limit: 30 maximum;
//
//   - Forum:
//
//     > TOPIC_FORUM_ALL, TOPIC_FORUM_COSPLAY, TOPIC_FORUM_ANIMANGA, TOPIC_FORUM_constants.SITE,
//     TOPIC_FORUM_GAMES, TOPIC_FORUM_VN, TOPIC_FORUM_CONTEST, TOPIC_FORUM_OFFTOPIC,
//     TOPIC_FORUM_CLUBS, TOPIC_FORUM_MYCLUBS, TOPIC_FORUM_CRITIQUES,
//     TOPIC_FORUM_NEWS, TOPIC_FORUM_COLLECTIONS, TOPIC_FORUM_ARTICLES;
//
//   - Linked_id: number without limit;
//
//   - Linked_type:
//
//     > TOPIC_LINKED_TYPE_ANIME, TOPIC_LINKED_TYPE_MANGA, TOPIC_LINKED_TYPE_RANOBE,
//     TOPIC_LINKED_TYPE_CHARACTER, TOPIC_LINKED_TYPE_PERSON, TOPIC_LINKED_TYPE_CLUB,
//     TOPIC_LINKED_TYPE_CLUBPAGE, TOPIC_LINKED_TYPE_CRITIQUE, TOPIC_LINKED_TYPE_REVIEW,
//     TOPIC_LINKED_TYPE_CONTEST, TOPIC_LINKED_TYPE_COSPLAYGALLYRY,
//     TOPIC_LINKED_TYPE_COLLECTION, TOPIC_LINKED_TYPE_ARTICLE;
//
// REMARK: linked_id and linked_type are only used together.
//
//   - Type: not supported;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopics(r Result) ([]models.Topics, int, error) {
	var t []models.Topics

	opt := r.OptionsTopicsV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 7(topics?) + ?(Result)
		concatination.Url(33+len(opt), []string{constants.SITE, "topics?", opt}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return nil, status, err
	}

	return t, status, nil
}

// FIXME (ghostemeow): Limit always returns +1 of the given number.
//
// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Page: 100000 maximum;
//   - Limit: 30 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsUpdates(r Result) ([]models.TopicsUpdates, int, error) {
	var t []models.TopicsUpdates

	opt := r.OptionsOnlyPageLimitV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 15(topics/updates?) + ?(Result)
		concatination.Url(41+len(opt), []string{constants.SITE, "topics/updates?", opt}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return nil, status, err
	}

	return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// 'Options' settings:
//   - Limit: 10 maximum;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsHot(r Result) ([]models.Topics, int, error) {
	var t []models.Topics

	opt := r.OptionsTopicsHotV2()

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 11(topics/hot?) + ?(Result)
		concatination.Url(37+len(opt), []string{constants.SITE, "topics/hot?", opt}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return nil, status, err
	}

	return t, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (c *Configuration) SearchTopicsId(id int) (models.TopicsId, int, error) {
	var t models.TopicsId

	str_id := strconv.Itoa(id)

	data, status, err := request.NewGetRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + 7(topics/) + ?(id)
		concatination.Url(33+len(str_id), []string{constants.SITE, "topics/", str_id}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return t, status, err
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return t, status, err
	}

	return t, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (c *Configuration) AddIgnoreTopic(id int) (models.IgnoreTopic, int, error) {
	var i models.IgnoreTopic

	str_id := strconv.Itoa(id)

	data, status, err := request.NewPostRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 10(v2/topics/) + ?(id) + 7(/ignore)
		concatination.Url(43+len(str_id), []string{
			constants.SITE, "v2/topics/", str_id, "/ignore",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return i, 0, err
	}

	if err := json.Unmarshal(data, &i); err != nil {
		return i, status, err
	}

	return i, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// You can find the ID in functions such as: SearchTopics(), SearchTopicsUpdates(), SearchTopicsHot()
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/topics
func (c *Configuration) RemoveIgnoreTopic(id int) (models.IgnoreTopic, int, error) {
	var i models.IgnoreTopic

	str_id := strconv.Itoa(id)

	data, status, err := request.NewDeleteRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 10(v2/topics/) + ?(id) + 7(/ignore)
		concatination.Url(43+len(str_id), []string{
			constants.SITE, "v2/topics/", str_id, "/ignore",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return i, 0, err
	}

	if err := json.Unmarshal(data, &i); err != nil {
		return i, status, err
	}

	return i, status, nil
}

// Only the application needs to be specified in SetConfiguration().
//
// Schema: customized request.
//
// How to use and all the information you need [here].
//
// [here]: https://github.com/ghostemeow/goshikimori/blob/master/examples/GraphQL.md
func (c *Configuration) SearchGraphql(schema string) (models.GraphQL, int, error) {
	var g models.GraphQL

	data, status, err := request.NewGraphQLPostRequestWithCancel(
		c.Application,
		// 26(constants.SITE) + ?(schema)
		concatination.Url(26+len(schema), []string{constants.SITE, schema}),
		constants.CUSTOM_MAX_EXPECTATION_GRAPHQL,
	)
	if err != nil {
		return g, status, err
	}

	if err := json.Unmarshal(data, &g); err != nil {
		return g, status, err
	}

	return g, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Id: message id.
//
// Ignore:
//   - Linked_type: nil;
//   - Linked: nil;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/message
func (c *Configuration) ReadMessage(id int) (models.Messages, int, error) {
	var m models.Messages

	str_id := strconv.Itoa(id)

	data, status, err := request.NewGetRequestWithCancelAndBearer(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 9(messages/) + ?(id)
		concatination.Url(35+len(str_id), []string{constants.SITE, "messages/", str_id}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return m, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return m, status, err
	}

	return m, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// From_id: your Id.
//
// To_id: the Id of the person you want to send the message to.
//
// Message: message text.
//
// Returns a status of 201.
//
// Ignore:
//   - Linked_type: nil;
//   - Linked: nil;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/message
func (c *Configuration) SendMessage(from_id, to_id int, message string) (models.Messages, int, error) {
	var m models.Messages

	data, status, err := request.NewSendMessagePostRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 8(messages)
		concatination.Url(34, []string{constants.SITE, "messages"}),
		message, from_id, to_id, constants.MAX_EXPECTATION,
	)
	if err != nil {
		return m, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return m, status, err
	}

	return m, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Id: message id.
//
// Message: message text.
//
// Ignore:
//   - Linked_type: nil;
//   - Linked: nil;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/message
func (c *Configuration) ChangeMessage(id int, message string) (models.Messages, int, error) {
	var m models.Messages

	str_id := strconv.Itoa(id)

	data, status, err := request.NewChangeMessagePutRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 9(messages/) + ?(id)
		concatination.Url(35+len(str_id), []string{constants.SITE, "messages/", str_id}),
		message, constants.MAX_EXPECTATION,
	)
	if err != nil {
		return m, status, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return m, status, err
	}

	return m, status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Id: message id.
//
// Only status 204 is returned.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/message
func (c *Configuration) DeleteMessage(id int) (int, error) {
	str_id := strconv.Itoa(id)

	_, status, err := request.NewDeleteMessageDeleteRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 9(messages/) + ?(id)
		concatination.Url(35+len(str_id), []string{constants.SITE, "messages/", str_id}),
		constants.MAX_EXPECTATION,
	)
	if err != nil {
		return status, err
	}

	return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Ids: array of ids converted to a string.
//
// Is_read: mark a message as read or unread.
//
// 'Is_read' settings:
//   - 1 (read)
//   - 0 (unread)
//
// Only status 200 is returned.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/message
func (c *Configuration) MarkReadMessages(ids string, is_read int) (int, error) {
	_, status, err := request.NewMarkReadPostRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 18(messages/mark_read)
		concatination.Url(44, []string{
			constants.SITE, "messages/mark_read",
		}), ids, is_read, constants.MAX_EXPECTATION,
	)
	if err != nil {
		return status, err
	}

	return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Name: unread message type.
//
// 'Name' settings:
//
// > UNREAD_MESSAGES_IDS_NEWS, UNREAD_MESSAGES_IDS_MESSAGES,
// UNREAD_MESSAGES_IDS_NOTIFICATIONS;
//
// Empty array to be filled with ids for messages.
func (f *FastId) UnreadMessagesIds(name string) ([]int, int, error) {
	var um models.UnreadMessages

	str_id := strconv.Itoa(f.Id)

	data, status, err := request.NewGetRequestWithCancelAndBearer(
		f.Conf.Application, f.Conf.AccessToken,
		// 26(constants.SITE) + 6(users/) + ?(id) + 16(/unread_messages)
		concatination.Url(48+len(str_id), []string{
			constants.SITE, "users/", str_id, "/unread_messages",
		}), constants.MAX_EXPECTATION,
	)
	if err != nil {
		return nil, status, err
	}

	if err := json.Unmarshal(data, &um); err != nil {
		return nil, status, err
	}

	switch name {
	case "messages":
		if um.Messages == 0 {
			return nil, status, errors.New("unread messages not found")
		}
		return make([]int, um.Messages), status, nil
	case "news":
		if um.News == 0 {
			return nil, status, errors.New("unread news not found")
		}
		return make([]int, um.News), status, nil
	case "notifications":
		if um.Notifications == 0 {
			return nil, status, errors.New("unread notifications not found")
		}
		return make([]int, um.Notifications), status, nil
	default:
		return nil, status, errors.New("wrong name... try messages, news or notifications")
	}
}

// In SetConfiguration(), you must specify the application and the token.
//
// Name: type in the mail.
//
// Returns a status of 200.
//
// 'Name' settings:
//
// > UNREAD_MESSAGES_IDS_NEWS, UNREAD_MESSAGES_IDS_NOTIFICATIONS;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/message
func (c *Configuration) ReadAllMessages(name string) (int, error) {
	_, status, err := request.NewReadDeleteAllPostRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 17(messages/read_all)
		concatination.Url(43, []string{constants.SITE, "messages/read_all"}),
		name, constants.MAX_EXPECTATION,
	)
	if err != nil {
		return status, err
	}

	return status, nil
}

// In SetConfiguration(), you must specify the application and the token.
//
// Name: type in the mail.
//
// Returns a status of 200.
//
// 'Name' settings:
//
// > UNREAD_MESSAGES_IDS_NEWS, UNREAD_MESSAGES_IDS_NOTIFICATIONS;
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/message
func (c *Configuration) DeleteAllMessages(name string) (int, error) {
	_, status, err := request.NewReadDeleteAllPostRequestWithCancel(
		c.Application, c.AccessToken,
		// 26(constants.SITE) + 19(messages/delete_all)
		concatination.Url(45, []string{constants.SITE, "messages/delete_all"}),
		name, constants.MAX_EXPECTATION,
	)
	if err != nil {
		return 0, err
	}

	return status, nil
}
