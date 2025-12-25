package chsync

import "sync"

// this works also but bad idea because embedding types are public so your lock is exposed
// -> no encapsulation on mutex
//type Counter2 struct {
//	sync.Mutex
//	value int
//}
//
//func (c *Counter2) Inc2() {
//	c.Lock()
//	defer c.Unlock()
//	c.value++
//}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
