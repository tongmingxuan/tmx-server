// Package tmxServer /*
package tmxServer

import (
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
	"runtime"
	"sync"
	"unsafe"
)

type RouteReg func(gin *gin.Engine)

type Listen struct {
	Name string
	Addr string
	Port string
}

type Server struct {
	goMaxProcess     int
	ServerConfigList []Listen
	ServerEntryList  sync.Map
	Router           map[string]RouteReg
}

func (server *Server) SetRouter(router map[string]RouteReg) {
	server.Router = router
}

func (server *Server) Key() string {
	return "server"
}

func (server *Server) Handle() *BaseFrame {

	frameConfig := (*Config)(unsafe.Pointer(frameContainer.Get(new(Config).Key())))
	serverConfig, ok := frameConfig.BaseConfigMap["server"]

	if ok != true {
		panic("server配置不存在")
	}

	maxP := frameConfig.GetConfigMap("go_max_process")

	var goMaxProcess int

	if php2go.Empty(maxP) {
		goMaxProcess = runtime.NumCPU()
	} else {
		v, ok := maxP.(string)

		if ok != true {
			panic("断言异常")
		}

		goMaxProcess = StringToInt(v, "go_max_process转化异常")
	}

	server.goMaxProcess = goMaxProcess

	runtime.GOMAXPROCS(goMaxProcess)

	server.ServerConfigList = make([]Listen, 0)

	for name, config := range serverConfig {
		func(c map[string]string, server *Server, name string) {
			addr, ok := c["listen_address"]
			if ok != true {
				panic("listen_address:不存在")
			}

			port, ok := c["listen_port"]
			if ok != true {
				panic("listen_port:不存在")
			}

			l := Listen{
				Name: name,
				Addr: addr,
				Port: port,
			}

			server.ServerConfigList = append(server.ServerConfigList, l)

			server.serverStart(l)
		}(config, server, name)
	}

	return (*BaseFrame)(unsafe.Pointer(server))
}

func (server *Server) serverStart(l Listen) {

	for key, function := range server.Router {
		func(key string, function RouteReg, l Listen) {
			if key == l.Name {
				ginServer := gin.Default()

				function(ginServer)

				server.ServerEntryList.Store(key, ginServer)

				addr := l.Addr + ":" + l.Port

				go func(ginServer *gin.Engine, addr string) {
					err := ginServer.Run(addr)
					if err != nil {
						panic("gin:err:" + err.Error())
					}
				}(ginServer, addr)

				server.ServerEntryList.Store(l.Name, ginServer)
			}
		}(key, function, l)
	}
}
