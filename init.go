package main

import (
	"flag"
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