package client

import "net/url"

type Container struct {
	ID     string            `json:"id"`
	Image  string            `json:"image"`
	State  string            `json:"state"`
	Status string            `json:"status"`
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

type Containers map[string]*Container

func (c *Container) Stream(uri *url.URL) *Stream {
	return &Stream{c, uri}
}

func (c Containers) Remove(id string) {
	c[id] = nil
}
