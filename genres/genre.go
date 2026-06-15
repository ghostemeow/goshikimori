package genres

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/ghostemeow/goshikimori/internal/concatination"
)

type Genres struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Russian    string `json:"russian"`
	Kind       string `json:"kind"`
	Entry_type string `json:"entry_type"`
}

var (
	// genre v2.
	GenreAnime = map[int]string{
		1:   "1-Action",
		2:   "2-Adventure",
		3:   "3-Cars",
		4:   "4-Comedy",
		5:   "5-Dementia",
		6:   "6-Demons",
		7:   "7-Mystery",
		8:   "8-Drama",
		9:   "9-Ecchi",
		10:  "10-Fantasy",
		11:  "11-Game",
		12:  "12-Hentai",
		13:  "13-Historical",
		14:  "14-Horror",
		15:  "15-Kids",
		16:  "16-Magic",
		17:  "17-Martial Arts",
		18:  "18-Mecha",
		19:  "19-Music",
		20:  "20-Parody",
		21:  "21-Samurai",
		22:  "22-Romance",
		23:  "23-School",
		24:  "24-Sci-Fi",
		25:  "25-Shoujo",
		26:  "26-Shoujo Ai",
		27:  "27-Shounen",
		28:  "28-Shounen Ai",
		29:  "29-Space",
		30:  "30-Sports",
		31:  "31-Super Power",
		32:  "32-Vampire",
		33:  "33-Yaoi",
		34:  "34-Yuri",
		35:  "35-Harem",
		36:  "36-Slice of Life",
		37:  "37-Supernatural",
		38:  "38-Military",
		39:  "39-Police",
		40:  "40-Psychological",
		41:  "41-Thriller",
		42:  "42-Seinen",
		43:  "43-Josei",
		539: "539-Erotica",
		541: "541-Work Life",
		543: "543-Gourmet",
	}

	GenreManga = map[int]string{
		46:  "46-Mystery",
		47:  "47-Shounen",
		48:  "48-Supernatural",
		49:  "49-Comedy",
		50:  "50-Drama",
		51:  "51-Ecchi",
		52:  "52-Seinen",
		53:  "53-Sci-Fi",
		54:  "54-Slice of Life",
		55:  "55-Shounen Ai",
		56:  "56-Action",
		57:  "57-Fantasy",
		58:  "58-Magic",
		59:  "59-Hentai",
		60:  "60-School",
		61:  "61-Doujinshi",
		62:  "62-Romance",
		63:  "63-Shoujo",
		64:  "64-Vampire",
		65:  "65-Yaoi",
		66:  "66-Martial Arts",
		67:  "67-Psychological",
		68:  "68-Adventure",
		69:  "69-Historical",
		70:  "70-Military",
		71:  "71-Harem",
		72:  "72-Demons",
		73:  "73-Shoujo Ai",
		74:  "74-Gender Bender",
		75:  "75-Yuri",
		76:  "76-Sports",
		77:  "77-Kids",
		78:  "78-Music",
		79:  "79-Game",
		80:  "80-Horror",
		81:  "81-Thriller",
		82:  "82-Super Power",
		83:  "83-Mecha",
		84:  "84-Cars",
		85:  "85-Space",
		86:  "86-Parody",
		87:  "87-Josei",
		88:  "88-Samurai",
		89:  "89-Police",
		90:  "90-Dementia",
		540: "540-Erotica",
		542: "542-Work Life",
		544: "544-Gourmet",
	}
)

// Write key to slice and check for duplicates.
func checkForDuplicates(target int, slice []int) bool {
	for i := 0; i < 50; i++ {
		if slice[i] == target {
			return false
		}
	}
	return true
}

// Anime value map search.
func MapGenresAnime(slice []int) string {
	var res bytes.Buffer
	var count int
	tempSlice := make([]int, 50)

	for i := 0; i < len(slice); i++ {
		_, ok := GenreAnime[slice[i]]
		if ok && checkForDuplicates(slice[i], tempSlice) {
			res.WriteString(GenreAnime[slice[i]])
			res.WriteString(",")
			tempSlice[count] = slice[i]
			count++
		}
	}

	return strings.TrimSuffix(res.String(), ",")
}

// Manga value map search.
func MapGenresManga(slice []int) string {
	var res bytes.Buffer
	var count int
	tempSlice := make([]int, 50)

	for i := 0; i < len(slice); i++ {
		_, ok := GenreManga[slice[i]]
		if ok && checkForDuplicates(slice[i], tempSlice) {
			res.WriteString(GenreManga[slice[i]])
			res.WriteString(",")
			tempSlice[count] = slice[i]
			count++
		}
	}

	return strings.TrimSuffix(res.String(), ",")
}

// Auxiliary function to get the correct list of genres.
//
// name:
//
// > GENERATE_GENRES_ANIME, GENERATE_GENRES_MANGA;
//
// genres: []genres.Genres;
func GenerateGenres(name string, genres []Genres) map[int]string {
	data := make(map[int]string)
	for _, v := range genres {
		if v.Entry_type == name {
			data[v.Id] = string(concatination.DataBuffer(
				[]string{strconv.Itoa(v.Id), "-", name},
			))
		}
	}
	return data
}
