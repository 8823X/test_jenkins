package z_lock

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

//var c = make(chan *sync.Mutex, 1000)

//var c = make(chan int)

//	func cLock() {
//		<-c
//	}
//
//	func cUnlock(l *sync.Mutex) {
//		c <- l
//	}
func CLock(k string) {
	g.Log().Info(gctx.New(), "====== 开始加锁 ====== %v", k)
	l := safeMap.GetOrSet(k, make(chan struct{}, 1))
	if v, ok := l.(chan struct{}); ok {
		time.Sleep(time.Second * 3)
		v <- struct{}{}
	} else {
		panic(ok)
	}
	g.Log().Info(gctx.New(), "====== 加锁 OK====== %v", k)
}

func CUnlock(k string) {
	g.Log().Info(gctx.New(), "====== 解锁 ====== %v", k)
	l := safeMap.Get(k)
	if v, ok := l.(chan struct{}); ok {
		time.Sleep(time.Second * 3)
		<-v
	} else {
		panic(ok)
	}
	g.Log().Info(gctx.New(), "====== 解锁 DEL ====== %v", k)
}
