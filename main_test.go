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