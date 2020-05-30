package main

import (
	"errors"
	"flag"
	"fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	"os"
	"strings"
	"myip/handler"
)

type Config struct {
	version bool
	output  string
}

// myip version
var ver = "0.1"

func startApp(conf *Config) error {

	if conf.version {
		fmt.Printf("myip, version: %v\n", ver)
		return nil
	}

	if url := genUrl(conf.output); url != "" {
		fmt.Println(handler.Get(url))
		return nil
	}

	return errors.New("invalid flag")

}

func main() {

	var conf Config
	flag.StringVar(&conf.output, "output", outDefault, outUsage)
	flag.StringVar(&conf.output, "o", outDefault, outUsage)
	flag.BoolVar(&conf.version, "version", verDefault, verUsage)
	flag.BoolVar(&conf.version, "v", verDefault, verUsage)

	flag.Parse()

	err := startApp(&conf)
	if err != nil {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

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
