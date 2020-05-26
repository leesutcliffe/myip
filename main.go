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

// initialise vars
var (
	ver        string
	outUsage   string
	outDefault string
	outFlag    string
	verUsage   string
	verDefault bool
	verFlag    bool
)

// define global vars in init function
func init() {
	ver = "0.1"
	outUsage = "text or json output {text|json}"
	outDefault = "text"
	verUsage = "version number"
	verDefault = false

	flag.StringVar(&outFlag, "output", outDefault, outUsage)
	flag.StringVar(&outFlag, "o", outDefault, outUsage)
	flag.BoolVar(&verFlag, "version", verDefault, verUsage)
	flag.BoolVar(&verFlag, "v", verDefault, verUsage)
}

func main() {
	err := StartApp()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func StartApp() error {

	flag.Parse()
	if outFlag == "" {
		return fmt.Errorf("unable to parse flags")
	}

	// output version and exit if flag selected
	if verFlag {
		version(ver)
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
func version(ver string) {
	fmt.Printf("myip, version:%v\n", ver)
	os.Exit(1)
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
