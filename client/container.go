package client

import (
	"net"
	"net/url"
)

type Container struct {
	ID     string            `json:"docker.id"`
	Image  string            `json:"docker.image"`
	State  string            `json:"docker.state"`
	Status string            `json:"docker.status"`
	Name   string            `json:"docker.name"`
	Labels map[string]string `json:"docker.labels"`
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
