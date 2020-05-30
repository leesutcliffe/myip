package main

// initialise vars
var (
	outUsage   string
	outDefault string
	verUsage   string
	verDefault bool
	//url 		   string
)

// define global vars in init function
func init() {
	outUsage = "text or json output {text|json}"
	outDefault = "text"
	verUsage = "version number"
	verDefault = false

	// flag.StringVar(&outFlag, "output", outDefault, outUsage)
	// flag.StringVar(&outFlag, "o", outDefault, outUsage)
	// flag.BoolVar(&verFlag, "version", verDefault, verUsage)
	// flag.BoolVar(&verFlag, "v", verDefault, verUsage)
	//flag.Parse()
}
