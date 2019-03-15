package main

import (
	"flag"
)

var _version string
var v bool
var version bool

func init() {
	flag.BoolVar(&v, "v", false, "print version")
	flag.BoolVar(&version, "version", false, "print version")
	flag.Parse()
}
func main() {
	if v || version {
		println(_version)
	}
}
