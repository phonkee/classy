package classy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMethod(t *testing.T) {

	var testdata = []struct {
		httpMethod Method
		methods    []Method
	}{
		{"GET", []Method{"Retrieve", "Other"}},
		{"POST", []Method{"Create", "Update"}},
		{"DELETE", []Method{"Destroy"}},
	}

	Convey("Test Method Map Contain", t, func() {
		for _, item := range testdata {
			mm := NewMethodMap()
			mm.Add(item.httpMethod, item.methods...)
			So(mm.Contains(item.httpMethod), ShouldBeTrue)
		}
	})

	Convey("Test Method Map Existing view method", t, func() {
		for _, item := range testdata {
			mm := NewMethodMap()
			mm.Add(item.httpMethod, item.methods...)

			methodset, ok := mm.GetMethodSet(item.httpMethod)
			So(ok, ShouldBeTrue)
			So(methodset.ContainsAll(item.methods...), ShouldBeTrue)
		}
	})

	Convey("Test Method Map Add/Delete", t, func() {
		for _, item := range testdata {
			mm := NewMethodMap()
			mm.Add(item.httpMethod)

			ms, ok := mm.GetMethodSet(item.httpMethod)
			So(ok, ShouldBeTrue)

			for _, m := range item.methods {

				// should not be found
				So(ms.Contains(m), ShouldBeFalse)

				// add item to set
				ms.Add(m)

				// now we have item
				So(ms.Contains(m), ShouldBeTrue)

				// remove
				ms.Remove(m)

				// not found
				So(ms.Contains(m), ShouldBeFalse)
			}
		}

		// Test methodmap delete
		for _, item := range testdata {
			mm := NewMethodMap()
			mm.Add(item.httpMethod)

			ms, ok := mm.GetMethodSet(item.httpMethod)
			So(ok, ShouldBeTrue)

			for _, m := range item.methods {

				// should not be found
				So(ms.Contains(m), ShouldBeFalse)

				// add item to set
				mm.Add(item.httpMethod, m)

				// now we have item
				So(ms.Contains(m), ShouldBeTrue)

				// remove
				mm.Delete(item.httpMethod, m)

				// not found
				So(ms.Contains(m), ShouldBeFalse)
			}
		}
	})
}
