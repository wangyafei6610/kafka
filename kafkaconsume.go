package main

import (
	"os"
	"strings"
	"os/signal"
	"fmt"
	"github.com/bsm/sarama-cluster"
	"github.com/Shopify/sarama"
)

var (
	brokerList string = "127.0.0.1:9092"
	groupId string  = "wang2"
	topics string = "test"
	signchan chan os.Signal
	consume *cluster.Consumer
)

func main()  {
	initConsumer()
	consumeInfo()
}

func initConsumer (){

	address := make([]string,0)
	topic := make([]string,0)
	for _,adrsOne := range strings.Split(brokerList,","){
		address = append(address, adrsOne)
	}
	for _,tipcOne := range strings.Split(topics,","){
		topic = append(topic,tipcOne)
	}

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	var err error
	fmt.Printf("配置文件|address=%+v,groupId=%+v,topic=%+v\n",address,groupId,topic)
	consume,err = cluster.NewConsumer(address,groupId,topic,config)
	if err != nil {
		fmt.Printf("cluster.NewConsumer|error err=%v\n",err)
		return
	}
	signchan = make(chan os.Signal,1)
	signal.Notify(signchan,os.Interrupt)
}

func consumeInfo()  {
	for {
		select {

		case msg := <-consume.Messages():
			fmt.Printf("pc.Messages|getMsg=%+v\n",msg)
			consume.MarkOffset(msg,"")
			consume.CommitOffsets()
		//case part,ok := <-consume.Partitions():
		//	if !ok {
		//		return
		//	}
		//	go func(pc cluster.PartitionConsumer) {
		//		for msg := range pc.Messages(){
		//			fmt.Printf("pc.Messages|getMsg=%+v\n",msg)
		//		}
		//	}(part)

		case nt := <-consume.Notifications():
			fmt.Printf(" consume.Notifications|nt=%+v\n",nt)

		case errs := <-consume.Errors():
			fmt.Printf("consume.Errors|errs=%+v\n",errs)
			
		case s := <-signchan:
			fmt.Printf("signchan|获取终端信号 = %v\n",s)
			consume.Close()
			return
		}
	}
}