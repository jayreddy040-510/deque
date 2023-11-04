package dll

import "testing"

func BenchmarkDllPushPop(b *testing.B) {
	dll := New()
	for i := 0; i < b.N; i++ {
		dll.PushBack(1)
		dll.PushBack(2)
		dll.PushBack(1)
		dll.PushBack(2)
		dll.PushBack(1)
		dll.PushBack(2)
		dll.PushBack(1)
		dll.PushBack(2)
		dll.PeekHead()
		dll.PeekTail()
		dll.PopBack(3)
		dll.PopBack(4)
		dll.PushBack(2)
		dll.PushBack(1)
		dll.PushBack(2)
		dll.PushBack(1)
		dll.PushBack(2)
		dll.PopBack(2)
		dll.PopBack()
	}
}
