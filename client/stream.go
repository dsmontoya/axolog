package client

import (
	"encoding/json"
	"net"
	"reflect"

	log "github.com/sirupsen/logrus"
)

type Data struct {
	Message   string     `json:"message"`
	Container *Container `json:"-"`
	Docker    map[string]interface{}
}

type Stream struct {
	Container *Container
	conn      net.Conn
}

func (d *Data) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"message": d.Message,
	}
	v := reflect.ValueOf(d.Container)
	el := v.Elem()
	n := el.NumField()
	for i := 0; i < n; i++ {
		tag := string(el.Type().Field(i).Tag)

		f := el.Field(i).Interface()
		data["docker."+tag] = f
	}

	return json.Marshal(data)
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
		log.Errorf("%s: %s | bytes size %d", s.Container.Name, err, len(b))
	}
	log.Debugf("Log sent from %s. Log size %d. Final message size %d.", s.Container.Name, len(p), len(b))
	return len(p), nil
}
