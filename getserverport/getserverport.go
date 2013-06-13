package main

import (
	"fmt"
	"github.com/jamesandariese/tomcat_info"
	"flag"
)

var serverdir = flag.String("d", ".", "Path to server (conf/server.xml should exist in this dir)")
func main() {
	flag.Parse()
	fmt.Println(tomcat_info.GetServerHttpPort(*serverdir + "/conf/server.xml"))
}