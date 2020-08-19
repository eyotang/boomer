package boomer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {
	Convey("test Sort no change", t, func() {
		var (
			t1 = &Task{Order: 1, Name: "t1"}
			t2 = &Task{Order: 2, Name: "t2"}
			t3 = &Task{Order: 3, Name: "t3"}
		)

		expTs := []*Task{t1, t2, t3}
		ts := []*Task{t1, t2, t3}
		Sort(ts)
		So(ts, ShouldResemble, expTs)
	})

	Convey("test Sort reverse", t, func() {
		var (
			t1 = &Task{Order: 1, Name: "t1"}
			t2 = &Task{Order: 2, Name: "t2"}
			t3 = &Task{Order: 3, Name: "t3"}
		)

		expTs := []*Task{t1, t2, t3}
		ts := []*Task{t3, t2, t1}
		Sort(ts)
		So(ts, ShouldResemble, expTs)
	})

	Convey("test Sort change", t, func() {
		var (
			t1 = &Task{Order: 1, Name: "t1"}
			t2 = &Task{Order: 2, Name: "t2"}
			t3 = &Task{Order: 3, Name: "t3"}
		)

		expTs := []*Task{t1, t2, t3}
		ts := []*Task{t2, t3, t1}
		Sort(ts)
		So(ts, ShouldResemble, expTs)
	})

	Convey("test Sort minus", t, func() {
		var (
			t1 = &Task{Order: -1, Name: "t1"}
			t2 = &Task{Order: 2, Name: "t2"}
			t3 = &Task{Order: 3, Name: "t3"}
		)

		expTs := []*Task{t1, t2, t3}
		ts := []*Task{t2, t3, t1}
		Sort(ts)
		So(ts, ShouldResemble, expTs)
	})
}
