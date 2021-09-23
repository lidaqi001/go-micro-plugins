package rabbitmq_test

import (
	"codeup.aliyun.com/shan/micro.protos/basic/event/event_hello"
	"context"
	"fmt"
	"github.com/asim/go-micro/v3/client"
	rabbitmq "github.com/lidaqi001/go-micro-plugins/rabbitmq"
	"github.com/streadway/amqp"
	"os"
	"testing"
	"time"

	micro "github.com/asim/go-micro/v3"
	broker "github.com/asim/go-micro/v3/broker"
	server "github.com/asim/go-micro/v3/server"
)

type Example struct{}

func (e *Example) Handler(ctx context.Context, r interface{}) error {
	return nil
}

func TestDurable(t *testing.T) {
	if tr := os.Getenv("TRAVIS"); len(tr) > 0 {
		t.Skip()
	}
	rabbitmq.DefaultRabbitURL = "amqp://rabbitmq:rabbitmq@127.0.0.1:5672"
	brkrSub := broker.NewSubscribeOptions(
		broker.Queue("queue.default"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)

	b := rabbitmq.NewBroker()
	b.Init()
	if err := b.Connect(); err != nil {
		t.Logf("cant conect to broker, skip: %v", err)
		t.Skip()
	}

	s := server.NewServer(server.Broker(b))

	service := micro.NewService(
		micro.Server(s),
		micro.Broker(b),
	)
	h := &Example{}
	// Register a subscriber
	micro.RegisterSubscriber(
		"topic",
		service.Server(),
		h.Handler,
		server.SubscriberContext(brkrSub.Context),
		server.SubscriberQueue("queue.default"),
	)

	//service.Init()

	if err := service.Run(); err != nil {
		t.Fatal(err)
	}

}

// 测试延时消息发布
func TestDelayPublish(t *testing.T) {
	var (
		err      error
		ctx      = context.Background()
		topic    = "test"
		delaySec = 20
		event    = &event_hello.HelloEvent{
			Who: "我是一条延时消息,投递时间:" + time.Now().Format("2006-01-02 15:04:05"),
		}
	)
	service := getService(t)
	publishOptsCtx := context.WithValue(ctx, rabbitmq.DeliveryMode{}, amqp.Persistent)
	publishOptsCtx = context.WithValue(publishOptsCtx, rabbitmq.Expire{}, delaySec*1000)
	p := micro.NewEvent(topic, service.Client())
	err = p.Publish(ctx, event, client.PublishContext(publishOptsCtx))
	if err != nil {
		t.Logf("[pub] failed: %v", err)
		t.FailNow()
	}
	t.Log("success")
}

// 测试延时消息订阅
type DelayExample struct{}

func (d *DelayExample) Handler(ctx context.Context, event *event_hello.HelloEvent) error {
	fmt.Println("消费时间:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("消费延时消息:", event.Who)
	return nil
}
func TestDelaySubscribe(t *testing.T) {
	var (
		queue = "test"
		topic = "test"
	)
	service := getService(t)
	brkrSub := broker.NewSubscribeOptions(
		broker.Queue(queue),
		rabbitmq.DurableQueue(),
		rabbitmq.IsDelayQueue(),
	)
	h := &DelayExample{}
	// Register a subscriber
	_ = micro.RegisterSubscriber(
		topic,
		service.Server(),
		h.Handler,
		server.SubscriberContext(brkrSub.Context),
		server.SubscriberQueue(queue),
	)

	if err := service.Run(); err != nil {
		t.Fatal(err)
	}
}

func getService(t *testing.T) micro.Service {

	rabbitmq.DefaultRabbitURL = "amqp://sxx:123456@rabbit-apples-rabbitmq.shan.svc.cluster.local:5672"
	b := rabbitmq.NewBroker()
	b.Init()
	if err := b.Connect(); err != nil {
		t.Logf("cant conect to broker, skip: %v", err)
		t.Skip()
	}

	s := server.NewServer(server.Broker(b))

	service := micro.NewService(
		micro.Server(s),
		micro.Broker(b),
	)
	return service
}
