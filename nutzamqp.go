package nutzamqp

import (
	"fmt"
	"regexp"

	"github.com/rusriver/config"
	"github.com/rusriver/filtertag"
	"github.com/streadway/amqp"
)

func AMQPBatchDeclare(
	log *filtertag.Entry,
	channel *amqp.Channel,
	cfg *config.Config,
) {
	var err error

	re_ddash_dot, err := regexp.Compile("::")
	if err != nil {
		panic(fmt.Errorf("!!! nutzamqp.go:19 / 0.1 at \"re_ddash_dot, err := regexp.Compile(\"::\")\": %v", err))
	}

	log.Fields["lib"] = "nutzamqp.AMQPBatchDeclare"

	for k, _ := range cfg.UMap("exchanges") {

		x, err := cfg.Get("exchanges." + k)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:26 / 1.1 at \"x, err := cfg.Get(\"exchanges.\" + k)\": %v", err))
		}

		k = re_ddash_dot.ReplaceAllString(k, ".")

		err = channel.ExchangeDeclare(
			k,                      // name
			x.UString("type"),      // type
			x.UBool("durable"),     // durable
			x.UBool("autodeleted"), // auto-deleted
			false,                  // internal
			false,                  // noWait
			nil,                    // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:38 / 1.2: %v", err))
		}

		log.Info("exchange declared OK: %v", k)
	}
	for k, _ := range cfg.UMap("queues") {

		q, err := cfg.Get("queues." + k)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:46 / 2.1 at \"q, err := cfg.Get(\"queues.\" + k)\": %v", err))
		}

		k = re_ddash_dot.ReplaceAllString(k, ".")

		_, err = channel.QueueDeclare(
			k,                             // name of the queue
			q.UBool("durable"),            // durable
			q.UBool("delete_when_unused"), // delete when unused
			false,                         // exclusive
			false,                         // noWait
			nil,                           // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:57 / 2.2: %v", err))
		}

		log.Info("queue declared OK: %v", k)
	}
	for _, v := range cfg.UList("bindings") {
		b := &config.Config{Root: v}

		x := b.UString("0")
		k := b.UString("1")
		q := b.UString("2")

		err = channel.QueueBind(
			q,     // name of the queue
			k,     // bindingKey
			x,     // sourceExchange
			false, // noWait
			nil,   // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:76 / 3.2: %v", err))
		}

		log.Info("binding declared OK (x--k--q): '%v' -- '%v' -- '%v'", x, k, q)
	}
}
