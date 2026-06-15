package neko

import "testing"

func TestNekoSearch(t *testing.T) {
	name_normal := "Lorem ipsum dolor Sit amet consectetur adipiscing Elit"
	compare_normal := "lorem_ipsum_dolor_sit_amet_consectetur_adipiscing_elit"
	output_normal, _ := Search(name_normal)
	if output_normal == compare_normal {
		t.Log("NekoSearch normal passed")
	} else {
		t.Error("NekoSearch normal failed")
	}

	name_spaces := "Lorem                ipsum             dolor          Sit             "
	compare_spaces := "lorem_ipsum_dolor_sit"
	output_spaces, _ := Search(name_spaces)
	if output_spaces == compare_spaces {
		t.Log("NekoSearch with big spaces passed")
	} else {
		t.Error("NekoSearch with big spaces failed")
	}

	_, err := Search("")
	if err != nil {
		t.Log("NekoSearch empty passed")
	} else {
		t.Error("NekoSearch empty failed")
	}
}
