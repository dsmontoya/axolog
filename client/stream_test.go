package client

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMarshalJSON(t *testing.T) {
	Convey("Given a container data", t, func() {
		data := &Data{
			Message: "Hi, there!",
			Container: &Container{
				ID:    "abc",
				Image: "dsmontoya/axolog:latest",
				Labels: map[string]string{
					"label1": "abc",
				},
			},
		}

		Convey("Then the data is decoded to json", func() {
			b, err := json.Marshal(data)

			Convey("err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The json fields should be valid", func() {
				m := map[string]interface{}{}
				json.Unmarshal(b, &m)

				So(m["message"], ShouldEqual, data.Message)
				So(m["docker.id"], ShouldEqual, data.Container.ID)
				So(m["docker.image"], ShouldEqual, data.Container.Image)
			})
		})
	})
}
