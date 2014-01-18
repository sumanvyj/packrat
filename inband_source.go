package main

import (
	irc "github.com/fluffle/goirc/client"
)

type InBandSource struct {
	host Host
	ListCommand string
}

func (s InBandSource) Host() Host {
	return s.host
}

func (s InBandSource) Connect(c *irc.Conn) (err error) {
	if c.Connected {
		return nil
	}

	c.AddHandler(irc.CONNECTED, func (c *irc.Conn, line *irc.Line) {
		c.Join(s.Host().Channel)
	})

	if err := c.Connect(s.Host().Server); err != nil {
		return err
	}


	return nil
}

func (s InBandSource) CanQuery() bool {
	return true
}

func (s InBandSource) Query() []string {
	return []string{}
}



