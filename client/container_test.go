package client

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemove(t *testing.T) {
	Convey("Given a container map", t, func() {
		containers := Containers{
			"abc": &Container{},
		}
		Convey("When an item is removed", func() {
			containers.Remove("abc")

			Convey("The item should not exist anymore", func() {
				So(containers["abc"], ShouldBeNil)
			})
		})
	})
}
