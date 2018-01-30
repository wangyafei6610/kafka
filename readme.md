# GO kafka生产消费Demo
本项目可作为kafka生产消费入门学习

# 安装本地kafka环境

### 获取安装包
```
cd /usr/local/src
wget http://mirrors.tuna.tsinghua.edu.cn/apache/kafka/0.10.2.1/kafka_2.10-0.10.2.1.tgz
tar -zxvf kafka_2.10-0.10.2.1.tgz
cd kafka_2.10-0.10.2.1
```
### 启动zookeeper和kafka

```
bin/zookeeper-server-start.sh config/zookeeper.properties &
bin/kafka-server-start.sh config/server.properties &
```
### 创建查看topic
```
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test

or
bin/kafka-topics.sh --create --zookeeper 10.55.3.192:2181 --replication-factor 1 --partitions 1 --topic test

bin/kafka-topics.sh --list --zookeeper localhost:2181
```

### 发送 消费消息
```
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test


./bin/kafka-console-consumer.sh --zookeeper 10.55.3.192:2181 --topic test --from-beginning
or
./bin/kafka-console-consumer.sh --bootstrap-server 10.55.3.192:9092 --topic test --from-beginning
```

# 下载/安装
克隆项目到本地
### 启动producer
cd kafka && go run kafkaprducer.go

### 启动consume
cd kafka && go run kafkaconsume.go
