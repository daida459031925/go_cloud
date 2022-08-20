package goZeroServer

import (
	"github.com/daida459031925/common/error/try"
	"github.com/gogf/gf/v2/container/glist"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"service/common/connect/pulsar"
)

type server struct {
	s      *rest.Server         //go-zero service 服务
	pulsar pulsar.MessagePulsar //分布式消息队列pulsar
	g      *glist.List
}

func Setserver(s *rest.Server) *server {
	return &server{s: s, g: glist.New()}
}

func (s *server) setFunc(f func()) {
	if f == nil {
		return
	}
	s.g.PushBack(f)
}

func (s *server) Stop() {
	s.execute()
	s.s.Stop()
}

func (s *server) Start() {
	s.s.Start()
}

func (s *server) execute() {
	g := s.g
	if g != nil && g.Len() > 0 {
		for i := 0; i < g.Len(); i++ {
			f := g.PopBack()
			r, e := f.(func())
			if e {
				try.Try(r).CatchAll(func(err error) {
					logx.Errorf("server func error: %s", err.Error())
				})
			}
		}
	}
}

// 开启pulsar消息服务生产者
func (s *server) onPulsarProducer() {
	//pulsar.NewPulsarClient()
}

// 开启pulsar消息服务消费者
func (s *server) onPulsarConsumer() {

}

// 开启redis
func (s *server) onRedis() {

}
