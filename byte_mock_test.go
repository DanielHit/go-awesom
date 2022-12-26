package service

import (
	"fmt"
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

func Foo(in string) string {
	return in
}

type A struct{}

func (a A) Foo(in string) string { return in }

var Bar = 0

func TestMockXXX(t *testing.T) {
	PatchConvey("TestMockXXX", t, func() {
		Mock(Foo).Return("c").Build()   // mock function
		Mock(A.Foo).Return("c").Build() // mock method
		MockValue(&Bar).To(1)           // mock variable

		So(Foo("a"), ShouldEqual, "c")        // assert `Foo` is mocked
		So(new(A).Foo("b"), ShouldEqual, "c") // assert `A.Foo` is mocked
		So(Bar, ShouldEqual, 1)               // assert `Bar` is mocked
	})
	// mock is released automatically outside `PatchConvey`
	fmt.Println(Foo("a"))        // a
	fmt.Println(new(A).Foo("b")) // b
	fmt.Println(Bar)             // 0
}
