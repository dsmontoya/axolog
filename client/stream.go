package client

import (
	"encoding/json"
	"net"

	log "github.com/sirupsen/logrus"
)

type Data struct {
	Message   string     `json:"message"`
	Container *Container `json:"-,"`
}

type Stream struct {
	Container *Container
	conn      net.Conn
}

func (s *Stream) Write(p []byte) (int, error) {
	data := &Data{
		Message:   string(p),
		Container: s.Container,
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	if _, err = s.conn.Write(b); err != nil {
		log.Error(err)
		return 0, err
	}
	return len(p), nil
}
