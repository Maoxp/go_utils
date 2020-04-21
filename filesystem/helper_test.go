package filesystem

import (
	"testing"
)

func TestBasename(t *testing.T) {
	var str string = "/usr/local/ttt/work/one.jpg"
	var ext string
	ext = Basename(str, ".jpg")
	t.Errorf(ext)
	ext = Basename("/", "")
	t.Errorf(ext)
	ext = Basename("/usr/local", "")
	t.Errorf(ext)
	ext = Basename(".", "")
	t.Errorf(ext)
}