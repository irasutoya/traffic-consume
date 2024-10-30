package main

import (
	"fmt"
)

var (
	appName = "traffic-consume"
	version = "1.0.0"
	date    = "1989-06-04"
)

func getVersion() {
	fmt.Printf("%s [%s], built at %s\n", appName, version, date)
}
