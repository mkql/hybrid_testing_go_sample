package hoge

import "testing"
func TestMessage(t *testing.T) {
	want := "hello"
	got := message()
	if got != want {
		t.Fatalf("want %s, but got %s", want, got)
	}
}