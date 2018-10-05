package client

import (
	"net"
	"net/url"
)

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

	//simple write
	return conn.Write(p)
}
