package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"os"
	"time"
)


var (
	producer sarama.SyncProducer
	host = []string{"localhost:9092"}
)


func main()  {
	defer func() {
		if err:= recover();err!= nil{
			fmt.Printf("main|defer error|err=%v",err)
		}
	}()
	initProducer()
	for{
		fmt.Println("start")
		produce("test","my name is lin")
		fmt.Println("end")
		time.Sleep(time.Second)

	}
}

func initProducer()  {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//config.Producer.Timeout = 3*time.Second
	config.Producer.Return.Successes = true
	var err error
	producer,err = sarama.NewSyncProducer(host,config)
	if err != nil{
		fmt.Printf("initProducer|sarama.NewSyncProducer error|err=%v\n",err)
		os.Exit(0)
	}
	//defer producer.Close()
}

func produce(topic string,data string) {
	msg := &sarama.ProducerMessage{
		Topic:topic,
		Key:sarama.StringEncoder("key"),
		Partition:int32(0),
	}

	//_,err := fmt.Scanf("%s",&data)
	//if err != nil{
	//	fmt.Printf("produce|fmt.Scanf error|err=%v\n",err)
	//}
	msg.Value = sarama.ByteEncoder(data)
	fmt.Printf("%+v\n",data)
	partition,offset,err := producer.SendMessage(msg)
	if err != nil{
		fmt.Printf("produce|producer.SendMessage error|partition=%+v,offset=%+v,err=%+v\n",partition,offset,err)
		return
	}
	fmt.Printf("sendMsg|topic=%s,msg=%+v, partition=%+v,offset=%+v,err=%v\n",topic,msg,partition,offset,err)
}
