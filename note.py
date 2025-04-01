"""
客户端 或者 服务端挂掉了


kafka 做分布式 作消息的存储或者效率 比 rabbitMQ快一倍

- Queue 将数据存醋当服务器的内存
- redis 列表
- rabiitMQ /kafka/zeroMQ(专业作消息队列)

无论什么场景 都是要推送服务器

- 1000台服务器  salstack
- saltstack
 - ssh  安装方便  但是执行效率慢
 - agent: 执行效率高（RPC）（基于消息队列zeroMQ做的RPC）.
 监听队列 执行一条命令
2. 公司什么时候会用消息队列？
    请求数量太多，需要把消息临时放到某个地方
    100个人消费者的目标一致
    发布/订阅

    服务端:


    客户端:
        pip3 install pika

"""
