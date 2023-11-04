package dllpool

import "testing"

func BenchmarkDllPoolPushPop(b *testing.B) {
	dllPool := New()	
	for i := 0; i < b.N; i++ {
		dllPool.PushBack(1)
		dllPool.PushBack(2)
		dllPool.PushBack(1)
		dllPool.PushBack(2)
		dllPool.PushBack(1)
		dllPool.PushBack(2)
		dllPool.PushBack(1)
		dllPool.PushBack(2)
		dllPool.PeekHead()
		dllPool.PeekTail()
		dllPool.PopFront(3)
		dllPool.PopFront(4)
		dllPool.PushBack(2)
		dllPool.PushBack(1)
		dllPool.PushBack(2)
		dllPool.PushBack(1)
		dllPool.PushBack(2)
		dllPool.PopFront(2)
		dllPool.PopFront()
	}
}
