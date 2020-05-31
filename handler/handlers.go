package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Function - Get
// @Param: url - full URL to ipify API
// returns the HTTP body
func Get(url string) string {
	var fmtBody string
	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	if response != nil {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmtBody = (string(body))
		response.Body.Close()
	}

	return fmtBody
}

// Function - Path
// @Param: output - text or json output as per ipify API
// returns the ipify api endpoint depending on output preference
func Path(output string) string {
	url := "https://api.ipify.org"
	query := "/?format=json"

	if output == strings.ToLower("json") {
		url += query
	} else if output != "text" {
		return ""
	}

	return url
}
