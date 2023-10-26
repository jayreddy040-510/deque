package dll

import "testing"

func BenchmarkDllPushPop(b *testing.B) {
	dll := New()
	for i := 0; i < b.N; i++ {
		dll.PushRight(1)
		dll.PushRight(2)
		dll.PushRight(1)
		dll.PushRight(2)
		dll.PushRight(1)
		dll.PushRight(2)
		dll.PushRight(1)
		dll.PushRight(2)
		dll.PeekHead()
		dll.PeekTail()
		dll.PopLeft(3)
		dll.PopLeft(4)
		dll.PushRight(2)
		dll.PushRight(1)
		dll.PushRight(2)
		dll.PushRight(1)
		dll.PushRight(2)
		dll.PopLeft(2)
		dll.PopLeft()
	}
}
