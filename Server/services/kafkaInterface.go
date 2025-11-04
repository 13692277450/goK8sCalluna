package services

import (
	"context"
	"fmt"
	"log"
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

	topic := "kubectlSystemLogs" //journalctl -u kubelet --no-pager -n 10000
	partition := 0

	// 连接到 Kafka 的指定分区的 leader
	conn, err := kafka.DialLeader(context.Background(), "tcp", "192.168.1.211:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer conn.Close()
	// 设置写入超时时间
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	// 发送消息
	for i := range 100000 {
		go func() {
			_, err = conn.WriteMessages(
				kafka.Message{Key: []byte("Hellokey"), Value: []byte("Hello Kafka")},
				kafka.Message{Key: []byte("Anotherkey"), Value: []byte("Another Message")},
			)
			fmt.Printf("Message %v write.  \n", i)
		}()
	}

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	fmt.Println("Messages sent successfully")
}

// func WriterKafka(ctx context.Context) {
// 	// 强制使用IPv4
// 	dialer := &net.Dialer{
// 		Control: func(network, address string, c syscall.RawConn) error {
// 			// 强制使用IPv4
// 			if network == "tcp" {
// 				return nil
// 			}
// 			return nil
// 		},
// 	}

// 	writer := &kafka.Writer{
// 		Addr:                   kafka.TCP("192.168.1.211:9092"), // 确保使用正确的IP地址
// 		Topic:                  topic,
// 		Balancer:               &kafka.Hash{},
// 		WriteTimeout:           time.Second * 5,  // 增加超时时间
// 		RequiredAcks:           kafka.RequireOne, // 更可靠的确认
// 		AllowAutoTopicCreation: true,
// 		Transport: &kafka.Transport{
// 			Dial: dialer.DialContext,
// 		},
// 	}
// 	defer writer.Close()
// 	for i := 0; i < 10; i++ {
// 		strValue := fmt.Sprintf("hello world %s", strconv.Itoa(i))
// 		if err := writer.WriteMessages(
// 			ctx,
// 			kafka.Message{
// 				Key:   []byte("Key"),
// 				Value: []byte(strValue),
// 			},
// 		); err != nil {
// 			if err == kafka.LeaderNotAvailable {
// 				fmt.Printf("Leader not available: %v\n", err)
// 				time.Sleep(time.Second * 2)
// 				continue
// 			} else {
// 				fmt.Printf("Error writing to Kafka: %v\n", err)
// 				// 如果是Kafka错误，打印更多详细信息
// 				if kafkaErr, ok := err.(kafka.Error); ok {
// 					fmt.Printf("Kafka error: %s\n", kafkaErr.Error())
// 				}
// 			}
// 		} else {
// 			fmt.Printf("Message %v written to Kafka\n", i)

// 		}
// 	}
// 	//return &writer
// }

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
		Brokers:        []string{"192.168.1.211:9092"},
		Topic:          topic,
		GroupID:        "tesrec_team",
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
	defer reader.Close()

	// 读取所有消息，直到遇到错误或达到最大消息数
	maxMessages := 10 // 设置最大读取消息数，避免无限循环
	messageCount := 0
	for messageCount < maxMessages {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("Error reading from Kafka: %v\n", err)
			break
		}
		fmt.Printf("Message at %v: %s = %s\n", msg.Time, string(msg.Key), string(msg.Value))
		messageCount++
	}
	fmt.Printf("Total messages read: %d\n", messageCount)
}

func KafkaMain() {
	ctx := context.Background()
	WriterKafka(ctx)
	time.Sleep(time.Second * 3)
	go listenSignal()

	//readKafka(ctx)
}
