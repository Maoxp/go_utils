package mtime

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	now := time.Now()
	date := Date("Y-m-d H:i:s", now)
	t.Log(date)
	date = Date("Y-m-d", now)
	t.Log(date)
	date = Date("H:i:s", now)
	t.Log(date)
	date = Date("W", now)
	t.Log(date)
	date = Date("z", now)
	t.Log(date)
}

func TestFromUnixTime(t *testing.T) {
	s := FromUnixTime(int64(1587366437), "Y-m-d H:i:s")
	t.Log(s)
}

func TestDateToUnix(t *testing.T) {
	ux := DateToUnix("2020-04-20", "Y-m-d")
	t.Log(ux)
}