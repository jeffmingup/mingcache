package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}
type Group struct {
	sync.Mutex
	m map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok { //存在正执行的查询
		g.Unlock()
		c.wg.Wait() //等到查询结束
		return c.val, c.err
	}
	c := new(call) //没有正在执行的查询
	c.wg.Add(1)
	g.m[key] = c
	g.Unlock()

	c.val, c.err = fn()
	c.wg.Done() //查询结束，减锁

	g.Lock()
	delete(g.m, key) //更新正在执行的map记录
	g.Unlock()

	return c.val, c.err
}
