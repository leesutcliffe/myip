package main

import (
	"flag"
	hnd "github.com/leesutcliffe/myip/handler"
	"os"
)

// struct for command line options
type Config struct {
	version bool
	output  string
}

// initialise global vars
var (
	outUsage   string
	outDefault string
	verUsage   string
	verDefault bool
	ver        string
)

// global vars
func init() {
	outUsage = "text or json output {text|json}"
	outDefault = "text"
	verUsage = "version number"
	verDefault = false

	// myip version
	ver = "1.0"
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
