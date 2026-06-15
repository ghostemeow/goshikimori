package models

type GraphQL struct {
	Data   Data             `json:"data"`
	Errors []MessageGraphQL `json:"errors"`
}

type Data struct {
	Animes     []AnimesGraphQL     `json:"animes"`
	Mangas     []MangasGraphQL     `json:"mangas"`
	Characters []CharactersGraphQL `json:"characters"`
	People     []PeopleGraphQL     `json:"people"`
	UserRates  []UserRatesGraphQL  `json:"userrates"`
}

type MessageGraphQL struct {
	Message string `json:"message"`
}

/*
# Removed:

	ExternalLinks: nil;
*/
type AnimesGraphQL struct {
	Id                string                  `json:"id"`
	MalId             string                  `json:"malId"`
	Name              string                  `json:"name"`
	Russian           string                  `json:"russian"`
	LicenseNameRu     string                  `json:"licenseNameRu"`
	English           string                  `json:"english"`
	Japanese          string                  `json:"japanese"`
	Synonyms          []string                `json:"synonyms"`
	Kind              string                  `json:"kind"`
	Rating            string                  `json:"rating"`
	Status            string                  `json:"status"`
	Episodes          int                     `json:"episodes"`
	EpisodesAired     int                     `json:"episodesAired"`
	Duration          int                     `json:"duration"`
	AiredOn           AiredOnGraphQL          `json:"airedOn"`
	ReleasedOn        AiredOnGraphQL          `json:"releasedOn"`
	Url               string                  `json:"url"`
	Season            string                  `json:"season"`
	Poster            PosterGraphQL           `json:"poster"`
	Fansubbers        []string                `json:"fansubbers"`
	Fandubbers        []string                `json:"fandubbers"`
	Licensors         []string                `json:"licensors"`
	CreatedAt         string                  `json:"createdAt"`
	UpdatedAt         string                  `json:"updatedAt"`
	NextEpisodeAt     string                  `json:"nextEpisodeAt"`
	Genres            []GenresGraphQL         `json:"genres"`
	Studios           []StudiosGraphQL        `json:"studios"`
	PersonRoles       []PersonRolesGraphQL    `json:"personRoles"`
	CharacterRoles    []CharacterRolesGraphQL `json:"characterRoles"`
	Related           []RelatedGraphQL        `json:"related"`
	Videos            []VideosGraphQL         `json:"videos"`
	Screenshots       []ScreenshotsGraphQL    `json:"screenshots"`
	ScoresStats       []ScoresStatsGraphQL    `json:"scoresStats"`
	StatusesStats     []StatusesStatsGraphQL  `json:"statusesStats"`
	Description       string                  `json:"description"`
	DescriptionHTML   string                  `json:"descriptionHtml"`
	DescriptionSource string                  `json:"descriptionSource"`
	Score             float32                 `json:"score"`
	IsCensored        bool                    `json:"isCensored"`
}

/*
# Removed:

	ExternalLinks: nil;
*/
type MangasGraphQL struct {
	Id                string                  `json:"id"`
	MalId             string                  `json:"malId"`
	Name              string                  `json:"name"`
	Russian           string                  `json:"russian"`
	LicenseNameRu     string                  `json:"licenseNameRu"`
	English           string                  `json:"english"`
	Japanese          string                  `json:"japanese"`
	Synonyms          []string                `json:"synonyms"`
	Kind              string                  `json:"kind"`
	Status            string                  `json:"status"`
	Volumes           int                     `json:"volumes"`
	Chapters          int                     `json:"chapters"`
	AiredOn           AiredOnGraphQL          `json:"airedOn"`
	ReleasedOn        AiredOnGraphQL          `json:"releasedOn"`
	Url               string                  `json:"url"`
	Poster            PosterGraphQL           `json:"poster"`
	Licensors         []string                `json:"licensors"`
	CreatedAt         string                  `json:"createdAt"`
	UpdatedAt         string                  `json:"updatedAt"`
	Genres            []GenresGraphQL         `json:"genres"`
	Publishers        []ObjectInfo            `json:"publishers"`
	PersonRoles       []PersonRolesGraphQL    `json:"personRoles"`
	CharacterRoles    []CharacterRolesGraphQL `json:"characterRoles"`
	Related           []RelatedGraphQL        `json:"related"`
	ScoresStats       []ScoresStatsGraphQL    `json:"scoresStats"`
	StatusesStats     []StatusesStatsGraphQL  `json:"statusesStats"`
	Description       string                  `json:"description"`
	DescriptionHTML   string                  `json:"descriptionHtml"`
	DescriptionSource string                  `json:"descriptionSource"`
	Score             float32                 `json:"score"`
	IsCensored        bool                    `json:"isCensored"`
}

type UserRatesGraphQL struct {
	Anime     AnimesGraphQL `json:"anime"`
	Manga     MangasGraphQL `json:"manga"`
	Id        string        `json:"id"`
	Text      string        `json:"text"`
	CreatedAt string        `json:"createdAt"`
	UpdatedAt string        `json:"updatedAt"`
	Rewatches int           `json:"rewatches"`
	Score     int           `json:"score"`
	Status    string        `json:"status"`
	Episodes  int           `json:"episodes"`
	Chapters  int           `json:"chapters"`
	Volumes   int           `json:"volumes"`
}

type AiredOnGraphQL struct {
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
	Date  string `json:"date"`
}

type PosterGraphQL struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	MainUrl     string `json:"mainUrl"`
}

type GenresGraphQL struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Kind    string `json:"kind"`
}

type StudiosGraphQL struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

type PersonRolesGraphQL struct {
	Id      string   `json:"id"`
	RolesRu []string `json:"rolesRu"`
	RolesEn []string `json:"rolesEn"`
	Person  Persons  `json:"person"`
}

type Persons struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Poster Posters `json:"poster"`
}

type Posters struct {
	Id string `json:"id"`
}

type CharacterRolesGraphQL struct {
	Id        string   `json:"id"`
	RolesRu   []string `json:"rolesRu"`
	RolesEn   []string `json:"rolesEn"`
	Character Persons  `json:"character"`
}

// If the Anime or Manga is not found, nil is returned.
type RelatedGraphQL struct {
	Id           string     `json:"id"`
	Anime        ObjectInfo `json:"anime"`
	Manga        ObjectInfo `json:"manga"`
	RelationKind string     `json:"relationKind"`
	RelationText string     `json:"relationText"`
}

type ObjectInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type VideosGraphQL struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type ScreenshotsGraphQL struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	X166Url     string `json:"x166Url"`
	X332Url     string `json:"x332Url"`
}

type ScoresStatsGraphQL struct {
	Score int `json:"score"`
	Count int `json:"count"`
}

type StatusesStatsGraphQL struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type CharactersGraphQL struct {
	Id                string        `json:"id"`
	MalId             string        `json:"malId"`
	Name              string        `json:"name"`
	Russian           string        `json:"russian"`
	Japanese          string        `json:"japanese"`
	Synonyms          []string      `json:"synonyms"`
	Url               string        `json:"url"`
	CreatedAt         string        `json:"createdAt"`
	UpdatedAt         string        `json:"updatedAt"`
	IsAnime           bool          `json:"isAnime"`
	IsManga           bool          `json:"isManga"`
	IsRanobe          bool          `json:"isRanobe"`
	Poster            PosterGraphQL `json:"poster"`
	Description       string        `json:"description"`
	DescriptionHTML   string        `json:"descriptionHtml"`
	DescriptionSource string        `json:"descriptionSource"`
}

type PeopleGraphQL struct {
	Id         string         `json:"id"`
	MalId      string         `json:"malId"`
	Name       string         `json:"name"`
	Russian    string         `json:"russian"`
	Japanese   string         `json:"japanese"`
	Synonyms   []string       `json:"synonyms"`
	Url        string         `json:"url"`
	IsSeyu     bool           `json:"isSeyu"`
	IsMangaka  bool           `json:"isMangaka"`
	IsProducer bool           `json:"isProducer"`
	Website    string         `json:"website"`
	BirthOn    AiredOnGraphQL `json:"birthOn"`
	DeceasedOn AiredOnGraphQL `json:"deceasedOn"`
	Poster     PosterGraphQL  `json:"poster"`
}
