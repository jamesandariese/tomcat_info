package main

import (
	"fmt"
	"github.com/jamesandariese/tomcat_info"
	"github.com/jamesandariese/easy_error"
	"flag"
	"os"
)

var serverdir = flag.String("d", ".", "Path to server (conf/server.xml should exist in this dir)")
func main() {
	defer func() {
		err := easy_error.Apply(recover())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}()
	flag.Parse()
	server := easy_error.Wrap(tomcat_info.ReadServer(*serverdir)).(*tomcat_info.Server)
	tomcatusers := easy_error.Wrap(server.ReadTomcatUsers()).(*tomcat_info.TomcatUsers)
	usermap := tomcatusers.GetUsersWithRole("manager")
	manager, ok := usermap["manager"]
	if ok {
		fmt.Println(manager.Password)
		return
	} else {
		for _, v := range(usermap) {
			fmt.Println(v.Password)
			return
		}
	}
	fmt.Println("No manager user exists")
	os.Exit(1)
}