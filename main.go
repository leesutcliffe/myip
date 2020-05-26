package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := StartApp()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func StartApp() error {

	// outFlag & varFlag defined in init.go
	flag.Parse()
	if outFlag == "" {
		return fmt.Errorf("unable to parse flags")
	}

	// output version and exit if flag selected
	if verFlag {
		fmt.Println(version(ver))
		os.Exit(1)
	}

	// generate url, print flag defaults if incorrect flag
	url := genUrl(outFlag)
	if url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println(getReq(url))

	return nil

}

// show version and exit
func version(ver string) string {
	return "myip, version: "+ver
}

// gets HTTP request to ipify api
// returns http body
func getReq(url string) string {
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

// returns the URL and option query if JSON format is requested
// takes output pointer as argument
func genUrl(output string) string {
	url := "https://api.ipify.org"
	query := "/?format=json"

	if output == strings.ToLower("json") {
		url += query
	} else if output != "text" {
		return ""
	}

	return url
}
