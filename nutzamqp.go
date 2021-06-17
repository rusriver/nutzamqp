package nutzamqp

import (
	"fmt"

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

	log.Fields["lib"] = "nutzamqp.AMQPBatchDeclare"

	for _, v := range cfg.UList("exchanges") {
		x := &config.Config{Root: v}

		log.Info("exchange declare: %v", x.UString("name"))

		err = channel.ExchangeDeclare(
			x.UString("name"),      // name
			x.UString("type"),      // type
			x.UBool("durable"),     // durable
			x.UBool("autodeleted"), // auto-deleted
			false,                  // internal
			false,                  // noWait
			nil,                    // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:34 / 1.2: %v", err))
		}

		log.Info("exchange declared OK: %v", x.UString("name"))
	}
	for _, v := range cfg.UList("queues") {
		q := &config.Config{Root: v}

		log.Info("queue declare: %v", q.UString("name"))

		_, err = channel.QueueDeclare(
			q.UString("name"),             // name of the queue
			q.UBool("durable"),            // durable
			q.UBool("delete_when_unused"), // delete when unused
			false,                         // exclusive
			false,                         // noWait
			nil,                           // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:52 / 2.2: %v", err))
		}

		log.Info("queue declared OK: %v", q.UString("name"))
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
			panic(fmt.Errorf("!!! nutzamqp.go:71 / 3.2: %v", err))
		}

		log.Info("binding declared OK (x--k--q): '%v' -- '%v' -- '%v'", x, k, q)
	}
	for _, v := range cfg.UList("xbindings") {
		b := &config.Config{Root: v}

		s := b.UString("0")
		k := b.UString("1")
		d := b.UString("2")

		err = channel.ExchangeBind(
			d,     // dest
			k,     // bindingKey
			s,     // source
			false, // noWait
			nil,   // arguments
		)
		if err != nil {
			panic(fmt.Errorf("!!! nutzamqp.go:90 / 4.2: %v", err))
		}

		log.Info("binding declared OK (x--k--q): '%v' -- '%v' -- '%v'", x, k, q)
	}
}
