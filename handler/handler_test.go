package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`OK`))
	}))
	defer ts.Close()

	want := "OK"
	got := Get(ts.URL)
	if got != want {
		t.Errorf("getReq() = %q, want %q", got, want)
	}
}

func TestDefaultUrl(t *testing.T) {
	expected := "https://api.ipify.org"
	actual := Path("text")
	if actual != expected {
		t.Errorf("Path() = %q, expected %q", actual, expected)
	}
}

func TestJsonUrl(t *testing.T) {
	expected := "https://api.ipify.org/?format=json"
	actual := Path("json")
	if actual != expected {
		t.Errorf("Path() = %q, expected %q", actual, expected)
	}
}

func TestUrlErr(t *testing.T) {
	expected := ""
	actual := Path("invalid")
	if actual != expected {
		t.Errorf("Path() = %q, expected %q", actual, expected)
	}
}


