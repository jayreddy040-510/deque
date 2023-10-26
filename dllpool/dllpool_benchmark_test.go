package dllpool

import "testing"

func BenchmarkDllPoolPushPop(b *testing.B) {
	dllPool := New()	
	for i := 0; i < b.N; i++ {
		dllPool.PushRight(1)
		dllPool.PushRight(2)
		dllPool.PushRight(1)
		dllPool.PushRight(2)
		dllPool.PushRight(1)
		dllPool.PushRight(2)
		dllPool.PushRight(1)
		dllPool.PushRight(2)
		dllPool.PeekHead()
		dllPool.PeekTail()
		dllPool.PopLeft(3)
		dllPool.PopLeft(4)
		dllPool.PushRight(2)
		dllPool.PushRight(1)
		dllPool.PushRight(2)
		dllPool.PushRight(1)
		dllPool.PushRight(2)
		dllPool.PopLeft(2)
		dllPool.PopLeft()
	}
}