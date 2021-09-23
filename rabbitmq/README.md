# rabbitmq broker driver

> 基于 [asim/go-micro](https://github.com/asim/go-micro)框架中的[rabbitmq](https://github.com/asim/go-micro/tree/master/plugins/broker/rabbitmq) 组件修改

> 用以解决 micro.NewEvent 不能设置rabbitmq持久化消息的配置
    
   - rabbitmq.DeliveryMode(amqp.Persistent)
   
> 新增及修改

- 新增: 延时队列
- 修改: 
    - 投递模式: DeliveryMode(设置为2时,rabbitMQ即时挂了,重启后也能恢复数据)
    - 消息等级: Priority