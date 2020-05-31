package main

import (
	// "net/http"
	// "net/http/httptest"
	"testing"
	"errors"
)

// func TestGetRequest(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		res.Write([]byte(`OK`))
// 	}))
// 	defer ts.Close()

// 	want := "OK"
// 	got := getReq(ts.URL)
// 	if got != want {
// 		t.Errorf("getReq() = %q, want %q", got, want)
// 	}
// }

func TestStartAppVersion(t *testing.T) {
		var testVerConf Config
		testVerConf.output = "text"
		testVerConf.version = true

		getMock := func(u string) string {
			return u
		}

		var want error = nil
		got := startApp(&testVerConf, getMock)
		if got != want {
			t.Errorf("genUrl() = %q, want %q", got, want)
		}
}

func TestStartApp(t *testing.T) {
	var testVerConf Config
	testVerConf.output = "text"
	testVerConf.version = false

	getMock := func(u string) string {
		return u
	}

	var want error = nil
	got := startApp(&testVerConf, getMock)
	if got != want {
		t.Errorf("genUrl() = %q, want %q", got, want)
	}
}

// func TestStartAppError(t *testing.T) {
// 	var testVerConf Config
// 	testVerConf.output = "invalid"
// 	testVerConf.version = false

// 	getMock := func(u string) string {
// 		return u
// 	}

// 	var want error = errors.New("invalid flag")
// 	got := startApp(&testVerConf, getMock)
// 	if got != want {
// 		t.Errorf("genUrl() = %q, want %q", got, want)
// 	}
// }

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
