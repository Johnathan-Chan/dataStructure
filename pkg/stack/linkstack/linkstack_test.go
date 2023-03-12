package linkstack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLinkStack(t *testing.T) {
	convey.Convey("link stack: link stack test case", t, func() {
		linkStack := NewLinkStack[int64, int64]()
		linkStack.Push(1)
		linkStack.Push(2)
		linkStack.Push(3)
		convey.So(linkStack.Pop(), convey.ShouldEqual, 3)
		convey.So(linkStack.Pop(), convey.ShouldEqual, 2)
		convey.So(linkStack.Pop(), convey.ShouldEqual, 1)
	})
}
