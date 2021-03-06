package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "testTopic"
	// kafkaConn = "localhost:2202"
	// topic     = "testTopic"
	// kafkaConn = "b-1.tlio-dev-sg-msk-c.3its0c.c4.kafka.ap-southeast-1.amazonaws.com:9092"
	// topic     = "testTopic"
)

func main() {
	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	// read command line input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		// publish without goroutene
		publish(msg, producer)

		// publish with go routene
		// go publish(msg, producer)
	}
}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Net.TLS.Enable = false
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	//brokers := []string{"b-1.tlio-dev-sg-msk-c.3its0c.c4.kafka.ap-southeast-1.amazonaws.com:9092", "b-2.tlio-dev-sg-msk-c.3its0c.c4.kafka.ap-southeast-1.amazonaws.com:9092"}
	brokers := []string{"localhost:9092"}
	kafkaEventClient, _ := sarama.NewClient(brokers, config)
	prd, err := sarama.NewSyncProducerFromClient(kafkaEventClient)
	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	//	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}

func publish(message string, producer sarama.SyncProducer) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
