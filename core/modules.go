package main

import (
	"fmt"
	"net"
	"strconv"
)

type Module struct {
	id   string
	port int
}

func (m *Module) isPortSet() bool {
	return m.port != -1
}

func setModulePort(m *Module) {
	for !m.isPortSet() {
		port, err := getFreePort()
		if err != nil {
			panic(err)
		}
		m.port = port
	}
}

func DebugModules(modules []Module) {
	for _, m := range modules {
		fmt.Println(m.id + " " + strconv.Itoa(m.port))
	}
}

// Set ports of each module to available port
func InitializeModules(modules []Module) {
	for i := range modules {
		setModulePort(&modules[i])
	}
}

func getFreePort() (int, error) {
	// Find available port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return -1, err
	}
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port, nil
}
