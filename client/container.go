package client

import "net/url"

type Container struct {
	ID     string
	Image  string
	State  string
	Status string
	Name   string
	Labels map[string]string
}

type Containers map[string]*Container

func (c *Container) Stream(uri *url.URL) *Stream {
	return &Stream{c, uri}
}

func (c Containers) Remove(id string) {
	c[id] = nil
}
