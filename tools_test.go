package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var (
		testTools Tools
		want      int = 10
	)
	s := testTools.RandomString(want)
	if len(s) != want {
		t.Error("wrong length random string returned")
	}
}
