package circular

import "testing"

func BenchmarkCircularPushPop(b *testing.B) {
	circular, _ := New()
	for i := 0; i < b.N; i++ {
		circular.PushBackOne(1)
		circular.PushBackOne(2)
		circular.PushBackOne(1)
		circular.PushBackOne(2)
		circular.PushBackOne(1)
		circular.PushBackOne(2)
		circular.PushBackOne(1)
		circular.PushBackOne(2)
		circular.PeekHead()
		circular.PeekTail()
		circular.PopBackOne()
		circular.PopBackOne()
		circular.PushBackOne(2)
		circular.PushBackOne(1)
		circular.PushBackOne(2)
		circular.PushBackOne(1)
		circular.PushBackOne(2)
		circular.PopBackOne()
		circular.PopBackOne()
	}
}
