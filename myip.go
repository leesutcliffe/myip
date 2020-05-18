package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "flag"
  "strings"
  "os"
)

func main() {

  var ver float64 = 0.1

  // command line args - 'output' default value 'text'
  outputPtr := flag.String("output", "text", "text or json output for ipify api")
  verPtr := flag.Bool("version", false, "output version number")
  flag.Parse()

  showVer(ver, verPtr)

  url := getUrl(outputPtr)

  response, err := http.Get(url)
  if err != nil {
    log.Fatalln(err)
  }

  if response != nil {

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
      log.Fatalln(err)
    }

    fmt.Println(string(body))
    response.Body.Close()
  }
}

// Prints version if flag set to true and exits
func showVer(ver float64, verFlag *bool){
  if *verFlag {
    fmt.Printf("myip, version:%v\n", ver)
    os.Exit(1)
  }
}

// returns the URL and option query if JSON format is requested
// takes output pointer as argument
func getUrl(output *string) string {
  url := "https://api.ipify.org"
  query := "/?format=json"

  if (*output == strings.ToLower("json")) {
    url += query
  } else if (*output != "text") {
    log.Fatal("invalid flag: ", *output)
  }

  return url
}
