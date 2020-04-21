package filesystem

import (
	"encoding/base64"
	"net/url"
	"strings"
)

// Basename returns the file name part of the path
func Basename(path, suffix string) string {
	// discard last '/'
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			path = path[i+1:]
			break
		}
	}
	// when path value is /
	if len(path) == 0 {
		return ""
	}

	//when suffix is not empty string
	if suffix != "" {
		if strings.Contains(path, suffix) {
			slash := strings.LastIndex(path, suffix)
			return path[:slash]
		}
	}
	return path
}

// url encode string, is + not %20
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// url decode string
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

// base64 encode
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64 decode
func Base64Decode(str string) (string, error) {
	s, e := base64.StdEncoding.DecodeString(str)
	return string(s), e
}