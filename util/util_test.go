package util

import "testing"

func TestDefaultUrl(t *testing.T) {
	want := "https://api.ipify.org"
	got := GenUrl("text")
	if got != want {
		t.Errorf("genUrl() = %q, want %q", got, want)
	}
}

func TestJsonUrl(t *testing.T) {
	want := "https://api.ipify.org/?format=json"
	got := GenUrl("json")
	if got != want {
		t.Errorf("genUrl() = %q, want %q", got, want)
	}
}

func TestUrlErr(t *testing.T) {
	want := ""
	got := GenUrl("invalid")
	if got != want {
		t.Errorf("genUrl() = %q, want %q", got, want)
	}
}

func TestVersion(t *testing.T) {
	want := "myip, version: 0.1"
	got := Version("0.1")
	if got != want {
		t.Errorf("version() = %q, want %q", got, want)
	}
}