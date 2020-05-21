package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "flag"
  "strings"
  "os"
  ver "github.com/leesutcliffe/myip/version"
)

func main() {

  var ver = "0.1"
  const (
    outUsage = "text or json output {text|json}"
    outDefault = "text"
    verUsage = "version number"
    verDefault = false
  )

  // cmd line flags
  var verFlag bool
  var outFlag string

  flag.StringVar(&outFlag, "output", outDefault, outUsage)
  flag.StringVar(&outFlag, "o", outDefault, outUsage)
  flag.BoolVar(&verFlag, "version", verDefault, verUsage)
  flag.BoolVar(&verFlag, "v", verDefault, verUsage)
  flag.Parse()

  // output version and exit if flag selected
  if verFlag {
    showVer(ver)
  }

  url := generateUrl(outFlag)
  fmt.Println(getRequest(url))

}

// gets HTTP request to ipify api
// returns http body
func getRequest(url string) string {
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

// Prints version and exits
func showVer(ver string){
    fmt.Printf("myip, version:%v\n", ver)
    os.Exit(1)
}

// returns the URL and option query if JSON format is requested
// takes output pointer as argument
func generateUrl(output string) string {
  url := "https://api.ipify.org"
  query := "/?format=json"

  if (output == strings.ToLower("json")) {
    url += query
  } else if (output != "text") {
    flag.PrintDefaults()
    os.Exit(1)
  }

  return url
}
