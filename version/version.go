package ver

import (
	"fmt"
	"os"
)

// Prints version and exits
func ShowVer(ver string){
	fmt.Printf("myip, version:%v\n", ver)
	os.Exit(1)
}