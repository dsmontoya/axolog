package client

import (
	"net"
	"net/url"
)

type Container struct {
	ID     string            `id`
	Image  string            `image`
	State  string            `state`
	Status string            `status`
	Name   string            `name`
	Labels map[string]string `labels`
}

type Containers map[string]*Container

func (c *Container) Stream(uri *url.URL) (*Stream, error) {
	conn, err := net.Dial(uri.Scheme, uri.Host)
	if err != nil {
		return nil, err
	}
	return &Stream{c, conn}, nil
}

func (c Containers) Remove(id string) {
	c[id] = nil
}
