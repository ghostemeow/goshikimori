package concatination

import (
	"bytes"
	"testing"
)

func TestIdsToString(t *testing.T) {
	if IdsToString([]int{1336, 1337, 1338}) == "1336,1337,1338" {
		t.Log("IdsToString passed")
	} else {
		t.Error("IdsToString failed")
	}
}

func TestNekoSliceToString(t *testing.T) {
	if NekoSliceToString([]string{"initial", "d", "first", "stage"}) == "initial_d_first_stage" {
		t.Log("NekoSliceToString passed")
	} else {
		t.Error("NekoSliceToString failed")
	}
}

func TestUrl(t *testing.T) {
	if Url(46, []string{"https://shikimori.one/api/", "users/", "search=arctica"}) ==
		"https://shikimori.one/api/users/search=arctica" {
		t.Log("Url apssed")
	} else {
		t.Error("Url failed")
	}
}

func TestBearer(t *testing.T) {
	if Bearer("XXX_TOKEN_XXX") == "Bearer XXX_TOKEN_XXX" {
		t.Log("Bearer passed")
	} else {
		t.Error("Bearer failed")
	}
}

func TestDataBuffer(t *testing.T) {
	if bytes.Equal(DataBuffer([]string{"zero", "one", "1337"}), []byte("zeroone1337")) {
		t.Log("DataBuffer passed")
	} else {
		t.Error("DataBuffer failed")
	}
}

func TestDataCopy(t *testing.T) {
	if bytes.Equal(DataCopy(11, []string{"zero", "one", "1337"}), []byte("zeroone1337")) {
		t.Log("DataCopy passed")
	} else {
		t.Error("DataCopy failed")
	}
}
