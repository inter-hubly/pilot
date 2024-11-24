package broker

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/inter-hubly/pilot/testutils"
	"github.com/streadway/amqp"
)

const (
	withContainer     = true
	queueNun          = 10
	quantityOfMessage = 10
)

func TestQueue(t *testing.T) {
	for _, v := range []struct {
		testName string
		queues   func() []string
	}{
		{
			testName: "need send message to 20 queue",
			queues: func() []string {
				resp := make([]string, 0, queueNun*quantityOfMessage)
				for i := 0; i < queueNun; i++ {
					resp = append(resp, fmt.Sprintf("hello.%d", i))
				}
				return resp
			},
		},
	} {
		var wg sync.WaitGroup
		wg.Add(queueNun)

		ctx := context.Background()
		containerClose := startContainer(ctx, t)
		if containerClose != nil {
			defer containerClose(ctx)
		}

		received := make(chan string, queueNun)
		t.Run(v.testName, func(t *testing.T) {
			queues := v.queues()
			for _, q := range queues {
				go GetConnection().Consume(q, func(delivery amqp.Delivery) {
					received <- string(delivery.Body)
				})
				wg.Done()
			}
			wg.Wait()

			for _, q := range queues {
				GetConnection().QueueBind(NewQueueBinding(q, q, "test"))
			}

			for i := 0; i < quantityOfMessage; i++ {
				for _, q := range queues {
					GetConnection().Publish(q, []byte(fmt.Sprintf("%s-%d", q, i)))
				}
			}

			wg.Add(1)
			go func() {
				count := 1
				for msg := range received {
					count += 1
					log.Printf("%d %s\n", count, msg)
					if count == queueNun*quantityOfMessage {
						wg.Done()
						break
					}
				}

				close(received)
			}()
			wg.Wait()
		})

	}
}

func startContainer(ctx context.Context, t *testing.T) func(context.Context) error {
	if withContainer {
		host, close, err := testutils.RabbitMq(ctx)
		if err != nil {
			t.Fatal(err)
		}
		NewRabbitMQ("test", "topic", WithURL(host))
		return close
	}
	NewRabbitMQ("test", "topic", WithURL(rabbitMqDefault))
	return nil
}
