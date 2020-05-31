package main

import (
	"testing"
)

func TestStartAppVersion(t *testing.T) {
	var testVerConf Config
	testVerConf.output = "text"
	testVerConf.version = true

	// function signature to mock the HTTP get request
	getMock := func(u string) string {
		return u
	}

	var expected error = nil
	actual := startApp(&testVerConf, getMock)
	if actual != expected {
		t.Errorf("startApp() = %q, expected %q", actual, expected)
	}
}

func TestStartApp(t *testing.T) {
	var testVerConf Config
	testVerConf.output = "text"
	testVerConf.version = false

	// function signature to mock the HTTP get request
	getMock := func(u string) string {
		return u
	}

	var expected error = nil
	actual := startApp(&testVerConf, getMock)
	if actual != expected {
		t.Errorf("startApp() = %q, expected %q", actual, expected)
	}
}

func TestStartAppError(t *testing.T) {
	var testVerConf Config
	testVerConf.output = "invalid"
	testVerConf.version = false

	// function signature to mock the HTTP get request
	getMock := func(u string) string {
		return u
	}

	expected := "invalid flag"
	actual := startApp(&testVerConf, getMock)
	if actual.Error() != expected {
		t.Errorf("startApp() = %q, expected %q", actual, expected)
	}
}
