package main

import (
	"context"

	"github.com/mvrdgs/go-experts/rabbitmq/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	ctx := context.Background()

	rabbitmq.Publish(ctx, ch, "amq.direct", "Hello World")
}
