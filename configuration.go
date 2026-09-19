package goshikimori

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/ghostemeow/goshikimori/constants"
	"github.com/ghostemeow/goshikimori/genres"
)

type Configuration struct {
	Application string
	AccessToken string
}

type FastId struct {
	Id   int
	Conf Configuration
	Err  error
}

// Getting an id(anime, manga, ranobe, user, person, group).
//
// More information can be found in the [example1] and [example2].
//
// [example1]: https://github.com/ghostemeow/goshikimori/blob/master/examples/custom_fastid
// [example2]: https://github.com/ghostemeow/goshikimori/blob/master/examples/getter_setter
func (f *FastId) GetFastId() int { return f.Id }

// To create a custom id(anime, manga, ranobe, user, person, group).
//
// More information can be found in the [example1] and [example2].
//
// [example1]: https://github.com/ghostemeow/goshikimori/blob/master/examples/custom_fastid
// [example2]: https://github.com/ghostemeow/goshikimori/blob/master/examples/getter_setter
func (c *Configuration) SetFastId(id int) *FastId {
	return &FastId{Id: id, Conf: *c, Err: nil}
}

// Getting the configuration.
//
// More information can be found in the [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/getter_setter
func (c *Configuration) GetConfiguration() (string, string) {
	return c.Application, c.AccessToken
}

// To register the application, follow the link from [OAuth].
//
// More information can be found in the [example].
//
// [OAuth]: https://github.com/ghostemeow/goshikimori/blob/master/examples/first_steps
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/getter_setter
func SetConfiguration(appname, token string) *Configuration {
	return &Configuration{Application: appname, AccessToken: token}
}

type Options struct {
	Order            string
	Kind             string
	Status           string
	Season           string
	Rating           string
	Type             string
	Target_type      string
	Duration         string
	Mylist           string
	Forum            string
	Linked_type      string
	Commentable_type string
	Page             int
	Limit            int
	Score            int
	Linked_id        int
	Target_id        int
	Commentable_id   int
	Genre_v2         []int
	Censored         bool
	Desc             bool
}

type Result interface {
	OptionsOnlyPageLimit(int, int) string
	OptionsAnime() string
	OptionsManga() string
	OptionsRanobe() string
	OptionsCalendar() string
	OptionsAnimeRates() string
	OptionsMangaRates() string
	OptionsUserHistory() string
	OptionsMessages() string
	OptionsPeople() string
	OptionsTopics() string
	OptionsTopicsHot() string
	OptionsRandomAnime() string
	OptionsRandomManga() string
	OptionsRandomRanobe() string

	OptionsOnlyPageLimitV2() string
	OptionsAnimeV2() string
	OptionsMangaV2() string
	OptionsRanobeV2() string
	OptionsCalendarV2() string
	OptionsAnimeRatesV2() string
	OptionsMangaRatesV2() string
	OptionsUserHistoryV2() string
	OptionsMessagesV2() string
	OptionsPeopleV2() string
	OptionsTopicsV2() string
	OptionsTopicsHotV2() string
	OptionsCommentsV2() string
}

// Topic parameters for creating/updating a topic.
//
// Required for [Configuration.CreateTopic]:
//   - Body: topic text;
//   - Forum_id: forum id, can be found in SearchForums();
//   - Title: topic title;
//   - User_id: your id, can be found in WhoAmi() or FastIdUser();
//
// Optional:
//
//   - Linked_id and Linked_type are only used together:
//
//     > TOPIC_LINKED_TYPE_ANIME, TOPIC_LINKED_TYPE_MANGA, TOPIC_LINKED_TYPE_RANOBE,
//     TOPIC_LINKED_TYPE_CHARACTER, TOPIC_LINKED_TYPE_PERSON, TOPIC_LINKED_TYPE_CLUB,
//     TOPIC_LINKED_TYPE_CLUBPAGE, TOPIC_LINKED_TYPE_CRITIQUE, TOPIC_LINKED_TYPE_REVIEW,
//     TOPIC_LINKED_TYPE_CONTEST, TOPIC_LINKED_TYPE_COSPLAYGALLYRY,
//     TOPIC_LINKED_TYPE_COLLECTION, TOPIC_LINKED_TYPE_ARTICLE;
//
//   - Type: defaults to TOPIC_TYPE for [Configuration.CreateTopic],
//     must not be set for [Configuration.UpdateTopic].
type TopicParams struct {
	Body        string `json:"body,omitempty"`
	Forum_id    int    `json:"forum_id,omitempty"`
	Linked_id   int    `json:"linked_id,omitempty"`
	Linked_type string `json:"linked_type,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
	User_id     int    `json:"user_id,omitempty"`
}

// Comment parameters for creating/updating a comment.
//
// Required for [Configuration.CreateComment]:
//
//   - Body: comment text;
//
//   - Commentable_id: id of the commentable object;
//
//   - Commentable_type: one of the COMMENTABLE_TYPE_* constants,
//     for the comment form the API expects one of:
//
//     > COMMENTABLE_TYPE_TOPIC, COMMENTABLE_TYPE_USER, COMMENTABLE_TYPE_ANIME, COMMENTABLE_TYPE_MANGA,
//     COMMENTABLE_TYPE_CHARACTER, COMMENTABLE_TYPE_PERSON, COMMENTABLE_TYPE_ARTICLE,
//     COMMENTABLE_TYPE_CLUB, COMMENTABLE_TYPE_CLUBPAGE, COMMENTABLE_TYPE_COLLECTION,
//     COMMENTABLE_TYPE_CRITIQUE, COMMENTABLE_TYPE_REVIEW;
//
//     When set to Anime, Manga, Character, Person, Article, Club, ClubPage,
//     Collection, Critique, Review, the comment is attached to the commentable main topic.
//
// Optional:
//
//   - Is_offtopic: mark the comment as offtopic.
type CommentParams struct {
	Body             string `json:"body,omitempty"`
	Commentable_id   int    `json:"commentable_id,omitempty"`
	Commentable_type string `json:"commentable_type,omitempty"`
	Is_offtopic      bool   `json:"is_offtopic,string,omitempty"`
}

// TODO: (ghostemeow) abandon url.QueryEscape in the future.
func encodeParamEscaped(key, value string) string {
	return url.QueryEscape(key) + "=" + url.QueryEscape(value)
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsOnlyPageLimit(page, limit int) string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= page {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= limit {
		v.Add("limit", strconv.Itoa(o.Limit))
	}

	return v.Encode()
}

func (o *Options) OptionsOnlyPageLimitV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 2)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsTopics() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 30 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Forum != "" {
		v.Add("forum", o.Forum)
	}
	// linked_id and linked_type are only used together.
	if o.Linked_id >= 1 && o.Linked_type != "" {
		v.Add("linked_id", strconv.Itoa(o.Linked_id))
		v.Add("linked_type", o.Linked_type)
	}

	return v.Encode()
}

func (o *Options) OptionsTopicsV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 5)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("forum", o.Forum))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Linked_id), 10)
	pairs = append(pairs, encodeParamEscaped("linked_id", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("linked_type", o.Linked_type))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsMessages() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 100 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	// The type is required.
	if o.Type == "" {
		v.Add("type", constants.MESSAGE_TYPE_NEWS)
	} else {
		v.Add("type", o.Type)
	}

	return v.Encode()
}

func (o *Options) OptionsMessagesV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 3)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("type", o.Type))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsUserHistory() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 100 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Target_type != "" {
		v.Add("target_type", o.Target_type)
	}
	if o.Target_id > 0 {
		v.Add("target_id", strconv.Itoa(o.Target_id))
	}

	return v.Encode()
}

func (o *Options) OptionsUserHistoryV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 4)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("target_type", o.Target_type))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Target_id), 10)
	pairs = append(pairs, encodeParamEscaped("target_id", string(numBuf)))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsAnime() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Order != "" {
		v.Add("order", o.Order)
	}
	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Rating != "" {
		v.Add("rating", o.Rating)
	}
	if o.Duration != "" {
		v.Add("duration", o.Duration)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := genres.MapGenresAnime(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

func (o *Options) OptionsAnimeV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 12)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Score), 10)
	pairs = append(pairs, encodeParamEscaped("score", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("order", o.Order))
	pairs = append(pairs, encodeParamEscaped("kind", o.Kind))
	pairs = append(pairs, encodeParamEscaped("status", o.Status))
	pairs = append(pairs, encodeParamEscaped("season", o.Season))
	pairs = append(pairs, encodeParamEscaped("rating", o.Rating))
	pairs = append(pairs, encodeParamEscaped("duration", o.Duration))
	pairs = append(pairs, encodeParamEscaped("mylist", o.Mylist))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}
	if genre := genres.MapGenresAnime(o.Genre_v2); genre != "" {
		pairs = append(pairs, encodeParamEscaped("genre_v2", genre))
	}

	totalLength := 0
	for _, p := range pairs {
		totalLength += len(p) + 1
	}
	if totalLength > 0 {
		totalLength--
	}
	sb.Grow(totalLength)

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsManga() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Order != "" {
		v.Add("order", o.Order)
	}
	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := genres.MapGenresManga(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

func (o *Options) OptionsMangaV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 10)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Score), 10)
	pairs = append(pairs, encodeParamEscaped("score", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("order", o.Order))
	pairs = append(pairs, encodeParamEscaped("kind", o.Kind))
	pairs = append(pairs, encodeParamEscaped("status", o.Status))
	pairs = append(pairs, encodeParamEscaped("season", o.Season))
	pairs = append(pairs, encodeParamEscaped("mylist", o.Mylist))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}
	if genre := genres.MapGenresAnime(o.Genre_v2); genre != "" {
		pairs = append(pairs, encodeParamEscaped("genre_v2", genre))
	}

	totalLength := 0
	for _, p := range pairs {
		totalLength += len(p) + 1
	}
	if totalLength > 0 {
		totalLength--
	}
	sb.Grow(totalLength)

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsRanobe() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Order != "" {
		v.Add("order", o.Order)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := genres.MapGenresManga(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

func (o *Options) OptionsRanobeV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 10)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Score), 10)
	pairs = append(pairs, encodeParamEscaped("score", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("order", o.Order))
	pairs = append(pairs, encodeParamEscaped("kind", o.Kind))
	pairs = append(pairs, encodeParamEscaped("status", o.Status))
	pairs = append(pairs, encodeParamEscaped("season", o.Season))
	pairs = append(pairs, encodeParamEscaped("mylist", o.Mylist))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}
	if genre := genres.MapGenresAnime(o.Genre_v2); genre != "" {
		pairs = append(pairs, encodeParamEscaped("genre_v2", genre))
	}

	totalLength := 0
	for _, p := range pairs {
		totalLength += len(p) + 1
	}
	if totalLength > 0 {
		totalLength--
	}
	sb.Grow(totalLength)

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsCalendar() string {
	v := url.Values{}

	v.Add("censored", strconv.FormatBool(o.Censored))

	return v.Encode()
}

func (o *Options) OptionsCalendarV2() string {
	var sb strings.Builder
	pairs := make([]string, 0, 1)

	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsAnimeRates() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 5000 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	return v.Encode()
}

func (o *Options) OptionsAnimeRatesV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 4)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("status", o.Status))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsMangaRates() string {
	v := url.Values{}

	if o.Page >= 1 && o.Page <= 100000 {
		v.Add("page", strconv.Itoa(o.Page))
	}
	if o.Limit >= 1 && o.Limit <= 5000 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	return v.Encode()
}

// FIXME: (ghostemeow) The manga has no status, ranobe is missing.
// https://shikimori.one/api/doc/1.0/users/manga_rates.html
func (o *Options) OptionsMangaRatesV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 3)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	if o.Censored {
		pairs = append(pairs, encodeParamEscaped("censored", "true"))
	} else {
		pairs = append(pairs, encodeParamEscaped("censored", "false"))
	}

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsPeople() string {
	v := url.Values{}

	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}

	return v.Encode()
}

func (o *Options) OptionsPeopleV2() string {
	var sb strings.Builder
	pairs := make([]string, 0, 1)

	pairs = append(pairs, encodeParamEscaped("kind", o.Kind))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsTopicsHot() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 10 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}

	return v.Encode()
}

func (o *Options) OptionsTopicsHotV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 1)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

func (o *Options) OptionsCommentsV2() string {
	var numBuf []byte
	var sb strings.Builder
	pairs := make([]string, 0, 5)

	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Commentable_id), 10)
	pairs = append(pairs, encodeParamEscaped("commentable_id", string(numBuf)))
	pairs = append(pairs, encodeParamEscaped("commentable_type", o.Commentable_type))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Page), 10)
	pairs = append(pairs, encodeParamEscaped("page", string(numBuf)))
	numBuf = strconv.AppendInt(numBuf[:0], int64(o.Limit), 10)
	pairs = append(pairs, encodeParamEscaped("limit", string(numBuf)))
	if o.Desc {
		pairs = append(pairs, encodeParamEscaped("desc", "1"))
	} else {
		pairs = append(pairs, encodeParamEscaped("desc", "0"))
	}

	for i, p := range pairs {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(p)
	}

	return sb.String()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsRandomAnime() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Rating != "" {
		v.Add("rating", o.Rating)
	}
	if o.Duration != "" {
		v.Add("duration", o.Duration)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := genres.MapGenresAnime(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsRandomManga() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Kind != "" {
		v.Add("kind", o.Kind)
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := genres.MapGenresManga(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}

// NOTE: (ghostemeow) DEPRECATED and will be removed from future versions.
func (o *Options) OptionsRandomRanobe() string {
	v := url.Values{}

	if o.Limit >= 1 && o.Limit <= 50 {
		v.Add("limit", strconv.Itoa(o.Limit))
	}
	if o.Score >= 1 && o.Score <= 9 {
		v.Add("score", strconv.Itoa(o.Score))
	}
	if o.Status != "" {
		v.Add("status", o.Status)
	}
	if o.Season != "" {
		v.Add("season", o.Season)
	}
	if o.Mylist != "" {
		v.Add("mylist", o.Mylist)
	}
	v.Add("censored", strconv.FormatBool(o.Censored))

	genre := genres.MapGenresManga(o.Genre_v2)
	if genre != "" {
		v.Add("genre_v2", genre)
	}

	return v.Encode()
}
