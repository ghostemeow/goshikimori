package genres

import "testing"

func TestMapGenresAnime(t *testing.T) {
	if MapGenresAnime([]int{2, 14, 10, 88, 31, 12, 2, 539, 10, 31, 29}) ==
		"2-Adventure,14-Horror,10-Fantasy,31-Super Power,12-Hentai,539-Erotica,29-Space" {
		t.Log("MapGenresAnime passed")
	} else {
		t.Error("MapGenresAnime failed")
	}
}

func TestMapGenresManga(t *testing.T) {
	if MapGenresManga([]int{49, 58, 66, 45, 49, 540, 78, 78, 85, 88, 63}) ==
		"49-Comedy,58-Magic,66-Martial Arts,540-Erotica,78-Music,85-Space,88-Samurai,63-Shoujo" {
		t.Log("MapGenresManga passed")
	} else {
		t.Error("MapGenresManga failed")
	}
}
