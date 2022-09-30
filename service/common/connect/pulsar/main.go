package pulsar

/**
分布式消息队列
*/
import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	cont "github.com/daida459031925/common/context"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"os"
	"service/common/constant"
	"strings"
	"time"
)

type MessagePulsar struct {
	pulsar.Client
	exit func()
}

var exit = func() { os.Exit(3) }

// NewPulsarClient Url 单个路径:端口就是单地址 路径:端口，路径:端口，路径:端口 就是集群
// 创建连接器
func NewPulsarClient(addr []net.TCPAddr, cli ...pulsar.ClientOptions) MessagePulsar {
	messagePulsar := MessagePulsar{nil, exit}

	opt := pulsar.ClientOptions{}

	if cli != nil && len(cli) > 0 {
		opt = cli[0]
	}

	if addr != nil && len(addr) > 0 {
		addrs := make([]string, 0)
		for i := range addr {
			addrs = append(addrs, addr[i].String())
		}

		opt.URL = fmt.Sprintf(constant.UseClientPulsar01s_01, strings.Join(addrs, constant.SysComma))
	} else {
		logx.Errorf(constant.ErrClientPulsar00_01)
		exit()
	}

	client, err := pulsar.NewClient(opt)

	if err != nil {
		logx.Errorf(constant.ErrClientPulsar01s_01, err)
		messagePulsar.exit()
	}

	messagePulsar.Client = client

	return messagePulsar
}

// NewPulsarClientString Url 单个路径:端口就是单地址 路径:端口，路径:端口，路径:端口 就是集群
// 创建连接器
func NewPulsarClientString(addr string, cli ...pulsar.ClientOptions) MessagePulsar {
	messagePulsar := MessagePulsar{nil, exit}

	opt := pulsar.ClientOptions{}

	if cli != nil && len(cli) > 0 {
		opt = cli[0]
	}

	if len(addr) > 0 {
		opt.URL = addr
	} else {
		logx.Error(constant.ErrClientPulsar00_01)
		exit()
	}

	client, err := pulsar.NewClient(opt)

	if err != nil {
		logx.Errorf(constant.ErrClientPulsar01s_01, err)
		messagePulsar.exit()
	}

	messagePulsar.Client = client

	return messagePulsar
}

// CreateProducer 创建生产者
func (m MessagePulsar) NewProducer(topic string, pro ...pulsar.ProducerOptions) pulsar.Producer {
	opt := pulsar.ProducerOptions{}

	if pro != nil && len(pro) > 0 {
		opt = pro[0]
	}

	opt.Topic = topic
	//消息压缩（四种压缩方式：LZ4，ZLIB，ZSTD，SNAPPY），consumer端不用做改动就能消费，开启后大约可以降低3/4带宽消耗和存储（官方测试）
	opt.CompressionType = pulsar.LZ4

	producer, err := m.Client.CreateProducer(opt)

	if err != nil {
		logx.Errorf(constant.ErrPulsarProducer01s_01, err)
		m.exit()
	}

	return producer

}

// CreateSubscribe 创建消费者
func (m MessagePulsar) NewSubscribe(topic string, subscriptionName string, t pulsar.SubscriptionType, channel chan pulsar.ConsumerMessage) pulsar.Consumer {
	return m.NewSubscribeOptions(topic, subscriptionName, t, channel)
}

// CreateSubscribeOptions 创建消费者
func (m MessagePulsar) NewSubscribeOptions(topic string, subscriptionName string, t pulsar.SubscriptionType, channel chan pulsar.ConsumerMessage, opts ...pulsar.ConsumerOptions) pulsar.Consumer {
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
		logx.Errorf(constant.ErrPulsarSubscribe01s_01, err)
		m.exit()
	}

	return consumer
}

// GetConsumerMsg 通道缓冲个数
func GetConsumerMsg(buffers ...int) chan pulsar.ConsumerMessage {
	def := 100
	if buffers != nil && len(buffers) > 0 && buffers[0] > 0 {
		def = buffers[0]
	}
	return make(chan pulsar.ConsumerMessage, def)
}

// GetContext 执行超时，取消执行
func GetContext(i ...time.Duration) func() (context.Context, context.CancelFunc) {
	return cont.GetContext(i...)
}

// SendProducerMsg 发送消息，发送完毕，关闭内容,同步
func SendProducerMsg(producer pulsar.Producer, con func() (context.Context, context.CancelFunc), msg *pulsar.ProducerMessage) (pulsar.MessageID, error) {
	ctx, cancle := con()
	defer cancle()
	msgID, err := producer.Send(ctx, msg)
	defer producer.Close()
	return msgID, err
}

// SendProducerAsyncMsg 发送消息，发送完毕，关闭内容,异步
func SendProducerAsyncMsg(producer pulsar.Producer, con func() (context.Context, context.CancelFunc), msg *pulsar.ProducerMessage,
	errFunc func(id pulsar.MessageID, message *pulsar.ProducerMessage, err error)) {
	ctx, cancle := con()
	defer cancle()
	producer.SendAsync(ctx, msg, errFunc)
	defer producer.Close()
}

// CreateListenerMsg 监听消费
func NewListenerMsg(consumer pulsar.Consumer, channel chan pulsar.ConsumerMessage, f func(consumerMsg pulsar.ConsumerMessage, consumer pulsar.Consumer)) {

	// Receive messages from channel. The channel returns a struct which contains message and the consumer from where
	// the message was received. It's not necessary here since we have 1 single consumer, but the channel could be
	// shared across multiple consumers as well
	for cm := range channel {
		msg := cm.Message
		logx.Info(fmt.Sprintf(constant.UsePulsarConsumerMessage02vs_01, msg.ID(), msg.Payload()))
		f(cm, consumer)
	}

}
