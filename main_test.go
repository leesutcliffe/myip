package main

import (
	// "net/http"
	// "net/http/httptest"
	"testing"
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

		var want error = nil
		got := startApp(&testVerConf)
		if got != want {
			t.Errorf("genUrl() = %q, want %q", got, want)
		}
}

//TODO: handler.Get(url) - dependency on hander package
// look to pass function by reference and mock using dummy
// func in test

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
