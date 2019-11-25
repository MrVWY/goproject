[English document](https://github.com/MrVWY/go/blob/master/micro/emamples/README.md) | Chinese document  
  
这是微服务的项目

如果您对Micro感兴趣，则可以使用该网址（https://github.com/micro-in-cn/tutorials) 来学习Micro  
现在，我将向您介绍此目录的结构以及如何运行

如果您对Micro有任何疑问，可以在Issues上留言并与您联系  

由于我仍处于学习阶段，它将稍后在docker上运行，这将是一个很大的麻烦  
____
一、  
====
您需要了解什么是Micro，什么是go-micro。
  
Micro是微服务框架（或工具的集合），它提供了各种组件来解决微服务中的不同问题架构，服务监视，服务发现，保险丝机制，负载平衡...，并且您还可以看到[Micro的文档]（https://micro.mu/docs/index.html),中文版文档，请参阅[Micro的文档]（https://micro.mu/docs/cn/index.html)  
  
go-micro框架提供了服务发现，负载平衡，同步传输，
异步通信和事件驱动。 它试图简化分布式系统之间的通信，
使我们能够专注于开发自己的业务逻辑。 
  
go-micro 的架构图如下所示:  
  
![alt](https://image-static.segmentfault.com/180/641/1806415878-5c28d8815645a)  
  
go-micro的组件包括 :  
* Registry组件：服务发现组件，提供服务发现机制：解析服务名字至服务地址。目前支持的注册中心有consul、etcd、 zookeeper、dns、gossip等.(Tip: In a recent update, the consul isn't use.)
- Selector组件：构建在Registry之上的客户端智能负载均衡组件，用于Client组件对Registry返回的服务进行智能选择.
* Broker组件：发布/订阅组件，服务之间基于消息中间件的异步通信方式，默认使用http方式，线上通常使用消息中间件，如Kafka、RabbitMQ等.
- Transport组件：服务之间同步通信方式.
* Codec组件：服务之间消息的编码/解码.
- Server组件：服务主体，该组件基于上面的Registry/Selector/Transport/Broker组件，对外提供一个统一的服务请求入口.
* Client组件：提供访问微服务的客户端。类似Server组件，它也是通过Registry/Selector/Transport/Broker组件实现查找服务、负载均衡、同步通信、异步消息等功能.  
  
所有以上组件功能共同构成一个go-micro微服务.  
 
Sencond
====
you need download the tools ( but Micro seems to have included all the tools,so you just need download the micro.But in order to clearly understand the tools we need to get into this project, I will list them below. Note : there may be some tools which is you need,like protoc-gen-gorm..., you need use by yourself.In here I will not say much):  
* protobuf  
```
go get github.com/micro/protoc-gen-micro  
go get github.com/micro/protoc-gen-go
```
- Micro  
```
go get github.com/micro/micro
```
  
Third
====
* Micro demo  
you can see [this](https://github.com/micro-in-cn/tutorials/tree/master/examples/middle-practices/micro-new)
