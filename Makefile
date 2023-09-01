# Makefile for Kafka Topic and Consumer Operations

# Kafka Broker地址和端口
BROKER=localhost:9092

# 主题名称
TOPIC=my_topic

# 查看主题列表
list-topics:
	kafka-topics --list --bootstrap-server $(BROKER)

# 创建新主题
create-topic:
	kafka-topics --create --topic $(TOPIC) --bootstrap-server $(BROKER) --partitions 1 --replication-factor 1

# 删除主题
delete-topic:
	kafka-topics --delete --topic $(TOPIC) --bootstrap-server $(BROKER)

# 消费主题消息
consume-topic:
	kafka-console-consumer --topic $(TOPIC) --bootstrap-server $(BROKER) --from-beginning

# 发送测试消息到主题
send-test-message:
	kafka-console-producer --topic $(TOPIC) --broker-list $(BROKER)

.PHONY: list-topics create-topic delete-topic consume-topic send-test-message
