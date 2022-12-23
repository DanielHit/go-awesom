package service

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

// convey
func TestUseSyncMap(t *testing.T) {
	UseSyncMap()

	Convey("should equals", t, func() {
		x := 2
		So(x, ShouldEqual, 2)
		So(x, ShouldBeGreaterThan, 1)
	})

	Convey("test random res", t, func() {
		x := rand.Int()
		So(x, ShouldEqual, 2)
		So(x, ShouldEqual, 2)
		So(x, ShouldEqual, 2)
	})

}
