package main

import (
	irc "github.com/fluffle/goirc/client"
)

type PackSource struct {
	host Host
	Bots []Bot
}

func (s PackSource) Host() Host {
	return s.host
}

type Bot struct {
	Username string
	Dialect Dialect
}

type Dialect struct {
	PoundSign bool
}

func (s PackSource) Connect(c *irc.Conn) error {
	return nil
}

func (s PackSource) CanQuery() bool {
	return false
}

func (s PackSource) Query() []string {
	return []string{}
}

