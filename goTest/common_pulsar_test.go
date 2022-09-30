package goTest

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	IP "github.com/daida459031925/common/ip"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

type messagePulsar struct {
	pulsar.Client
	exit func()
}

var exit = func() { os.Exit(3) }

// Url 单个路径:端口就是单地址 路径:端口，路径:端口，路径:端口 就是集群
// 创建连接器
func NewPulsarClient(addr []net.TCPAddr, cli ...pulsar.ClientOptions) messagePulsar {
	messagePulsar := messagePulsar{nil, exit}

	opt := pulsar.ClientOptions{}

	if cli != nil && len(cli) > 0 {
		opt = cli[0]
	}

	if addr != nil && len(addr) > 0 {
		addrs := make([]string, 0)
		for i := range addr {
			addrs = append(addrs, addr[i].String())
		}

		opt.URL = fmt.Sprintf("pulsar://%s", strings.Join(addrs, ","))
	} else {
		exit()
	}

	client, err := pulsar.NewClient(opt)

	if err != nil {
		logx.Errorf("pulsar client error: %s", err)
		messagePulsar.exit()
	}

	messagePulsar.Client = client

	return messagePulsar
}

// 创建生产者
func (m messagePulsar) NewProducer(topic string, pro ...pulsar.ProducerOptions) pulsar.Producer {
	opt := pulsar.ProducerOptions{}

	if pro != nil && len(pro) > 0 {
		opt = pro[0]
	}

	opt.Topic = topic
	//消息压缩（四种压缩方式：LZ4，ZLIB，ZSTD，SNAPPY），consumer端不用做改动就能消费，开启后大约可以降低3/4带宽消耗和存储（官方测试）
	opt.CompressionType = pulsar.LZ4

	producer, err := m.Client.CreateProducer(opt)

	if err != nil {
		logx.Errorf("pulsar producer error: %s", err)
		m.exit()
	}

	return producer

}

// 创建消费者
func (m messagePulsar) NewSubscribe(topic string, subscriptionName string, t pulsar.SubscriptionType, channel chan pulsar.ConsumerMessage) pulsar.Consumer {
	return m.NewSubscribeOptions(topic, subscriptionName, t, channel)
}

// 创建消费者
func (m messagePulsar) NewSubscribeOptions(topic string, subscriptionName string, t pulsar.SubscriptionType, channel chan pulsar.ConsumerMessage, opts ...pulsar.ConsumerOptions) pulsar.Consumer {
	opt := pulsar.ConsumerOptions{}
	if opts != nil && len(opts) > 0 {
		opt = opts[0]
	}

	opt.MessageChannel = channel
	opt.Topic = topic
	opt.SubscriptionName = subscriptionName
	opt.Type = t

	consumer, err := m.Subscribe(opt)
	if err != nil {
		logx.Errorf("pulsar Subscribe error: %s", err)
		m.exit()
	}

	return consumer
}

// 通道缓冲个数
func GetConsumerMsg(buffers ...int) chan pulsar.ConsumerMessage {
	def := 100
	if buffers != nil && len(buffers) > 0 && buffers[0] > 0 {
		def = buffers[0]
	}
	return make(chan pulsar.ConsumerMessage, def)
}

// 执行超时，取消执行
func GetContext(i ...time.Duration) func() (context.Context, context.CancelFunc) {
	return func() (context.Context, context.CancelFunc) {
		var t time.Duration = 5
		if i != nil && len(i) > 0 && i[0] > 0 {
			t = i[0]
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*t)
		return ctx, cancel
	}
}

// 发送消息，发送完毕，关闭内容,同步
func SendProducerMsg(producer pulsar.Producer, con func() (context.Context, context.CancelFunc), msg *pulsar.ProducerMessage) (pulsar.MessageID, error) {
	ctx, cancle := con()
	defer cancle()
	msgID, err := producer.Send(ctx, msg)
	defer producer.Close()
	return msgID, err
}

// 发送消息，发送完毕，关闭内容,异步
func SendProducerAsyncMsg(producer pulsar.Producer, con func() (context.Context, context.CancelFunc), msg *pulsar.ProducerMessage,
	errFunc func(id pulsar.MessageID, message *pulsar.ProducerMessage, err error)) {
	ctx, cancle := con()
	defer cancle()
	producer.SendAsync(ctx, msg, errFunc)
	defer producer.Close()
}

// 监听消费
func NewListenerMsg(consumer pulsar.Consumer, channel chan pulsar.ConsumerMessage, f func(consumerMsg pulsar.ConsumerMessage, consumer pulsar.Consumer)) {

	// Receive messages from channel. The channel returns a struct which contains message and the consumer from where
	// the message was received. It's not necessary here since we have 1 single consumer, but the channel could be
	// shared across multiple consumers as well
	for cm := range channel {
		msg := cm.Message
		logx.Infof("Received message  msgId: %v -- content: '%s'\n", msg.ID(), string(msg.Payload()))
		f(cm, consumer)
	}

}

// 创建IP地址
func DistinctIp(addr []net.TCPAddr) []net.TCPAddr {
	if addr == nil || len(addr) == 0 {
		exit()
	}

	ips := make([]net.TCPAddr, 0)

	fx.From(func(source chan<- interface{}) {
		for i := range addr {
			source <- addr[i]
		}
	}).Distinct(func(item interface{}) interface{} {
		var a = item.(net.TCPAddr)
		return fmt.Sprintf("%s:%d", a.IP, a.Port)
	}).ForEach(func(item interface{}) {
		var a = item.(net.TCPAddr)
		ips = append(ips, a)
	})

	return ips
}

func TestProducerMsg(t *testing.T) {

	addr := make([]net.TCPAddr, 0)
	addr = append(addr, IP.NewIp(192, 168, 0, 100, 6550))
	addr = append(addr, IP.NewIp(192, 168, 0, 100, 6550))
	DistinctIp(addr)

	//client := NewPulsarClient(ip,port)
	//defer client.Close()
	//
	//producer := client.CreateProducer("my-topic")
	//defer producer.Close()
	//
	//msgId, e := SendProducerMsg(producer, GetContext(), &pulsar.ProducerMessage{
	//	Payload: []byte("hello"),
	//})
	//
	//if e != nil {
	//	logx.Errorf("发送失败")
	//	return
	//}
	//
	//logx.Infof("发送成功 , msgId为: %s", msgId)

}

func TestListenerMsg(t *testing.T) {

	//client := NewPulsarClient()
	//defer client.Close()
	//
	//channel := GetConsumerMsg()
	//options := pulsar.ConsumerOptions{
	//	Topic:            "my-topic",
	//	SubscriptionName: "my-subscription",
	//	Type:             pulsar.Shared,
	//}
	//
	//consumer := client.CreateSubscribeOptions(options, channel)
	//defer consumer.Close()
	//
	//CreateListenerMsg(consumer, channel, func(consumerMsg pulsar.ConsumerMessage, consumer pulsar.Consumer) {
	//	msg := consumerMsg.Message
	//	fmt.Printf("Received message  msgId: %v -- content: '%s'\n",
	//		msg.ID(), string(msg.Payload()))
	//	consumer.Ack(msg)
	//})

}
