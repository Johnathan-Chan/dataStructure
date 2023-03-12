package the_stack_containing_the_min_function

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCQueue(t *testing.T) {
	convey.Convey("the_stack_containing_the_min_function: minStack test case", t, func() {
		minStack := Constructor[int64, int64]()
		minStack.Push(1)
		minStack.Push(0)
		minStack.Push(-1)
		data := minStack.Top()
		convey.So(data, convey.ShouldEqual, -1)
	})
}
