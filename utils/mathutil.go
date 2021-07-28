package utils



/*
数学工具

这里主要封装了一个获取指定范围内随机数的方法

考虑到任意两次调用间时间种子不得相同，我们使用了强制同步和阻塞睡眠
 */
import (
"time"
"math/rand"
"sync"
)

var(
	//随机数互斥锁（确保GetRandomInt不能被并发访问）
	randomMutex sync.Mutex
)

/*获取[start,end]范围内的随机数*/
func GetRandomInt(start, end int) int {
	randomMutex.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := start + r.Intn(end-start+1)
	randomMutex.Unlock()
	return n
}

