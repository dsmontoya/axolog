package client

import (
	"encoding/json"
	"net"
	"net/url"
)

type Data struct {
	Message   string     `json:"message"`
	Container *Container `json:"docker"`
}

type Stream struct {
	Container *Container
	URI       *url.URL
}

func (s *Stream) Write(p []byte) (int, error) {
	conn, err := net.Dial(s.URI.Scheme, s.URI.Host)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	data := &Data{
		Message:   string(p),
		Container: s.Container,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	return conn.Write(b)
}
