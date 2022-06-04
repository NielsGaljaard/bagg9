package main

import (
	"fmt"

	reader "gitlab.com/galjaardniels/bagg9/internal/config-reader"
)

func main() {
	conf, err := reader.ReadConfig()
	if err != nil {
		fmt.Println("cant read config")
	}
	fmt.Println(conf)
}
