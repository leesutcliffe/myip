package util

import "strings"

// returns the URL and option query if JSON format is requested
// takes output pointer as argument
func GenUrl(output string) string {
	url := "https://api.ipify.org"
	query := "/?format=json"

	if output == strings.ToLower("json") {
		url += query
	} else if output != "text" {
		return ""
	}

	return url
}

// show version and exit
func Version(ver string) string {
	return "myip, version: "+ver
}