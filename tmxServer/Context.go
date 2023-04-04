// Package tmxServer /*
package tmxServer

import (
	"sync"
	"unsafe"
)

type Context struct {
	AllContext sync.Map
}

// GetContext
//
//	@Description: 获取上下文实例
//	@return *Context
func GetContext() *Context {
	return (*Context)(unsafe.Pointer(frameContainer.Get(new(Context).Key())))
}

func (c *Context) Key() string {
	return "context"
}

func (c *Context) Handle() *BaseFrame {
	return (*BaseFrame)(unsafe.Pointer(c))
}

// Set
//
//	@Description: 设置上下文
//	@receiver c
//	@param id
//	@param value
func (c *Context) Set(id string, value interface{}) {
	coId := GetGoId()
	var coroutineMap map[string]interface{}

	if v, ok := c.AllContext.Load(coId); ok {
		coroutineMap = v.(map[string]interface{})
	} else {
		coroutineMap = make(map[string]interface{})
	}

	coroutineMap[id] = value

	c.AllContext.Store(coId, coroutineMap)
}

// Get
//
//	@Description: 获取上下文
//	@receiver c
//	@param id
//	@return interface{}
//	@return bool
func (c *Context) Get(id string) (interface{}, bool) {
	coroutineMap, ok := c.AllContext.Load(GetGoId())

	if ok {
		m, ok := coroutineMap.(map[string]interface{})

		if ok != true {
			return nil, false
		}

		res, has := m[id]

		if has {
			return res, true
		}

		return nil, false
	}

	return nil, false
}

// Release
//
//	@Description: 释放上下文
//	@receiver c
func (c *Context) Release() {
	c.AllContext.Delete(GetGoId())
}
