package main

import "testing"

func TestDefaultUrl(t *testing.T) {
	want := "https://api.ipify.org"
	got := genUrl("text")
	if got != want {
		t.Errorf("genUrl() = %q, want %q", got, want)
	}
}

func TestJsonUrl(t *testing.T) {
	want := "https://api.ipify.org/?format=json"
	got := genUrl("json")
	if got != want {
		t.Errorf("genUrl() = %q, want %q", got, want)
	}
}

func TestUrlErr(t *testing.T) {
	want := ""
	got := genUrl("invalid")
	if got != want {
		t.Errorf("genUrl() = %q, want %q", got, want)
	}
}