package services

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	reader *kafka.Reader
	//writer *kafka.Writer
	topic = "user_click"
)

func WriterKafka(ctx context.Context) {
	fmt.Println("Kafka Writer started....")
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()
	for i := 0; i < 10; i++ {
		if err := writer.WriteMessages(
			ctx,
			kafka.Message{
				Key:   []byte("Key"),
				Value: []byte("Value"),
			},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				fmt.Println("Leader not available")
				time.Sleep(time.Second * 2)
				continue
			} else {
				fmt.Println("Error writing to Kafka")

			}
		} else {
			fmt.Println("Message written to Kafka")
			break
		}
	}
	//return &writer
}

func listenSignal() {
	fmt.Println("Kafka listenSignal started....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Println("Received message: ", sig.String())
	if reader != nil {
		reader.Close()
	}
	os.Exit(0)
}

func readKafka(ctx context.Context) {
	fmt.Println("Kafka reader started....")

	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          topic,
		GroupID:        "tesrec_team",
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
	defer reader.Close()
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading from Kafka")
			break
		}
		fmt.Printf("Message at %v: %s = %s\n", msg.Time, string(msg.Key), string(msg.Value))
	}
}

func main1() {
	ctx := context.Background()
	WriterKafka(ctx)
	time.Sleep(time.Second * 3)
	go listenSignal()
	time.Sleep(time.Second * 3)

	readKafka(ctx)
}
