package main

import (
	"fmt"
	"github.com/jamesandariese/tomcat_info"
	"github.com/jamesandariese/easy_error"
	"flag"
)

var serverdir = flag.String("d", ".", "Path to server (conf/server.xml should exist in this dir)")
func main() {
	defer func() {
		err := easy_error.Apply(recover())
		if err != nil {
			fmt.Println(err)
		}
	}()
	flag.Parse()
	server := easy_error.Wrap(tomcat_info.ReadServer(*serverdir)).(*tomcat_info.Server)
	port := easy_error.Wrap(server.GetHttpPort()).(int16)
	fmt.Println(port)
}