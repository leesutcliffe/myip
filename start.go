package main

import (
	"fmt"
	hnd "myip/handler"
	"errors"
)

func startApp(conf *Config, get func(string) string) error {
	if conf.version {
		fmt.Printf("myip, version: %v\n", ver)
		return nil
	}

	if url := hnd.Path(conf.output); url != "" {
		fmt.Println(get(url))
		return nil
	}

	return errors.New("invalid flag")
}