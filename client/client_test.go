package client

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClient(t *testing.T) {
	Convey("Given a client", t, func() {
		client, err := New("udp://localhost:5000")
		So(err, ShouldBeNil)
		Convey("When the containers are registered", func() {
			err := client.RegisterContainers()

			Convey("err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("container list should not be empty", func() {
				So(client.Containers, ShouldNotBeEmpty)
			})

			Convey("The uri should be valid", func() {
				So(client.URI.Scheme, ShouldEqual, "udp")
				So(client.URI.Host, ShouldEqual, "localhost:5000")
			})
		})
	})
}
