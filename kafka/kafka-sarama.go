package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var brokers = []string{"127.0.0.1:9092"}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}

func main() {

	producer, err := newProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	msg := prepareMessage("first_topic", "Sarama")
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("%s error occured.", err.Error())
	} else {
		fmt.Printf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
	}

	consumer, err := sarama.NewConsumer(brokers, nil)

	partitionList, err := consumer.Partitions("first_topic") //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetOldest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition("first_topic", partition, initialOffset)

		for message := range pc.Messages() {
			fmt.Println(string(message.Value))
		}

	}

}
