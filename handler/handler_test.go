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

	expected := "OK"
	actual := Get(ts.URL)
	if actual != expected {
		t.Errorf("Get() = %q, expected %q", actual, expected)
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
