package models

import "time"

type UserSize struct {
	X160 string `json:"x160"`
	X148 string `json:"x148"`
	X80  string `json:"x80"`
	X64  string `json:"x64"`
	X48  string `json:"x48"`
	X32  string `json:"x32"`
	X16  string `json:"x16"`
}

type StatusesAnimeManga struct {
	Id         int    `json:"id"`
	Grouped_id string `json:"grouped_id"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Type       string `json:"type"`
}

/*
# Removed:

	Full_statuses: no response after request;
	Scores: empty slice;
	Types: empty slice;
	Ratings: empty slice;
	Has_anime? || Has_manga?: always one boolean state: false;
	Genres: empty slice;
	Publishers: empty slice;
	Activity: no response after request;
*/
type StatusesInfo struct {
	Anime []StatusesAnimeManga `json:"anime"`
	Manga []StatusesAnimeManga `json:"manga"`
}

type StatsInfo struct {
	Statuses StatusesInfo `json:"statuses"`
}

type Users struct {
	Id             int       `json:"id"`
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Image          UserSize  `json:"image"`
	Last_online_at time.Time `json:"last_online_at"`
	Name           string    `json:"name"`
	Sex            string    `json:"sex"`
	Full_years     int       `json:"full_years"`
	Last_online    string    `json:"last_online"`
	Website        string    `json:"website"`
	Location       string    `json:"location"`
	Banned         bool      `json:"banned"`
	About          string    `json:"about"`
	AboutHTML      string    `json:"about_html"`
	Common_info    []string  `json:"common_info"`
	Show_comments  bool      `json:"show_comments"`
	In_friends     bool      `json:"in_friends"`
	Is_ignored     bool      `json:"is_ignored"`
	Stats          StatsInfo `json:"stats"`
	Style_id       int       `json:"style_id"`
}

type UserFriends struct {
	Id             int       `json:"id"`
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Image          UserSize  `json:"image"`
	Last_online_at time.Time `json:"last_online_at"`
	Url            string    `json:"url"`
}

/*
# Removed:

	Chapters: nil;
	Volumes: nil;
	Manga: nil;
*/
type UserAnimeRates struct {
	Id         int         `json:"id"`
	Score      int         `json:"score"`
	Status     string      `json:"status"`
	Text       string      `json:"text"`
	Episodes   int         `json:"episodes"`
	Text_html  string      `json:"text_html"`
	Rewatches  int         `json:"rewatches"`
	Created_at time.Time   `json:"created_at"`
	Updated_at time.Time   `json:"updated_at"`
	User       UserFriends `json:"user"`
	Anime      Animes      `json:"anime"`
}

// # Removed:
//
//	Episodes: nil;
//	Anime: nil;
type UserMangaRates struct {
	Id         int         `json:"id"`
	Score      int         `json:"score"`
	Status     string      `json:"status"`
	Text       string      `json:"text"`
	Chapters   int         `json:"chapters"`
	Volumes    int         `json:"volumes"`
	Text_html  string      `json:"text_html"`
	Rewatches  int         `json:"rewatches"`
	Created_at time.Time   `json:"created_at"`
	Updated_at time.Time   `json:"updated_at"`
	User       UserFriends `json:"user"`
	Manga      Mangas      `json:"manga"`
}

// # Removed:
//
//	Url: nil;
type UserFavouritesInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Image   string `json:"image"`
}

type UserFavourites struct {
	Animes     []UserFavouritesInfo `json:"animes"`
	Mangas     []UserFavouritesInfo `json:"mangas"`
	Characters []UserFavouritesInfo `json:"characters"`
	People     []UserFavouritesInfo `json:"people"`
	Mangakas   []UserFavouritesInfo `json:"mangakas"`
	Seyu       []UserFavouritesInfo `json:"seyu"`
	Producers  []UserFavouritesInfo `json:"producers"`
}

// # Notes:
//
//	Episodes and Episodes_aired only for anime.
//	Volumes and Chapters only for manga.
type TargetInfo struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Russian        string    `json:"russian"`
	Image          AnimeSize `json:"image"`
	Url            string    `json:"url"`
	Kind           string    `json:"kind"`
	Score          string    `json:"score"`
	Status         string    `json:"status"`
	Episodes       int       `json:"episodes"`
	Episodes_aired int       `json:"episodes_aired"`
	Volumes        int       `json:"volumes"`
	Chapters       int       `json:"chapters"`
	Aired_on       string    `json:"aired_on"`
	Released_on    string    `json:"released_on"`
}

type UserHistory struct {
	Id          int        `json:"id"`
	Created_at  time.Time  `json:"created_at"`
	Description string     `json:"description"`
	Target      TargetInfo `json:"target"`
}

type Who struct {
	Id             int       `json:"id"`
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Image          UserSize  `json:"image"`
	Last_online_at time.Time `json:"last_online_at"`
	Name           string    `json:"name"`
	Sex            string    `json:"sex"`
	Website        string    `json:"website"`
	Birth_on       int       `json:"birth_on"`
	Locale         string    `json:"locale"`
}

type AnimeSize struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
	X96      string `json:"x96"`
	X48      string `json:"x48"`
}

type Stats[T int | string] struct {
	Name  T   `json:"name"`
	Value int `json:"value"`
}

type AnimeGenres struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Russian    string `json:"russian"`
	Kind       string `json:"kind"`
	Entry_type string `json:"entry_type"`
}

type AnimeStudios struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Filtered_name string `json:"filtered_name"`
	Real          bool   `json:"real"`
	Image         string `json:"image"`
}

type UserRate struct {
	Id         int       `json:"id"`
	Score      int       `json:"score"`
	Status     string    `json:"status"`
	Text       string    `json:"text"`
	Episodes   int       `json:"episodes"`
	Chapters   int       `json:"chapters"`
	Volumes    int       `json:"volumes"`
	Text_html  string    `json:"text_html"`
	Rewatches  int       `json:"rewatches"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// # Removed:
//
//	Description_source: nil;
type Anime struct {
	Id                 int                `json:"id"`
	Name               string             `json:"name"`
	Russian            string             `json:"russian"`
	Image              AnimeSize          `json:"image"`
	Url                string             `json:"url"`
	Kind               string             `json:"kind"`
	Score              string             `json:"score"`
	Status             string             `json:"status"`
	Episodes           int                `json:"episodes"`
	Episodes_aired     int                `json:"episodes_aired"`
	Aired_on           string             `json:"aired_on"`
	Released_on        string             `json:"released_on"`
	Rating             string             `json:"rating"`
	English            []string           `json:"english"`
	Japanese           []string           `json:"japanese"`
	Synonyms           []string           `json:"synonyms"`
	LicenseNameRu      string             `json:"license_name_ru"`
	Duration           int                `json:"duration"`
	Description        string             `json:"description"`
	DescriptionHTML    string             `json:"description_html"`
	Franchise          string             `json:"franchise"`
	Favoured           bool               `json:"favoured"`
	Anons              bool               `json:"anons"`
	Ongoing            bool               `json:"ongoing"`
	Thread_id          int                `json:"thread_id"`
	Topic_id           int                `json:"topic_id"`
	Myanimelist_id     int                `json:"myanimelist_id"`
	RatesScoresStats   []Stats[int]       `json:"rates_scores_stats"`
	RatesStatusesStats []Stats[string]    `json:"rates_statuses_stats"`
	Updated_at         time.Time          `json:"updated_at"`
	Next_episode_at    time.Time          `json:"next_episode_at"`
	Fansubbers         []string           `json:"fansubbers"`
	Fandubbers         []string           `json:"fandubbers"`
	Licensors          []string           `json:"licensors"`
	Genres             []AnimeGenres      `json:"genres"`
	Studios            []AnimeStudios     `json:"studios"`
	Videos             []AnimeVideos      `json:"videos"`
	Screenshots        []AnimeScreenshots `json:"screenshots"`
	User_rate          UserRate           `json:"user_rate"`
}

type Animes struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Russian        string    `json:"russian"`
	Image          AnimeSize `json:"image"`
	Url            string    `json:"url"`
	Kind           string    `json:"kind"`
	Score          string    `json:"score"`
	Status         string    `json:"status"`
	Episodes       int       `json:"episodes"`
	Episodes_aired int       `json:"episodes_aired"`
	Aired_on       string    `json:"aired_on"`
	Released_on    string    `json:"released_on"`
}

// # Removed:
//
//	Description_source: nil;
type Manga struct {
	Id                 int             `json:"id"`
	Name               string          `json:"name"`
	Russian            string          `json:"russian"`
	Image              AnimeSize       `json:"image"`
	Url                string          `json:"url"`
	Kind               string          `json:"kind"`
	Score              string          `json:"score"`
	Status             string          `json:"status"`
	Volumes            int             `json:"volumes"`
	Chapters           int             `json:"chapters"`
	Aired_on           string          `json:"aired_on"`
	Released_on        string          `json:"released_on"`
	English            []string        `json:"english"`
	Japanese           []string        `json:"japanese"`
	Synonyms           []string        `json:"synonyms"`
	LicenseNameRu      string          `json:"license_name_ru"`
	Description        string          `json:"description"`
	DescriptionHTML    string          `json:"description_html"`
	Franchise          string          `json:"franchise"`
	Favoured           bool            `json:"favoured"`
	Anons              bool            `json:"anons"`
	Ongoing            bool            `json:"ongoing"`
	Thread_id          int             `json:"thread_id"`
	Topic_id           int             `json:"topic_id"`
	Myanimelist_id     int             `json:"myanimelist_id"`
	RatesScoresStats   []Stats[int]    `json:"rates_scores_stats"`
	RatesStatusesStats []Stats[string] `json:"rates_statuses_stats"`
	Licensors          []string        `json:"licensors"`
	Genres             []AnimeGenres   `json:"genres"`
	Publishers         []Publishers    `json:"publishers"`
	User_rate          UserRate        `json:"user_rate"`
}

type Mangas struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Russian     string    `json:"russian"`
	Image       AnimeSize `json:"image"`
	Url         string    `json:"url"`
	Kind        string    `json:"kind"`
	Score       string    `json:"score"`
	Status      string    `json:"status"`
	Volumes     int       `json:"volumes"`
	Chapters    int       `json:"chapters"`
	Aired_on    string    `json:"aired_on"`
	Released_on string    `json:"released_on"`
}

type ClubSize struct {
	Original string `json:"original"`
	Main     string `json:"main"`
	X96      string `json:"x96"`
	X73      string `json:"x73"`
	X48      string `json:"x48"`
}

type Clubs struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Logo           ClubSize `json:"logo"`
	Is_censored    bool     `json:"is_censored"`
	Join_policy    string   `json:"join_policy"`
	Comment_policy string   `json:"comment_policy"`
}

type Achievements struct {
	Id         int       `json:"id"`
	Neko_id    string    `json:"neko_id"`
	Level      int       `json:"level"`
	Progress   int       `json:"progress"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// # Removed:
//
//	Manga: bool;
type RelatedAnimes struct {
	Relation         string `json:"relation"`
	Relation_Russian string `json:"relation_russian"`
	Anime            Animes `json:"anime"`
}

// # Removed:
//
//	Anime: bool;
type RelatedMangas struct {
	Relation         string `json:"relation"`
	Relation_Russian string `json:"relation_russian"`
	Manga            Mangas `json:"manga"`
}

type AnimeScreenshots struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
}

type AnimeVideos struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Image_url  string `json:"image_url"`
	Player_url string `json:"player_url"`
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Hosting    string `json:"hosting"`
}

type CharacterInfo struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Russian string    `json:"russian"`
	Image   AnimeSize `json:"image"`
	Url     string    `json:"url"`
}

// # Removed:
//
//	Person: nil;
type Roles struct {
	Roles         []string      `json:"roles"`
	Roles_Russian []string      `json:"roles_russian"`
	Character     CharacterInfo `json:"character"`
}

type CommentInfo struct {
	Id               int       `json:"id"`
	Commentable_id   int       `json:"commentable_id"`
	Commentable_type string    `json:"commentable_type"`
	Body             string    `json:"body"`
	User_id          int       `json:"user_id"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updated_at"`
	Is_summary       bool      `json:"is_summary"`
	Is_offtopic      bool      `json:"is_offtopic"`
}

type UserInfo struct {
	Id             int       `json:"id"`
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Image          UserSize  `json:"image"`
	Last_online_at time.Time `json:"last_online_at"`
}

type Bans struct {
	Id               int         `json:"id"`
	User_id          int         `json:"user_id"`
	Comment          CommentInfo `json:"comment"`
	Moderator_id     int         `json:"moderator_id"`
	Reason           string      `json:"reason"`
	Created_at       time.Time   `json:"created_at"`
	Duration_minutes int         `json:"duration_minutes"`
	User             UserInfo    `json:"user"`
	Moderator        UserInfo    `json:"moderator"`
}

type Calendar struct {
	Next_episode    int       `json:"next_episode"`
	Next_episode_at time.Time `json:"next_episode_at"`
	Duration        int       `json:"duration"`
	Anime           Animes    `json:"anime"`
}

type LinksInfo struct {
	Id        int    `json:"id"`
	Source_id int    `json:"source_id"`
	Target_id int    `json:"target_id"`
	Source    int    `json:"source"`
	Target    int    `json:"target"`
	Weight    int    `json:"weight"`
	Relation  string `json:"relation"`
}

type NodesInfo struct {
	Id        int    `json:"id"`
	Date      int    `json:"date"`
	Name      string `json:"name"`
	Image_url string `json:"image_url"`
	Year      int    `json:"year"`
	Kind      string `json:"kind"`
	Weight    int    `json:"weight"`
}

type Franchise struct {
	Links      []LinksInfo `json:"links"`
	Nodes      []NodesInfo `json:"nodes"`
	Current_id int         `json:"current_id"`
}

type ExternalLinks struct {
	Id          int       `json:"id"`
	Kind        string    `json:"kind"`
	Url         string    `json:"url"`
	Source      string    `json:"source"`
	Entry_id    int       `json:"entry_id"`
	Entry_type  string    `json:"entry_type"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Imported_at time.Time `json:"imported_at"`
}

// # Removed:
//
//	Image: nil;
type Studios struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Filtered_name string `json:"filtered_name"`
	Real          bool   `json:"real"`
}

type Publishers struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Forums struct {
	Id        int    `json:"id"`
	Position  int    `json:"position"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	Url       string `json:"url"`
}

type FriendRequest struct {
	Notice string `json:"notice"`
}

type UnreadMessages struct {
	Messages      int `json:"messages"`
	News          int `json:"news"`
	Notifications int `json:"notifications"`
}

// # Removed:
//
//	Released_on: nil;
//	Aired_io: nil;
type LinkedMessages struct {
	Id             int       `json:"id"`
	Topic_url      string    `json:"topic_url"`
	Thread_id      int       `json:"thread_id"`
	Topic_id       int       `json:"topic_id"`
	Type           string    `json:"type"`
	Name           string    `json:"name"`
	Russian        string    `json:"russian"`
	Image          AnimeSize `json:"image"`
	Url            string    `json:"url"`
	Kind           string    `json:"kind"`
	Score          string    `json:"score"`
	Status         string    `json:"status"`
	Episodes       int       `json:"episodes"`
	Episodes_aired int       `json:"episodes_aired"`
}

type FromToMessages struct {
	Id             int       `json:"id"`
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Image          UserSize  `json:"image"`
	Last_online_at time.Time `json:"last_online_at"`
	Url            string    `json:"url"`
}

type Messages struct {
	Id          int            `json:"id"`
	Kind        string         `json:"kind"`
	Read        bool           `json:"read"`
	Body        string         `json:"body"`
	HTMLBody    string         `json:"html_body"`
	Created_at  time.Time      `json:"created_at"`
	Linked_id   int            `json:"linked_id"`
	Linked_type string         `json:"linked_type"`
	Linked      LinkedMessages `json:"linked"`
	From        FromToMessages `json:"from"`
	To          FromToMessages `json:"to"`
}

type Constants struct {
	Kind   []string `json:"kind"`
	Status []string `json:"status"`
}

type ConstantsUserRate struct {
	Status []string `json:"status"`
}

type ConstantsClub struct {
	Join_policy         []string `json:"join_policy"`
	Comment_policy      []string `json:"comment_policy"`
	Image_upload_policy []string `json:"image_upload_policy"`
}

type ConstantsSmileys struct {
	Bbcode string `json:"bbcode"`
	Path   string `json:"path"`
}

type Birth struct {
	Day   int `json:"day"`
	Year  int `json:"year"`
	Month int `json:"month"`
}

type RolesPeople struct {
	Characters []AllPeople `json:"characters"`
	Animes     []Animes    `json:"animes"`
}

// # Removed:
//
//	Manga: nil;
type WorksPeople struct {
	Anime Animes `json:"anime"`
	Role  string `json:"role"`
}

type AllPeople struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Russian string    `json:"russian"`
	Image   AnimeSize `json:"image"`
	Url     string    `json:"url"`
}

/*
# Removed:

	Deceased_on: no response after request;
	Person_favoured: bool;
	Producer/Producer_favoured: bool;
	Mangaka/Mangaka_favoured: bool;
	Seyu/Seyu_favoured: bool;
	Birthday: same as the Birth_on;
*/
type People struct {
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	Russian        string        `json:"russian"`
	Image          AnimeSize     `json:"image"`
	Url            string        `json:"url"`
	Japanese       string        `json:"japanese"`
	Job_title      string        `json:"job_title"`
	Birth_on       Birth         `json:"birth_on"`
	Website        string        `json:"website"`
	Groupped_roles [][]any       `json:"groupped_roles"`
	Roles          []RolesPeople `json:"roles"`
	Works          []WorksPeople `json:"works"`
	Topic_id       int           `json:"topic_id"`
	Updated_at     time.Time     `json:"updated_at"`
	Thread_id      int           `json:"thread_id"`
}

type ClubImages struct {
	Id           int    `json:"id"`
	Original_url string `json:"original_url"`
	Main_url     string `json:"main_url"`
	Preview_url  string `json:"preview_url"`
	Can_destroy  bool   `json:"can_destroy"`
	User_id      int    `json:"user_id"`
}

/*
# Removed:

	Linked: nil;
	Last_comment_viewed: nil;
	Event: nil;
	Episode: nil;
*/
type ClubCollections struct {
	Id             int            `json:"id"`
	Topic_title    string         `json:"topic_title"`
	Body           string         `json:"body"`
	HTMLBody       string         `json:"html_body"`
	HTMLFooter     string         `json:"html_footer"`
	Created_at     time.Time      `json:"created_at"`
	Comments_count int            `json:"comments_count"`
	Forum          Forums         `json:"forum"`
	User           FromToMessages `json:"user"`
	Type           string         `json:"type"`
	Linked_id      int            `json:"linked_id"`
	Linked_type    string         `json:"linked_type"`
	Viewed         bool           `json:"viewed"`
}

type Favorites struct {
	Success bool   `json:"success"`
	Notice  string `json:"notice"`
}

// # Removed:
//
//	Linked_type: nil;
//	Linked: nil;
type Target_message struct {
	Id         int       `json:"id"`
	Kind       string    `json:"kind"`
	Read       bool      `json:"read"`
	Body       string    `json:"body"`
	HTMLBody   string    `json:"html_body"`
	Created_at time.Time `json:"created_at"`
	Linked_id  int       `json:"linked_id"`
}

type Dialogs struct {
	Target_user UserFriends    `json:"target_user"`
	Message     Target_message `json:"message"`
}

// # Removed:
//
//	Linked_type: nil;
//	Linked: nil;
type SearchDialogs struct {
	Id         int         `json:"id"`
	Kind       string      `json:"kind"`
	Read       bool        `json:"read"`
	Body       string      `json:"body"`
	HTMLBody   string      `json:"html_body"`
	Created_at time.Time   `json:"created_at"`
	Linked_id  int         `json:"linked_id"`
	From       UserFriends `json:"from"`
	To         UserFriends `json:"to"`
}

type Info struct {
	Id             int       `json:"id"`
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Image          UserSize  `json:"image"`
	Last_online_at time.Time `json:"last_online_at"`
	Name           string    `json:"name"`
	Sex            string    `json:"sex"`
	Website        string    `json:"website"`
	Birth_on       int       `json:"birth_on"`
	Full_years     int       `json:"full_years"`
	Locale         string    `json:"locale"`
}

// Topics solution. In this struct there are both anime and manga fields.
type Animanga struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Russian        string    `json:"russian"`
	Image          AnimeSize `json:"image"`
	Url            string    `json:"url"`
	Kind           string    `json:"kind"`
	Score          string    `json:"score"`
	Status         string    `json:"status"`
	Episodes       int       `json:"episodes"`
	Episodes_aired int       `json:"episodes_aired"`
	Volumes        int       `json:"volumes"`
	Chapters       int       `json:"chapters"`
	Aired_on       string    `json:"aired_on"`
	Released_on    string    `json:"released_on"`
}

// # Removed:
//
//	Event: nil;
//	Episode: nil;
type Topics struct {
	Id                  int         `json:"id"`
	Topic_title         string      `json:"topic_title"`
	Body                string      `json:"body"`
	HTMLBody            string      `json:"html_body"`
	HTMLFooter          string      `json:"html_footer"`
	Created_at          time.Time   `json:"created_at"`
	Comments_count      int         `json:"comments_count"`
	Forum               Forums      `json:"forum"`
	User                UserFriends `json:"user"`
	Type                string      `json:"type"`
	Linked_id           int         `json:"linked_id"`
	Linked_type         string      `json:"linked_type"`
	Linked              Animanga    `json:"linked"`
	Viewed              bool        `json:"viewed"`
	Last_comment_viewed bool        `json:"last_comment_viewed"`
}

type TopicsUpdates struct {
	Id         int       `json:"id"`
	Linked     Animanga  `json:"linked"`
	Event      string    `json:"event"`
	Episode    int       `json:"episode"`
	Created_at time.Time `json:"created_at"`
	Url        string    `json:"url"`
}

// TopicsId solution. In this struct there are both anime and manga fields.
type AnimangaId struct {
	Id             int         `json:"id"`
	Name           string      `json:"name"`
	Russian        string      `json:"russian"`
	Image          AnimeSize   `json:"image"`
	Url            string      `json:"url"`
	Kind           string      `json:"kind"`
	Score          string      `json:"score"`
	Status         string      `json:"status"`
	Episodes       int         `json:"episodes"`
	Episodes_aired int         `json:"episodes_aired"`
	Volumes        int         `json:"volumes"`
	Chapters       int         `json:"chapters"`
	Aired_on       string      `json:"aired_on"`
	Released_on    string      `json:"released_on"`
	User           UserFriends `json:"user"`
	Votes_count    int         `json:"votes_count"`
	Votes_for      int         `json:"votes_for"`
	Body           string      `json:"body"`
	HTMLBody       string      `json:"html_body"`
	Overall        int         `json:"overall"`
	Storyline      int         `json:"storyline"`
	Music          int         `json:"music"`
	Characters     int         `json:"characters"`
	Animation      int         `json:"animation"`
	Created_at     time.Time   `json:"created_at"`
}

// # Removed:
//
//	Event: nil;
//	Episode: nil;
type TopicsId struct {
	Id                  int         `json:"id"`
	Topic_title         string      `json:"topic_title"`
	Body                string      `json:"body"`
	HTMLBody            string      `json:"html_body"`
	HTMLFooter          string      `json:"html_footer"`
	Created_at          time.Time   `json:"created_at"`
	Comments_count      int         `json:"comments_count"`
	Forum               Forums      `json:"forum"`
	User                UserFriends `json:"user"`
	Type                string      `json:"type"`
	Linked_id           int         `json:"linked_id"`
	Linked_type         string      `json:"linked_type"`
	Linked              AnimangaId  `json:"linked"`
	Viewed              bool        `json:"viewed"`
	Last_comment_viewed bool        `json:"last_comment_viewed"`
}

type Character struct {
	Id                int               `json:"id"`
	Name              string            `json:"name"`
	Russian           string            `json:"russian"`
	Image             AnimeSize         `json:"image"`
	Url               string            `json:"url"`
	Altname           string            `json:"altname"`
	Japanese          string            `json:"japanese"`
	Description       string            `json:"description"`
	DescriptionHTML   string            `json:"description_html"`
	DescriptionSource string            `json:"description_source"`
	Favoured          bool              `json:"favoured"`
	Thread_id         int               `json:"thread_id"`
	Topic_id          int               `json:"topic_id"`
	Updated_at        time.Time         `json:"updated_at"`
	Seyu              []AllPeople       `json:"seyu"`
	Animes            []CharacterAnimes `json:"animes"`
	Mangas            []CharacterMangas `json:"mangas"`
}

type CharacterAnimes struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Russian        string    `json:"russian"`
	Image          AnimeSize `json:"image"`
	Url            string    `json:"url"`
	Kind           string    `json:"kind"`
	Score          string    `json:"score"`
	Status         string    `json:"status"`
	Episodes       int       `json:"episodes"`
	Episodes_aired int       `json:"episodes_aired"`
	Aired_on       string    `json:"aired_on"`
	Released_on    string    `json:"released_on"`
	Roles          []string  `json:"roles"`
	Role           string    `json:"role"`
}

type CharacterMangas struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Russian     string    `json:"russian"`
	Image       AnimeSize `json:"image"`
	Url         string    `json:"url"`
	Kind        string    `json:"kind"`
	Score       string    `json:"score"`
	Status      string    `json:"status"`
	Volumes     int       `json:"volumes"`
	Chapters    int       `json:"chapters"`
	Aired_on    string    `json:"aired_on"`
	Released_on string    `json:"released_on"`
	Roles       []string  `json:"roles"`
	Role        string    `json:"role"`
}
