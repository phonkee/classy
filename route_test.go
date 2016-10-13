package classy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRoute(t *testing.T) {

	Convey("Test New Route", t, func() {

		var tests = []struct {
			path    string
			name    string
		}{
			{
				"/", "routename",
			},
		}

		So(tests, ShouldNotBeNil)

		for _, item := range tests {
			route := NewViewRoute(item.path)
			So(route.GetName(), ShouldBeEmpty)

			route.Name(item.name)
			So(route.GetName(), ShouldEqual, item.name)
		}
	})

}
