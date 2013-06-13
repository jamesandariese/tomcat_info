package tomcat_info

import (
	"encoding/xml"
	"errors"
	"os"
	"github.com/jamesandariese/easy_error"
)

type Valve struct {
	ClassName string `xml:"className,attr"`
	Directory string `xml:"directory,attr"`
	Prefix string `xml:"prefix,attr"`
	Suffix string `xml:"suffix,attr"`
	Pattern string `xml:"pattern,attr"`
	ResolveHosts bool `xml:"resolveHosts,attr"`
}

type Host struct {
	Valves []Valve `xml:"Valve"`

	Name string `xml:"name,attr"`
	AppBase string `xml:"appBase,attr"`
	UnpackWars string `xml:"unpackWARs,attr"`
	AutoDeploy string `xml:"autoDeploy,attr"`
}

type Realm struct {
	Realms []Realm `xml:"Realm"`
	ClassName string `xml:"className,attr"`
	ResourceName string `xml:"resourceName,attr,omitempty"`
}

type Engine struct {
	Realms []Realm `xml:"Realm"`
	Hosts []Host `xml:"Host"`

	Name string `xml:"name,attr"`
	DefaultHost string `xml:"defaultHost,attr"`
}

type Connector struct {
	Port int16 `xml:"port,attr"`
	MaxThreads int32 `xml:"maxThreads,attr,omitempty"`
	Protocol string `xml:"protocol,attr"`
	RedirectPort int16 `xml:"redirectPort,attr"`
}

type Service struct {
	Connectors []Connector `xml:"Connector"`
	Engines []Engine `xml:"Engine"`

	Name string `xml:"name,attr"`
}

type Resource struct {
	Name string `xml:"name,attr"`
	Auth string `xml:"auth,attr"`
	Type string `xml:"type,attr"`
	Description string `xml:"description,attr"`
	Factory string `xml:"factory,attr"`
	Pathname string `xml:"pathname,attr"`
}

type GlobalNamingResources struct {
	Resources []Resource `xml:"Resource"`
}

type Listener struct {
	ClassName string `xml:"className,attr"`
	SslEngine string `xml:"SSLEngine,attr,omitempty"`
}

type Server struct {
	Port     int16 `xml:"port,attr"`
	Shutdown string `xml:"shutdown,attr"`
	Listeners []Listener `xml:"Listener"`
	GlobalNamingResources GlobalNamingResources `xml:"GlobalNamingResources"`
	Services []Service `xml:"Service"`

	serverpath string // this is the path given to the reader.
}

var ErrServerHasNoPort = errors.New("No port is defined for HTTP/1.1")
func (server *Server) GetHttpPort() (port int16, err error) {
	for _, service := range(server.Services) {
		for _, connector := range(service.Connectors) {
			switch connector.Protocol {
			case "HTTP/1.1", "":
				port = connector.Port
				return
			}
		}
	}
	err = ErrServerHasNoPort
	return
}

var ErrServerHasNoUserFile = errors.New("No user file is defined")
func (server *Server) GetUserFile() (path string, err error) {
	for _, v := range(server.GlobalNamingResources.Resources) {
		if v.Name == "UserDatabase" {
			path = v.Pathname
			return
		}
	}
	err = ErrServerHasNoUserFile
	return
}

func ReadServer(serverpath string) (server *Server, err error) {
	defer func() {
		err = easy_error.Apply(recover())
	}()
	file := easy_error.Wrap(os.Open(serverpath + "/conf/server.xml")).(*os.File)
	decoder := xml.NewDecoder(file)
	server = &Server{}
	server.serverpath = serverpath
	decoder.Decode(server)
	return
}