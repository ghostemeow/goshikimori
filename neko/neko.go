package neko

import (
	"errors"
	"strings"

	"github.com/ghostemeow/goshikimori/internal/concatenation"
)

// String formatting for achievements search. Check [example].
//
// [example]: https://github.com/ghostemeow/goshikimori/blob/master/examples/achievements/main.go
func Search(name string) (string, error) {
	// Extra spaces removed.
	words := strings.Fields(name)
	if len(words) == 0 {
		return "", errors.New("too short string")
	}
	return strings.ToLower(concatenation.NekoSliceToString(words)), nil
}
