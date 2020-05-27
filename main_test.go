package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestGetRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`OK`))
	}))
	defer ts.Close()

	want:= "OK"
	got := getReq(ts.URL)
	if got != want {
		t.Errorf("getReq() = %q, want %q", got, want)
	}
}

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

func TestVersion(t *testing.T) {
	want := "myip, version: 0.1"
	got := version("0.1")
	if got != want {
		t.Errorf("version() = %q, want %q", got, want)
	}
}