package goZeroServer

import (
	"github.com/zeromicro/go-zero/rest"
	"service/common/algorithm/queue/bqueue"
	"service/common/connect/pulsar"
)

type server struct {
	*rest.Server                        //go-zero service 服务
	Pulsar         pulsar.MessagePulsar //分布式消息队列pulsar
	*bqueue.BQueue                      //自己配合gf2实现的队列
}

func NewServer(s *rest.Server) *server {
	return &server{Server: s, BQueue: bqueue.NewFuncList()}
}

func (s *server) Stop() {
	s.BQueue.ExecutePopBack()
	s.Server.Stop()
}

func (s *server) Start() {
	s.Server.Start()
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
