package main

import (
	irc "github.com/fluffle/goirc/client"
)

type Host struct {
	Server string
	Channel string
}

type Source interface {
	Host() Host
	Connect(connection *irc.Conn) error
	CanQuery() bool
	Query() []string
}

