/**

Filename: 		mutex.go
Author: 		gengming - gengming.zb@ccbft.com
Description:	loongrpc mutex tool logic
Create:			2022-07-02 09:00:05
Last Modified:	2022-07-02 11:10:27

*/

package utils

import (
	"sync"
	"sync/atomic"
)

var mutex sync.Mutex
var mutexMap = make(map[string]*innerMutex)

type innerMutex struct {
	count int64
	mutex sync.Mutex
}

func (m *innerMutex) Lock() {
	atomic.AddInt64(&m.count, 1)
	m.mutex.Lock()
}

func (m *innerMutex) Unlock() {
	m.mutex.Unlock()
	atomic.AddInt64(&m.count, -1)
}

func MutexLock(key string) {
	mutex.Lock()
	lock, ok := mutexMap[key]
	if !ok {
		lock = new(innerMutex)
		mutexMap[key] = lock
	}
	mutex.Unlock()
	lock.Lock()
}

func MutexUnlock(key string) {
	mutex.Lock()
	lock, ok := mutexMap[key]
	mutex.Unlock()
	if ok {
		lock.Unlock()
		if lock.count <= 0 {
			mutex.Lock()
			delete(mutexMap, key)
			mutex.Unlock()
		}
	}
}
