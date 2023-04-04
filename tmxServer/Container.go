// Package tmxServer /*
package tmxServer

var frameContainer *Container

func CreateContainer() *Container {
	return new(Container)
}

type Container struct {
	Kv map[string]*BaseFrame
}

func (container *Container) Set(key string, v *BaseFrame) {
	if container.Kv == nil {
		container.Kv = make(map[string]*BaseFrame, 10)
	}

	container.Kv[key] = v
}

func (container *Container) Get(key string) *BaseFrame {

	v, ok := container.Kv[key]

	if ok != true {
		return nil
	}

	return v
}
