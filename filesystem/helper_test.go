package filesystem

import (
	"testing"
)

func TestBasename(t *testing.T) {
	var str string = "/usr/local/ttt/work/one.jpg"
	var ext string
	ext = Basename(str, ".jpg")
	t.Log(ext)
	ext = Basename("/", "")
	t.Log(ext)
	ext = Basename("/usr/local", "")
	t.Log(ext)
	ext = Basename(".", "")
	t.Log(ext)
}