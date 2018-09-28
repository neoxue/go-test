package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"fmt"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	kazoo "github.com/wvanbergen/kazoo-go"
)

/*
https://github.com/Shopify/sarama
支持 kafka 0.8 + 版本
现在支持 0.9 以上，0.8 有bug需要自己解决
sarama 是底层库，确切封装0.8以前及以后分别有2个库
https://godoc.org/github.com/Shopify/sarama

https://github.com/wvanbergen/kafka/blob/master/examples/consumergroup/main.go
*/

const (
	DefaultConsumerGroup = "test"
	Defaultzk            = "localhost:2181"
	DefaultKafkaTopics   = "test"
)

var (
	consumerGroup  = flag.String("group", DefaultConsumerGroup, "The name of the consumer group, used for coordination and load balancing")
	kafkaTopicsCSV = flag.String("topics", DefaultKafkaTopics, "The comma-separated list of topics to consume")
	zookeeper      = flag.String("zookeeper", Defaultzk, "A comma-separated Zookeeper connection string (e.g. `zookeeper1.local:2181,zookeeper2.local:2181,zookeeper3.local:2181`)")

	zookeeperNodes []string
)

func init() {
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
}

func main() {
	flag.Parse()

	if *zookeeper == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	zookeeperNodes, config.Zookeeper.Chroot = kazoo.ParseConnectionString(*zookeeper)

	kafkaTopics := strings.Split(*kafkaTopicsCSV, ",")

	consumer, consumerErr := consumergroup.JoinConsumerGroup(*consumerGroup, kafkaTopics, zookeeperNodes, config)
	if consumerErr != nil {
		log.Fatalln(consumerErr)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if err := consumer.Close(); err != nil {
			sarama.Logger.Println("Error closing the consumer", err)
		}
	}()

	go func() {
		for err := range consumer.Errors() {
			log.Println(err)
		}
	}()

	eventCount := 0
	offsets := make(map[string]map[int32]int64)

	for message := range consumer.Messages() {
		if offsets[message.Topic] == nil {
			offsets[message.Topic] = make(map[int32]int64)
		}

		eventCount += 1
		if offsets[message.Topic][message.Partition] != 0 && offsets[message.Topic][message.Partition] != message.Offset-1 {
			log.Printf("Unexpected offset on %s:%d. Expected %d, found %d, diff %d.\n", message.Topic, message.Partition, offsets[message.Topic][message.Partition]+1, message.Offset, message.Offset-offsets[message.Topic][message.Partition]+1)
		}
		fmt.Print(" key:", string(message.Key))
		fmt.Print(" partition:", message.Partition)
		fmt.Print(" offset:", message.Offset)
		fmt.Println()

		// Simulate processing time
		//time.Sleep(10 * time.Millisecond)
		//time.Sleep(1 * time.Second)

		offsets[message.Topic][message.Partition] = message.Offset
		consumer.CommitUpto(message)
	}

	log.Printf("Processed %d events.", eventCount)
	log.Printf("%+v", offsets)
}
