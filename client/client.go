package client

import (
	"net/url"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	c          *docker.Client
	URI        *url.URL
	Containers Containers
}

func New(target string) (*Client, error) {
	uri, err := url.ParseRequestURI(target)
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{c: client, URI: uri}, nil
}

func (c *Client) ReadLogs() error {
	for _, container := range c.Containers {
		go c.Logs(container)
	}
	return nil
}

func (c *Client) RegisterContainers() error {
	c.Containers = map[string]*Container{}
	containers, err := c.c.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		return err
	}
	log.Infof("Found %d containers", len(containers))
	time.Sleep(2 * time.Second)
	for _, container := range containers {
		c.Containers[container.ID] = &Container{
			ID:     container.ID,
			Image:  container.Image,
			State:  container.State,
			Status: container.Status,
			Name:   container.Names[0],
			Labels: container.Labels,
		}
	}
	return nil
}

func (c *Client) Logs(container *Container) error {
	outputStream, err := container.Stream(c.URI)
	if err != nil {
		return err
	}
	errorStream, err := container.Stream(c.URI)
	if err != nil {
		return err
	}
	if err := c.c.Logs(docker.LogsOptions{
		Container:    container.ID,
		Stdout:       true,
		Stderr:       true,
		OutputStream: outputStream,
		ErrorStream:  errorStream,
		Follow:       true,
		Since:        time.Now().Unix(),
	}); err != nil {
		log.Errorf("unable to read logs from %s: %s", container.Name, err)
		return err
	}
	return nil
}

func (c *Client) ListenEvents() error {
	listener := make(chan *docker.APIEvents)
	if err := c.c.AddEventListener(listener); err != nil {
		return nil
	}

	for {
		select {
		case msg := <-listener:
			action := msg.Action
			typ := msg.Type
			id := msg.ID
			attributes := msg.Actor.Attributes
			log.Infof("%s %s %s\n", action, typ, id)
			if typ == "container" {
				switch action {
				case "stop":
					c.Containers.Remove(id)
				case "start":
					container := &Container{
						ID:     id,
						Image:  attributes["image"],
						Status: msg.Status,
						Name:   attributes["name"],
					}
					c.Containers[id] = container
					go c.Logs(container)
				}
			}
		}
	}
	return nil

}
