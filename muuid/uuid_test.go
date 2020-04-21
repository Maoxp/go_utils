package muuid

import "testing"

func TestGetUUID(t *testing.T) {
	ud := GetUUID()
	t.Fatal(ud)
}
