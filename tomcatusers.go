package tomcat_info

import (
	"encoding/xml"
	"os"
	"github.com/jamesandariese/easy_error"
	"strings"
)

type Role struct {
	Rolename string `xml:"rolename,attr"`
}

type User struct {
	Username string `xml:"username,attr"`
	Password string `xml:"password,attr"`
	Roles string `xml:"roles,attr"`
}

type TomcatUsers struct {
	Roles []Role `xml:"role"`
	Users []User `xml:"user"`
}

func (users *TomcatUsers) GetUsersWithRole(role string) (usermap map[string]User) {
	splitter := func(r rune) bool {
		switch r {
		case ',', ' ', '\t':
			return true
		}
		return false
	}
	usermap = make(map[string]User, 1) // usually just one.
	for _, v := range(users.Users) {
		for _, r := range(strings.FieldsFunc(v.Roles, splitter)) {
			if r == role {
				usermap[v.Username] = v
			}
		}
	}
	return
} 

func (server *Server) ReadTomcatUsers() (users *TomcatUsers, err error) {
	defer func() {
		err = easy_error.Apply(recover())
	}()
	userfile := easy_error.Wrap(server.GetUserFile()).(string)
	file := easy_error.Wrap(os.Open(server.serverpath + "/" + userfile)).(*os.File)
	decoder := xml.NewDecoder(file)
	users = &TomcatUsers{}
	decoder.Decode(users)
	return
}

