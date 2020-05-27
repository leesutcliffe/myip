package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"myip/util"
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
		fmt.Println(util.Version(ver))
		os.Exit(1)
	}

	// generate url, print flag defaults if incorrect flag
	url := util.GenUrl(outFlag)
	if url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println(getReq(url))

	return nil

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