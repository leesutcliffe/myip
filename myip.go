package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "flag"
  "strings"
)

func main() {

  // command line args - 'output' default value 'text'
  outputPtr := flag.String("output", "text", "text or json output for ipify api")
  flag.Parse()

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
