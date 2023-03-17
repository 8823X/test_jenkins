package z_lock

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"sync"
	"sync/atomic"
)

//订单锁模块，保证订单在同一时间只能有一个正在进行中的操作流程

var safeMap = gmap.NewHashMap(true) //打开并发安全

type lock struct {
	Count int64
	Lock  *sync.Mutex
}

var delLock sync.RWMutex

//	func TryLock(k string) bool {
//		l := safeMap.GetOrSet(k, &lock{Count: 0, Lock: &sync.Mutex{}})
//		if v, ok := l.(*lock); ok {
//			for {
//				if atomic.CompareAndSwapInt64(&v.Sign, 0, 0) {
//					atomic.AddInt64(&v.Count, 1) //原子++
//					if v.Lock.TryLock() {
//						return true
//					} else {
//						atomic.AddInt64(&v.Count, -1) //原子--
//						return false
//					}
//				}
//			}
//		} else {
//			panic(ok)
//		}
//	}
func Lock(k string) {
	//查询是否存在订单对应的锁
	delLock.RLock()
	defer delLock.RUnlock()
	l := safeMap.GetOrSet(k, &lock{Count: 0, Lock: &sync.Mutex{}})
	if v, ok := l.(*lock); ok {
		//g.Log().Info(gctx.New(), fmt.Sprintf("====== 开始加锁 !, count : %v", v.Count))
		atomic.AddInt64(&v.Count, 1) //原子++
		g.Log().Info(gctx.New(), fmt.Sprintf("====== 等待获锁 !, count : %v", v.Count))
		v.Lock.Lock()
		g.Log().Info(gctx.New(), fmt.Sprintf("====== 加锁完毕 !, count : %v", v.Count))
	} else {
		panic(ok)
	}
}

func Unlock(id string) {
	//解锁并从map中移除
	if l := safeMap.Get(id); l != nil {
		if v, ok := l.(*lock); ok {
			atomic.AddInt64(&v.Count, -1) //原子--
			v.Lock.Unlock()
			//g.Log().Info(gctx.New(), fmt.Sprintf("====== 解锁完毕 !, count : %v", v.Count))
			//判断是否可以删除该锁
			g.Log().Info(gctx.New(), fmt.Sprintf("====== 判断是否可删除锁 !, count : %v", v.Count))
			if atomic.CompareAndSwapInt64(&v.Count, 0, 0) {
				delLock.Lock()
				//g.Log().Info(gctx.New(), fmt.Sprintf("====== 进入删除等待5秒！！！！！！！！ !, count : %v", v.Count))
				//time.Sleep(time.Second * 5)
				safeMap.Remove(id)
				g.Log().Info(gctx.New(), fmt.Sprintf("====== 删除对应锁！！！！！！！！ !, count : %v", v.Count))
				delLock.Unlock()
			}
		} else {
			panic(ok)
		}
	} else {
		panic(fmt.Sprintf("unloc error there is no lock : %v", id))
	}
}
