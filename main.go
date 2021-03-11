package main

import (
	"github.com/artstylecode/artcoding-go/mq"
)

func main() {
	//TODO 加入拼团队列
	mqConfig := mq.MqConfig{
		Host: "fish",
		Port: 5672,
		User: "guest",
		Pass: "5Zgeu7ByiMw3s3rL",
	}
	rabbitMq := mq.New(mqConfig, "group_beta")
	rabbitMq.SendMessage("{\"name\":\"123\", \"age\":123}")
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
