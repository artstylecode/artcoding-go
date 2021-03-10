package mq

import (
	"fmt"
	"github.com/artstylecode/artcoding-go/utils"
	"github.com/streadway/amqp"
	"log"
)

var Channel = make(chan int)

type MqConfig struct {
	Host string
	Port int
	User string
	Pass string
}

type RabbitMQ struct {
	Connection    *amqp.Connection
	Channel       *amqp.Channel
	producerList  []Producer
	customerList  []Customer
	queueMap      map[string]int
	producerCount int
	queueName     string
}

//生产者接口
type Producer interface {
	GetQueueName() string
	MsgContent() string
}

//消费者接口
type Customer interface {
	ReceiverMsg(msg string) bool
}

func New(config MqConfig, queueName string) *RabbitMQ {
	var rabbit = RabbitMQ{
		queueName: queueName,
		queueMap:  make(map[string]int),
	}
	connectUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.User, config.Pass, config.Host, config.Port)

	connection, err := amqp.Dial(connectUrl)
	utils.FailOnError(err, "链接rabbitmq服务失败！")
	rabbit.Connection = connection
	ch, err := connection.Channel()
	utils.FailOnError(err, "创建通道失败")
	rabbit.Channel = ch
	return &rabbit
}

//Close 关闭rabbit相关连接
func (r *RabbitMQ) Close() {
	err := r.Channel.Close()
	utils.FailOnError(err, "关闭通道失败！")
	err = r.Connection.Close()
	utils.FailOnError(err, "关闭链接失败！")
}

//GetQueueMsgCount 获取指定队列数量
func (r *RabbitMQ) getQueueMsgCount() int {
	channel := r.Channel
	q, err := channel.QueueDeclare(
		r.queueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		utils.FailOnError(err, "获取消息数量错误！")
		return 0
	}
	return q.Messages
}

//设置主线程等待数量
func (r *RabbitMQ) setChannelCounts() {

	var sum = r.getQueueMsgCount()
	r.queueMap[r.queueName] = sum

	fmt.Printf("set counts..:%d\n queuename:%s", sum, r.queueName)
	fmt.Println(r.queueMap)
	Channel <- sum
}

//减去消息数量
func (r *RabbitMQ) reduceChannelCounts() {
	fmt.Printf("reduce counts..queuename:%s", r.queueName)

	r.queueMap[r.queueName]--
	r.setChannelCounts()
}

//RegisterProducer 注册生产者
func (r *RabbitMQ) RegisterProducer(producer Producer) {
	r.producerList = append(r.producerList, producer)
	r.queueMap[r.queueName]++
}

//RegisterCustomer 注册消费者
func (r *RabbitMQ) RegisterCustomer(customer Customer) {
	r.customerList = append(r.customerList, customer)
}

//消费消息
func (r *RabbitMQ) listenCustomer(customer Customer) {
	customerObject := customer
	channel := r.Channel
	q, err := channel.QueueDeclare(
		r.queueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "定义队列失败！")
	msgs, err := channel.Consume(q.Name, "",
		false,
		false,
		false,
		false,
		nil)
	utils.FailOnError(err, "消费队列失败！")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			res := customerObject.ReceiverMsg(string(d.Body))
			if res {
				d.Ack(false)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

//生产消息
func (r *RabbitMQ) listenProducer(producer Producer) {
	q, err := r.Channel.QueueDeclare(
		r.queueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "定义队列失败！")
	msg := producer.MsgContent()
	err = r.Channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", msg)
}
func (r RabbitMQ) wait() {
	<-Channel
}
func (r *RabbitMQ) Start() {
	//生产消息
	for _, producerItem := range r.producerList {
		go r.listenProducer(producerItem)

	}
	//消费消息
	for _, customerItem := range r.customerList {
		go r.listenCustomer(customerItem)
	}

	r.wait()

}
