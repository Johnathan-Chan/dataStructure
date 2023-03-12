package implement_queues_with_two_stacks

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCQueue(t *testing.T) {
	convey.Convey("implement_queues_with_two_stacks: CQueue test case", t, func() {
		cQueue := Constructor[int64, int64]()
		cQueue.AppendTail(1)
		cQueue.AppendTail(0)
		cQueue.AppendTail(-1)
		data := cQueue.DeleteHead()
		convey.So(data, convey.ShouldEqual, 1)
	})
}
