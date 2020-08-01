package kafka

import (
	"context"
	"time"
	"github.com/segmentio/kafka-go/snappy"
	"github.com/segmentio/kafka-go"
)

func main(){

// to produce messages
topic := "first_topic"
partition := 0

conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

conn.SetWriteDeadline(time.Now().Add(10*time.Second))
conn.WriteMessages(
    kafka.Message{Value: string("one!")},
    kafka.Message{Value: []byte("two!")},
    kafka.Message{Value: []byte("three!")},
)

conn.Close()


}