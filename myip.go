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

  url := "https://api.ipify.org"
  query := "/?format=json"
  output := flag.String("output", "text", "text or json output for ipify api")

  flag.Parse()

  if (*output == strings.ToLower("json")) {
    url += query
  } else if *output {

  } else if (*output != "text") {
    log.Fatal("invalid flag ", *output)
  }

  resp, err := http.Get(url)
  if err != nil {
    log.Fatal("Error getting response. ", err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal("Error reading response. ", err)
  }

  fmt.Printf("%s\n", body)
}
