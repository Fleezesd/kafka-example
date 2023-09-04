# kafka demo 学习
因微服务后续需要链路追踪Opentelemetry的服务上需要用到Kafka
对消息队列不是很清晰 即对kafka go的client端进行学习

# v1 samara IBM的kafka的client学习
- 核心 syncProducer AsyncProducer 分别为同步和异步端的Producer
                                    channel * 的区别 
- 整体使用起来 
  - 缺点 
    文档较少
    源码文件较为混乱 
    上手较难  api不是特别透明
  - 优点
    底层源码架构部分设计的较为巧妙
    可以做为使用其他client端的参考
    性能很好


# v2 kafka-go client端学习
- 目前完成基本produce，topic consume 测试 
- Reader 简易test
- Writer 建设学习中

- 整体使用起来
  - 优点
    文档完善
    上手轻松， 以dial直接连通远端broker做操作
    主要是写代码方便很多 不用过于思考底层的kafka的过多原理 但凡事有利有弊 酌情处理 比较看好
    因代码不过于复杂 多人开发时还是比较易维护 若不是过分需求性能的话，个人推荐kafka-go


# v3 kafka-go reader writer SASL认证 logging
- reader
  - simple reader
  - consumer groups   consumer组
  - explicit commits 详述message 
  - managing commits  定期处理commits 即定时刷新commit 保证kafka的offset
- writer
  - simple writer
  - AutoTopicCreation 自动创建topic
  - Writing to multiple topics 写入多个topic
  - Compression writer压缩
- TLS
- SASL
- logging