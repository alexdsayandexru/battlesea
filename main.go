package main

import "flag"

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "Developer mode")
}

func main() {
	Run()
	//Test2()
}
