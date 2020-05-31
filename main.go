package main

import (
	"flag"
	hnd "myip/handler"
	"os"
)

type Config struct {
	version bool
	output  string
}

// initialise vars
var (
	outUsage   string
	outDefault string
	verUsage   string
	verDefault bool
	ver        string
	//conf       Config
)

// define global vars in init function
func init() {
	outUsage = "text or json output {text|json}"
	outDefault = "text"
	verUsage = "version number"
	verDefault = false

	// myip version
	ver = "0.1"
}

func main() {

	var conf Config
	flag.StringVar(&conf.output, "output", outDefault, outUsage)
	flag.StringVar(&conf.output, "o", outDefault, outUsage)
	flag.BoolVar(&conf.version, "version", verDefault, verUsage)
	flag.BoolVar(&conf.version, "v", verDefault, verUsage)

	flag.Parse()

	err := startApp(&conf, hnd.Get)
	if err != nil {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

