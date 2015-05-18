package server

import (
	"github.com/mondough/typhon/auth"
)

// Server is an interface that all servers must implement
// so that we can register endpoints, and serve requests
type Server interface {
	Init(*Config)
	Run()
	Close()
	NotifyConnected() chan bool

	Name() string
	Description() string

	RegisterEndpoint(endpoint *Endpoint)
	DeregisterEndpoint(pattern string)

	AuthenticationProvider() auth.AuthenticationProvider
	RegisterAuthenticationProvider(auth.AuthenticationProvider)
}

// DefaultServer stores a default implementation, for simple usage
var DefaultServer Server = NewAMQPServer()

// Init our DefaultServer with a Config
func Init(c *Config) Server {
	DefaultServer.Init(c)
	return DefaultServer
}

// RegisterEndpoint with the DefaultServer
func RegisterEndpoint(endpoint *Endpoint) {
	DefaultServer.RegisterEndpoint(endpoint)
}

// Run the DefaultServer
func Run() {
	DefaultServer.Run()
}

// Close the DefaultServer
func Close() {
	DefaultServer.Close()
}

// NotifyConnected delegates to DefaultServer
func NotifyConnected() chan bool {
	return DefaultServer.NotifyConnected()
}

// Config defines the config a server needs to start up, and serve requests
type Config struct {
	Name        string
	Description string
}
