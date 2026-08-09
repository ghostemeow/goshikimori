package concatenation

import (
	"bytes"
	"strconv"
	"strings"
)

// Convert a slice with an ids into a string.
func IdsToString(slice []int) string {
	var res bytes.Buffer
	for i := range slice {
		res.WriteString(strconv.Itoa(slice[i]))
		res.WriteString(",")
	}
	return strings.TrimSuffix(res.String(), ",")
}

// Convert a slice with an words into a string.
func NekoSliceToString(slice []string) string {
	var res bytes.Buffer
	for i := range slice {
		res.WriteString(slice[i])
		res.WriteString("_")
	}
	return strings.TrimSuffix(res.String(), "_")
}

// Quick creation of a url.
//
// maxLen is only a hint for pre-allocation: if the actual length of the
// joined parts exceeds it, the result is grown instead of truncated.
func Url(maxLen int, slice []string) string {
	total := 0
	for _, s := range slice {
		total += len(s)
	}
	if maxLen < total {
		maxLen = total
	}
	var offset int
	res := make([]byte, maxLen)
	for i := range slice {
		offset += copy(res[offset:], slice[i])
	}
	return string(res[:offset])
}

// Quick creation of a bearer token.
func Bearer(token string) string {
	res := make([]byte, 7+len(token))
	copy(res, "Bearer ")
	copy(res[7:], token)
	return string(res)
}

// Converting a slice to a []byte using a bytes.Buffer.
func DataBuffer(slice []string) []byte {
	var res bytes.Buffer
	for i := range slice {
		res.WriteString(slice[i])
	}
	return res.Bytes()
}

// Converting a slice to a []byte using a copy.
//
// maxLen is only a hint for pre-allocation: if the actual length of the
// joined parts exceeds it, the result is grown instead of truncated.
func DataCopy(maxLen int, slice []string) []byte {
	total := 0
	for _, s := range slice {
		total += len(s)
	}
	if maxLen < total {
		maxLen = total
	}
	var offset int
	res := make([]byte, maxLen)
	for i := range slice {
		offset += copy(res[offset:], []byte(slice[i]))
	}
	return res[:offset]
}
