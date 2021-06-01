package nutzamqp

import (
	"fmt"

	"github.com/rusriver/config"
	"github.com/streadway/amqp"
)

func AMQPBatchDeclare(
	channel *amqp.Channel,
	cfg *config.Config,
) {
	var err error

	for k, _ := range cfg.UMap("exchanges") {
		x, err := cfg.Get("exchanges." + k)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:18 / 1.1 at \"x, err := cfg.Get(\"exchanges.\" + k)\": %v", err))
		}
		err = channel.ExchangeDeclare(
			x.UString(k),           // name
			x.UString("type"),      // type
			x.UBool("durable"),     // durable
			x.UBool("autodeleted"), // auto-deleted
			false,                  // internal
			false,                  // noWait
			nil,                    // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:27 / 1.2: %v", err))
		}
	}
	for k, _ := range cfg.UMap("queues") {
		q, err := cfg.Get("queues." + k)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:32 / 2.1 at \"q, err := cfg.Get(\"queues.\" + k)\": %v", err))
		}
		_, err = channel.QueueDeclare(
			q.UString(k),                  // name of the queue
			q.UBool("durable"),            // durable
			q.UBool("delete_when_unused"), // delete when unused
			false,                         // exclusive
			false,                         // noWait
			nil,                           // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:40 / 2.2: %v", err))
		}
	}
	for _, v := range cfg.UMap("bindings") {
		b := &config.Config{Root: v}
		err = channel.QueueBind(
			b.UString("2"), // name of the queue
			b.UString("1"), // bindingKey
			b.UString("0"), // sourceExchange
			false,          // noWait
			nil,            // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:52 / 3.2: %v", err))
		}
	}
}
