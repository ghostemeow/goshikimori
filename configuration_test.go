package goshikimori

import (
	"testing"

	"github.com/ghostemeow/goshikimori/constants"
)

func TestOptionsTopics(t *testing.T) {
	empty := Options{}
	if empty.OptionsTopics() == "" {
		t.Log("Empty OptionsTopics passed")
	} else {
		t.Error("Empty OptionsTopics failed")
	}

	big := Options{
		Page: 100002, Limit: 32,
		Linked_id: 222222222, Linked_type: constants.TOPIC_LINKED_TYPE_ANIME,
	}
	if big.OptionsTopics() == "linked_id=222222222&linked_type=Anime" {
		t.Log("Big OptionsTopics passed")
	} else {
		t.Error("Big OptionsTopics failed")
	}

	zero := Options{
		Page: 0, Limit: 0,
		Linked_id: 0, Linked_type: constants.TOPIC_LINKED_TYPE_ANIME,
	}
	if zero.OptionsTopics() == "" {
		t.Log("Zero OptionsTopics passed")
	} else {
		t.Error("Zero OptionsTopics failed")
	}

	negative := Options{
		Page: -1, Limit: -1,
		Linked_id: -1, Linked_type: constants.TOPIC_LINKED_TYPE_ANIME,
	}
	if negative.OptionsTopics() == "" {
		t.Log("Negative OptionsTopics passed")
	} else {
		t.Error("Negative OptionsTopics failed")
	}

	normal_one := Options{
		Page: 5, Limit: 10, Forum: constants.TOPIC_FORUM_ANIMANGA,
		Linked_id: 342908, Linked_type: constants.TOPIC_LINKED_TYPE_ANIME,
	}
	if normal_one.OptionsTopics() == "forum=animanga&limit=10&linked_id=342908&linked_type=Anime&page=5" {
		t.Log("Normal-one OptionsTopics passed")
	} else {
		t.Error("Normal-one OptionsTopics failed")
	}

	normal_two := Options{
		Page: 3, Limit: 8, Forum: constants.TOPIC_FORUM_CLUBS,
		Linked_id: 2323, Linked_type: constants.TOPIC_LINKED_TYPE_MANGA,
	}
	if normal_two.OptionsTopics() == "forum=clubs&limit=8&linked_id=2323&linked_type=Manga&page=3" {
		t.Log("Normal-two OptionsTopics passed")
	} else {
		t.Error("Normal-two OptionsTopics failed")
	}
}

func TestOptionsTopicsV2(t *testing.T) {
	normal_one := Options{
		Page: 5, Limit: 10, Forum: constants.TOPIC_FORUM_ANIMANGA,
		Linked_id: 342908, Linked_type: constants.TOPIC_LINKED_TYPE_ANIME,
	}
	if normal_one.OptionsTopicsV2() == "page=5&limit=10&forum=animanga&linked_id=342908&linked_type=Anime" {
		t.Log("Normal-one OptionsTopicsV2 passed")
	} else {
		t.Error("Normal-one OptionsTopicsV2 failed")
	}

	normal_two := Options{
		Page: 3, Limit: 8, Forum: constants.TOPIC_FORUM_CLUBS,
		Linked_id: 2323, Linked_type: constants.TOPIC_LINKED_TYPE_MANGA,
	}
	if normal_two.OptionsTopicsV2() == "page=3&limit=8&forum=clubs&linked_id=2323&linked_type=Manga" {
		t.Log("Normal-two OptionsTopicsV2 passed")
	} else {
		t.Error("Normal-two OptionsTopicsV2 failed")
	}
}

func TestOptionsMessages(t *testing.T) {
	empty := Options{}
	if empty.OptionsMessages() == "type=news" {
		t.Log("Empty OptionsMessages passed")
	} else {
		t.Error("Empty OptionsMessages failed")
	}

	big := Options{Type: constants.MESSAGE_TYPE_NEWS, Page: 100002, Limit: 102}
	if big.OptionsMessages() == "type=news" {
		t.Log("Big OptionsMessages passed")
	} else {
		t.Error("Big OptionsMessages failed")
	}

	zero := Options{Type: constants.MESSAGE_TYPE_NEWS, Page: 0, Limit: 0}
	if zero.OptionsMessages() == "type=news" {
		t.Log("Zero OptionsMessages passed")
	} else {
		t.Error("Zero OptionsMessages failed")
	}

	negative := Options{Type: constants.MESSAGE_TYPE_NEWS, Page: -1, Limit: -1}
	if negative.OptionsMessages() == "type=news" {
		t.Log("Negative OptionsMessages passed")
	} else {
		t.Error("Negative OptionsMessages failed")
	}

	normal := Options{Type: constants.MESSAGE_TYPE_PRIVATE, Page: 2, Limit: 10}
	if normal.OptionsMessages() == "limit=10&page=2&type=private" {
		t.Log("Normal OptionsMessages passed")
	} else {
		t.Error("Normal OptionsMessages failed")
	}
}

func TestOptionsMessagesV2(t *testing.T) {
	normal := Options{Type: constants.MESSAGE_TYPE_PRIVATE, Page: 2, Limit: 10}
	if normal.OptionsMessagesV2() == "page=2&limit=10&type=private" {
		t.Log("Normal OptionsMessagesV2 passed")
	} else {
		t.Error("Normal OptionsMessagesV2 failed")
	}
}

func TestOptionsUserHistory(t *testing.T) {
	empty := Options{}
	if empty.OptionsUserHistory() == "" {
		t.Log("Empty OptionsUserHistory passed")
	} else {
		t.Error("Empty OptionsUserHistory failed")
	}

	big := Options{Page: 100002, Limit: 102}
	if big.OptionsUserHistory() == "" {
		t.Log("Big OptionsUserHistory passed")
	} else {
		t.Error("Big OptionsUserHistory failed")
	}

	zero := Options{Page: 0, Limit: 0, Target_id: 0}
	if zero.OptionsUserHistory() == "" {
		t.Log("Zero OptionsUserHistory passed")
	} else {
		t.Error("Zero OptionsUserHistory failed")
	}

	negative := Options{Page: -1, Limit: -1, Target_id: -1}
	if negative.OptionsUserHistory() == "" {
		t.Log("Negative OptionsUserHistory passed")
	} else {
		t.Error("Negative OptionsUserHistory failed")
	}

	normal := Options{Page: 3, Limit: 20, Target_id: 1337, Target_type: constants.TARGET_TYPE_MANGA}
	if normal.OptionsUserHistory() == "limit=20&page=3&target_id=1337&target_type=Manga" {
		t.Log("Zero OptionsUserHistory passed")
	} else {
		t.Error("Zero OptionsUserHistory failed")
	}
}

func TestOptionsUserHistoryV2(t *testing.T) {
	normal := Options{Page: 3, Limit: 20, Target_id: 1337, Target_type: constants.TARGET_TYPE_MANGA}
	if normal.OptionsUserHistoryV2() == "page=3&limit=20&target_type=Manga&target_id=1337" {
		t.Log("Zero OptionsUserHistoryV2 passed")
	} else {
		t.Error("Zero OptionsUserHistoryV2 failed")
	}
}

func TestOptionsOnlyPageLimit(t *testing.T) {
	empty := Options{}
	if empty.OptionsOnlyPageLimit(100000, 100) == "" {
		t.Log("Empty OptionsOnlyPageLimit passed")
	} else {
		t.Error("Empty OptionsOnlyPageLimit failed")
	}

	big := Options{Page: 100002, Limit: 102}
	if big.OptionsOnlyPageLimit(100000, 100) == "" {
		t.Log("Big OptionsOnlyPageLimit passed")
	} else {
		t.Error("Big OptionsOnlyPageLimit failed")
	}

	zero := Options{Page: 0, Limit: 0}
	if zero.OptionsOnlyPageLimit(100000, 100) == "" {
		t.Log("Zero OptionsOnlyPageLimit passed")
	} else {
		t.Error("Zero OptionsOnlyPageLimit failed")
	}

	negative := Options{Page: -1, Limit: -1}
	if negative.OptionsOnlyPageLimit(100000, 100) == "" {
		t.Log("Negative OptionsOnlyPageLimit passed")
	} else {
		t.Error("Negative OptionsOnlyPageLimit failed")
	}

	normal := Options{Page: 7, Limit: 37}
	if normal.OptionsOnlyPageLimit(100000, 100) == "limit=37&page=7" {
		t.Log("Normal OptionsOnlyPageLimit passed")
	} else {
		t.Error("Normal OptionsOnlyPageLimit failed")
	}
}

func TestOptionsOnlyPageLimitV2(t *testing.T) {
	normal := Options{Page: 7, Limit: 37}
	if normal.OptionsOnlyPageLimitV2() == "page=7&limit=37" {
		t.Log("Normal OptionsOnlyPageLimitV2 passed")
	} else {
		t.Error("Normal OptionsOnlyPageLimitV2 failed")
	}
}

func TestOptionsAnime(t *testing.T) {
	empty := Options{}
	if empty.OptionsAnime() == "censored=false" {
		t.Log("Empty OptionsAnime passed")
	} else {
		t.Error("Empty OptionsAnime failed")
	}

	big := Options{
		Page: 100002, Limit: 52, Score: 12, Censored: false, Genre_v2: []int{111111111},
	}
	if big.OptionsAnime() == "censored=false" {
		t.Log("Big OptionsAnime passed")
	} else {
		t.Error("Big OptionsAnime failed")
	}

	zero := Options{
		Page: 0, Limit: 0, Score: 0,
		Censored: false, Genre_v2: []int{0},
	}
	if zero.OptionsAnime() == "censored=false" {
		t.Log("Zero OptionsAnime passed")
	} else {
		t.Error("Zero OptionsAnime failed")
	}

	negative := Options{
		Page: -1, Limit: -1, Score: -1,
		Censored: false, Genre_v2: []int{-1},
	}
	if negative.OptionsAnime() == "censored=false" {
		t.Log("Negative OptionsAnime passed")
	} else {
		t.Error("Negative OptionsAnime failed")
	}

	normal := Options{
		Page: 2, Limit: 12, Order: constants.ANIME_ORDER_ID, Kind: constants.ANIME_KIND_TV,
		Status: constants.ANIME_STATUS_RELEASED, Season: constants.SEASON_199x,
		Score: 8, Rating: constants.ANIME_RATING_R, Duration: constants.ANIME_DURATION_D,
		Censored: true, Mylist: constants.MY_LIST_WATCHING, Genre_v2: []int{539, 539},
	}
	if normal.OptionsAnime() == "censored=true&duration=D&genre_v2=539-Erotica&kind=tv&limit=12&mylist=watching&order=id&page=2&rating=r&score=8&season=199x&status=released" {
		t.Log("Normal OptionsAnime passed")
	} else {
		t.Error("Normal OptionsAnime failed")
	}
}

func TestOptionsAnimeV2(t *testing.T) {
	normal := Options{
		Page: 2, Limit: 12, Order: constants.ANIME_ORDER_ID, Kind: constants.ANIME_KIND_TV,
		Status: constants.ANIME_STATUS_RELEASED, Season: constants.SEASON_199x,
		Score: 8, Rating: constants.ANIME_RATING_R, Duration: constants.ANIME_DURATION_D,
		Censored: true, Mylist: constants.MY_LIST_WATCHING, Genre_v2: []int{539, 539},
	}
	if normal.OptionsAnimeV2() == "page=2&limit=12&score=8&order=id&kind=tv&status=released&season=199x&rating=r&duration=D&mylist=watching&censored=true&genre_v2=539-Erotica" {
		t.Log("Normal OptionsAnimeV2 passed")
	} else {
		t.Error("Normal OptionsAnimeV2 failed")
	}
}

func TestOptionsManga(t *testing.T) {
	empty := Options{}
	if empty.OptionsManga() == "censored=false" {
		t.Log("Empty OptionsManga passed")
	} else {
		t.Error("Empty OptionsManga failed")
	}

	big := Options{
		Page: 100002, Limit: 52, Score: 1111111111,
		Censored: false, Genre_v2: []int{111111111111},
	}
	if big.OptionsManga() == "censored=false" {
		t.Log("Big OptionsManga passed")
	} else {
		t.Error("Big OptionsManga failed")
	}

	zero := Options{
		Page: 0, Limit: 0, Score: 0, Censored: false, Genre_v2: []int{0},
	}
	if zero.OptionsManga() == "censored=false" {
		t.Log("Zero OptionsManga passed")
	} else {
		t.Error("Zero OptionsManga failed")
	}

	negative := Options{
		Page: -1, Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
	}
	if negative.OptionsManga() == "censored=false" {
		t.Log("Negative OptionsManga passed")
	} else {
		t.Error("Negative OptionsManga failed")
	}

	normal := Options{
		Page: 4, Limit: 5, Order: constants.MANGA_ORDER_ID, Kind: constants.MANGA_KIND_MANGA,
		Status: constants.MANGA_STATUS_ANONS, Season: constants.SEASON_198x,
		Score: 7, Censored: true, Mylist: constants.MY_LIST_PLANNED, Genre_v2: []int{540, 540},
	}
	if normal.OptionsManga() == "censored=true&genre_v2=540-Erotica&kind=manga&limit=5&mylist=planned&order=id&page=4&score=7&season=198x&status=anons" {
		t.Log("Normal OptionsManga passed")
	} else {
		t.Error("Normal OptionsManga failed")
	}
}

func TestOptionsMangaV2(t *testing.T) {
	normal := Options{
		Page: 4, Limit: 5, Order: constants.MANGA_ORDER_ID, Kind: constants.MANGA_KIND_MANGA,
		Status: constants.MANGA_STATUS_ANONS, Season: constants.SEASON_198x,
		Score: 7, Censored: true, Mylist: constants.MY_LIST_PLANNED, Genre_v2: []int{540, 540},
	}
	if normal.OptionsMangaV2() == "page=4&limit=5&score=7&order=id&kind=manga&status=anons&season=198x&mylist=planned&censored=true" {
		t.Log("Normal OptionsMangaV2 passed")
	} else {
		t.Error("Normal OptionsMangaV2 failed")
	}
}

func TestOptionsRanobe(t *testing.T) {
	empty := Options{}
	if empty.OptionsRanobe() == "censored=false" {
		t.Log("Empty OptionsRanobe passed")
	} else {
		t.Error("Empty OptionsRanobe failed")
	}

	big := Options{
		Page: 100002, Limit: 52, Score: 1111111111,
		Censored: false, Genre_v2: []int{111111111111},
	}
	if big.OptionsRanobe() == "censored=false" {
		t.Log("Big OptionsRanobe passed")
	} else {
		t.Error("Big OptionsRanobe failed")
	}

	zero := Options{
		Page: 0, Limit: 0, Score: 0, Censored: false, Genre_v2: []int{0},
	}
	if zero.OptionsRanobe() == "censored=false" {
		t.Log("Zero OptionsRanobe passed")
	} else {
		t.Error("Zero OptionsRanobe failed")
	}

	negative := Options{
		Page: -1, Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
	}
	if negative.OptionsRanobe() == "censored=false" {
		t.Log("Negative OptionsRanobe passed")
	} else {
		t.Error("Negative OptionsRanobe failed")
	}

	normal := Options{
		Page: 4, Limit: 5, Order: constants.MANGA_ORDER_ID, Status: constants.MANGA_STATUS_ANONS,
		Season: constants.SEASON_198x, Score: 7, Censored: true, Mylist: constants.MY_LIST_PLANNED,
		Genre_v2: []int{540, 540},
	}
	if normal.OptionsRanobe() == "censored=true&genre_v2=540-Erotica&limit=5&mylist=planned&order=id&page=4&score=7&season=198x&status=anons" {
		t.Log("Normal OptionsRanobe passed")
	} else {
		t.Error("Normal OptionsRanobe failed")
	}
}

func TestOptionsRanobeV2(t *testing.T) {
	normal := Options{
		Page: 4, Limit: 5, Order: constants.MANGA_ORDER_ID, Status: constants.MANGA_STATUS_ANONS,
		Season: constants.SEASON_198x, Score: 7, Censored: true, Mylist: constants.MY_LIST_PLANNED,
		Genre_v2: []int{540, 540},
	}
	if normal.OptionsRanobeV2() == "page=4&limit=5&score=7&order=id&kind=&status=anons&season=198x&mylist=planned&censored=true" {
		t.Log("Normal OptionsRanobeV2 passed")
	} else {
		t.Error("Normal OptionsRanobeV2 failed")
	}
}

func TestOptionsCalendar(t *testing.T) {
	empty := Options{}
	if empty.OptionsCalendar() == "censored=false" {
		t.Log("Empty OptionsCalendar passed")
	} else {
		t.Error("Empty OptionsCalendar failed")
	}

	normal := Options{Censored: true}
	if normal.OptionsCalendar() == "censored=true" {
		t.Log("Normal OptionsCalendar passed")
	} else {
		t.Error("Normal OptionsCalendar failed")
	}
}

func TestOptionsCalendarV2(t *testing.T) {
	normal := Options{Censored: true}
	if normal.OptionsCalendarV2() == "censored=true" {
		t.Log("Normal OptionsCalendarV2 passed")
	} else {
		t.Error("Normal OptionsCalendarV2 failed")
	}
}

func TestOptionsAnimeRates(t *testing.T) {
	empty := Options{}
	if empty.OptionsAnimeRates() == "censored=false" {
		t.Log("Empty OptionsAnimeRates passed")
	} else {
		t.Error("Empty OptionsAnimeRates failed")
	}

	big := Options{Page: 100002, Limit: 5002}
	if big.OptionsAnimeRates() == "censored=false" {
		t.Log("Big OptionsAnimeRates passed")
	} else {
		t.Error("Big OptionsAnimeRates failed")
	}

	zero := Options{Page: 0, Limit: 0}
	if zero.OptionsAnimeRates() == "censored=false" {
		t.Log("Zero OptionsAnimeRates passed")
	} else {
		t.Error("Zero OptionsAnimeRates failed")
	}

	negative := Options{Page: -1, Limit: -1}
	if negative.OptionsAnimeRates() == "censored=false" {
		t.Log("Negative OptionsAnimeRates passed")
	} else {
		t.Error("Negative OptionsAnimeRates failed")
	}

	normal := Options{Page: 15, Limit: 405, Status: constants.MY_LIST_DROPPED, Censored: true}
	if normal.OptionsAnimeRates() == "censored=true&limit=405&page=15&status=dropped" {
		t.Log("Normal OptionsAnimeRates passed")
	} else {
		t.Error("Normal OptionsAnimeRates failed")
	}
}

func TestOptionsAnimeRatesV2(t *testing.T) {
	normal := Options{Page: 15, Limit: 405, Status: constants.MY_LIST_DROPPED, Censored: true}
	if normal.OptionsAnimeRatesV2() == "page=15&limit=405&status=dropped&censored=true" {
		t.Log("Normal OptionsAnimeRatesV2 passed")
	} else {
		t.Error("Normal OptionsAnimeRatesV2 failed")
	}
}

func TestOptionsMangaRates(t *testing.T) {
	empty := Options{}
	if empty.OptionsMangaRates() == "censored=false" {
		t.Log("Empty OptionsMangaRates passed")
	} else {
		t.Error("Empty OptionsMangaRates failed")
	}

	big := Options{Page: 100002, Limit: 5002}
	if big.OptionsMangaRates() == "censored=false" {
		t.Log("Big OptionsMangaRates passed")
	} else {
		t.Error("Big OptionsMangaRates failed")
	}

	zero := Options{Page: 0, Limit: 0}
	if zero.OptionsMangaRates() == "censored=false" {
		t.Log("Zero OptionsMangaRates passed")
	} else {
		t.Error("Zero OptionsMangaRates failed")
	}

	negative := Options{Page: -1, Limit: -1}
	if negative.OptionsMangaRates() == "censored=false" {
		t.Log("Negative OptionsMangaRates passed")
	} else {
		t.Error("Negative OptionsMangaRates failed")
	}

	normal := Options{Page: 33, Limit: 25, Censored: true}
	if normal.OptionsMangaRates() == "censored=true&limit=25&page=33" {
		t.Log("Normal OptionsMangaRates passed")
	} else {
		t.Error("Normal OptionsMangaRates failed")
	}
}

func TestOptionsMangaRatesV2(t *testing.T) {
	normal := Options{Page: 33, Limit: 25, Censored: true}
	if normal.OptionsMangaRatesV2() == "page=33&limit=25&censored=true" {
		t.Log("Normal OptionsMangaRatesV2 passed")
	} else {
		t.Error("Normal OptionsMangaRatesV2 failed")
	}
}

func TestOptionsPeople(t *testing.T) {
	empty := Options{}
	if empty.OptionsPeople() == "" {
		t.Log("Empty OptionsPeople passed")
	} else {
		t.Error("Empty OptionsPeople failed")
	}

	normal := Options{Kind: constants.PEOPLE_KIND_MANGAKA}
	if normal.OptionsPeople() == "kind=mangaka" {
		t.Log("Normal OptionsPeople passed")
	} else {
		t.Error("Normal OptionsPeople failed")
	}
}

func TestOptionsPeopleV2(t *testing.T) {
	normal := Options{Kind: constants.PEOPLE_KIND_MANGAKA}
	if normal.OptionsPeopleV2() == "kind=mangaka" {
		t.Log("Normal OptionsPeopleV2 passed")
	} else {
		t.Error("Normal OptionsPeopleV2 failed")
	}
}

func TestOptionsRandomAnime(t *testing.T) {
	empty := Options{}
	if empty.OptionsRandomAnime() == "censored=false" {
		t.Log("Empty OptionsRandomAnime passed")
	} else {
		t.Error("Empty OptionsRandomAnime failed")
	}

	big := Options{
		Limit: 52, Score: 111111111111,
		Censored: false, Genre_v2: []int{111111111},
	}
	if big.OptionsRandomAnime() == "censored=false" {
		t.Log("Big OptionsRandomAnime passed")
	} else {
		t.Error("Big OptionsRandomAnime failed")
	}

	zero := Options{
		Limit: 0, Score: 0, Censored: false, Genre_v2: []int{0},
	}
	if zero.OptionsRandomAnime() == "censored=false" {
		t.Log("Zero OptionsRandomAnime passed")
	} else {
		t.Error("Zero OptionsRandomAnime failed")
	}

	negative := Options{
		Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
	}
	if negative.OptionsRandomAnime() == "censored=false" {
		t.Log("Negative OptionsRandomAnime passed")
	} else {
		t.Error("Negative OptionsRandomAnime failed")
	}

	normal := Options{
		Limit: 12, Kind: constants.ANIME_KIND_TV, Status: constants.ANIME_STATUS_RELEASED,
		Mylist: constants.MY_LIST_ON_HOLD, Season: constants.SEASON_199x,
		Score: 8, Rating: constants.ANIME_RATING_R, Duration: constants.ANIME_DURATION_D,
		Censored: true, Genre_v2: []int{539, 539},
	}
	if normal.OptionsRandomAnime() == "censored=true&duration=D&genre_v2=539-Erotica&kind=tv&limit=12&mylist=on_hold&rating=r&score=8&season=199x&status=released" {
		t.Log("Normal OptionsRandomAnime passed")
	} else {
		t.Error("Normal OptionsRandomAnime failed")
	}
}

func TestOptionsRandomAnimeV2(t *testing.T) {
	normal := Options{
		Limit: 12, Kind: constants.ANIME_KIND_TV, Status: constants.ANIME_STATUS_RELEASED,
		Mylist: constants.MY_LIST_ON_HOLD, Season: constants.SEASON_199x,
		Score: 8, Rating: constants.ANIME_RATING_R, Duration: constants.ANIME_DURATION_D,
		Censored: true, Genre_v2: []int{539, 539},
	}
	if normal.OptionsAnimeV2() == "page=0&limit=12&score=8&order=&kind=tv&status=released&season=199x&rating=r&duration=D&mylist=on_hold&censored=true&genre_v2=539-Erotica" {
		t.Log("Normal OptionsRandomAnimeV2 passed")
	} else {
		t.Error("Normal OptionsRandomAnimeV2 failed")
	}
}

func TestOptionsRandomManga(t *testing.T) {
	empty := Options{}
	if empty.OptionsRandomManga() == "censored=false" {
		t.Log("Empty OptionsRandomManga passed")
	} else {
		t.Error("Empty OptionsRandomManga failed")
	}

	big := Options{
		Limit: 52, Score: 1111111111,
		Censored: false, Genre_v2: []int{111111111111},
	}
	if big.OptionsRandomManga() == "censored=false" {
		t.Log("Big OptionsRandomManga passed")
	} else {
		t.Error("Big OptionsRandomManga failed")
	}

	zero := Options{
		Limit: 0, Score: 0, Censored: false, Genre_v2: []int{0},
	}
	if zero.OptionsRandomManga() == "censored=false" {
		t.Log("Zero OptionsRandomManga passed")
	} else {
		t.Error("Zero OptionsRandomManga failed")
	}

	negative := Options{
		Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
	}
	if negative.OptionsRandomManga() == "censored=false" {
		t.Log("Negative OptionsRandomManga passed")
	} else {
		t.Error("Negative OptionsRandomManga failed")
	}

	normal := Options{
		Limit: 5, Kind: constants.MANGA_KIND_MANGA, Status: constants.MANGA_STATUS_ANONS,
		Season: constants.SEASON_198x, Score: 7, Censored: true,
		Mylist: constants.MY_LIST_PLANNED, Genre_v2: []int{540, 540},
	}
	if normal.OptionsRandomManga() == "censored=true&genre_v2=540-Erotica&kind=manga&limit=5&mylist=planned&score=7&season=198x&status=anons" {
		t.Log("Normal OptionsRandomManga passed")
	} else {
		t.Error("Normal OptionsRandomManga failed")
	}
}

func TestOptionsRandomMangaV2(t *testing.T) {
	normal := Options{
		Limit: 5, Kind: constants.MANGA_KIND_MANGA, Status: constants.MANGA_STATUS_ANONS,
		Season: constants.SEASON_198x, Score: 7, Censored: true,
		Mylist: constants.MY_LIST_PLANNED, Genre_v2: []int{540, 540},
	}
	if normal.OptionsMangaV2() == "page=0&limit=5&score=7&order=&kind=manga&status=anons&season=198x&mylist=planned&censored=true" {
		t.Log("Normal OptionsRandomMangaV2 passed")
	} else {
		t.Error("Normal OptionsRandomMangaV2 failed")
	}
}

func TestOptionsRandomRanobe(t *testing.T) {
	empty := Options{}
	if empty.OptionsRandomRanobe() == "censored=false" {
		t.Log("Empty OptionsRandomRanobe passed")
	} else {
		t.Error("Empty OptionsRandomRanobe failed")
	}

	big := Options{
		Limit: 52, Score: 1111111111,
		Censored: false, Genre_v2: []int{111111111111},
	}
	if big.OptionsRandomRanobe() == "censored=false" {
		t.Log("Big OptionsRandomRanobe passed")
	} else {
		t.Error("Big OptionsRandomRanobe failed")
	}

	zero := Options{
		Limit: 0, Score: 0, Censored: false, Genre_v2: []int{0},
	}
	if zero.OptionsRandomRanobe() == "censored=false" {
		t.Log("Zero OptionsRandomRanobe passed")
	} else {
		t.Error("Zero OptionsRandomRanobe failed")
	}

	negative := Options{
		Limit: -1, Score: -1, Censored: false, Genre_v2: []int{-1},
	}
	if negative.OptionsRandomRanobe() == "censored=false" {
		t.Log("Negative OptionsRandomRanobe passed")
	} else {
		t.Error("Negative OptionsRandomRanobe failed")
	}

	normal := Options{
		Limit: 5, Status: constants.MANGA_STATUS_ANONS,
		Season: constants.SEASON_198x, Score: 7, Censored: true,
		Mylist: constants.MY_LIST_PLANNED, Genre_v2: []int{540, 540},
	}
	if normal.OptionsRandomRanobe() == "censored=true&genre_v2=540-Erotica&limit=5&mylist=planned&score=7&season=198x&status=anons" {
		t.Log("Normal OptionsRandomRanobe passed")
	} else {
		t.Error("Normal OptionsRandomRanobe failed")
	}
}

func TestOptionsRandomRanobeV2(t *testing.T) {
	normal := Options{
		Limit: 5, Status: constants.MANGA_STATUS_ANONS,
		Season: constants.SEASON_198x, Score: 7, Censored: true,
		Mylist: constants.MY_LIST_PLANNED, Genre_v2: []int{540, 540},
	}
	if normal.OptionsRanobeV2() == "page=0&limit=5&score=7&order=&kind=&status=anons&season=198x&mylist=planned&censored=true" {
		t.Log("Normal OptionsRandomRanobeV2 passed")
	} else {
		t.Error("Normal OptionsRandomRanobeV2 failed")
	}
}
