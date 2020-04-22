package muuid

import "testing"

func TestGetUUID(t *testing.T) {
	ud := GetUUID()
	t.Log(ud)
}
