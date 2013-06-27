package main

import (
	"fmt"
	"flag"
	"github.com/jamesandariese/tomcat_info"
	"github.com/jamesandariese/easy_error"
	"os"
)

var serverdir = flag.String("d", ".", "Path to server (conf/server.xml should exist at this path)")
var server *tomcat_info.Server
func getServer() *tomcat_info.Server {
	if server == nil {
		server = easy_error.Wrap(tomcat_info.ReadServer(*serverdir)).(*tomcat_info.Server)
	}
	return server
}
var tomcatusers *tomcat_info.TomcatUsers
func getTomcatUsers() *tomcat_info.TomcatUsers {
	if tomcatusers == nil {
		tomcatusers = easy_error.Wrap(getServer().ReadTomcatUsers()).(*tomcat_info.TomcatUsers)		
	}
	return tomcatusers
}



var bestUserWithRole = flag.String("r", "", "Print a single user and password with this role, preferring a user with the same name as the role")
func doBestUserWithRole(role string) {
	user := getTomcatUsers().GetBestUserWithRole(role)
	if user == nil {
		fmt.Fprintf(os.Stderr, "No user matching role exists: %s\n", role)
		os.Exit(1)
	}
	fmt.Printf("%s %s\n", user.Username, user.Password)
}

var usersWithRole = flag.String("R", "", "Print all users with role along with their password")
func doUsersWithRole(role string) {
	users := getTomcatUsers().GetUsersWithRole(role)
	if users == nil {
		fmt.Fprintf(os.Stderr, "No users matching role exists: %s\n", role)
		os.Exit(1)
	}
	for _, v := range(users) {
		fmt.Printf("%s %s\n", v.Username, v.Password)
	}
}

var getServerPort = flag.Bool("p", false, "Print the port that tomcat is listening on")
func doGetServerPort() {
        server := easy_error.Wrap(tomcat_info.ReadServer(*serverdir)).(*tomcat_info.Server)
        port := easy_error.Wrap(server.GetHttpPort()).(uint16)
        fmt.Println(port)
}

func main() {
	defer func() {
		err := easy_error.Apply(recover())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}()

	flag.Parse()

	if *bestUserWithRole != "" {
		doBestUserWithRole(*bestUserWithRole)
	}
	if *usersWithRole != "" {
		doUsersWithRole(*usersWithRole)
	}
	if *getServerPort {
		doGetServerPort()
	}
}