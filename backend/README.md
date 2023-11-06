# mytiktok-backend

## Description

Gin + Gorm + JWT + MySQL + Redis + RabbitMQ

## Technical Selection

- Web Framework：[Gin](https://gin-gonic.com/)
- ORM Framework：[GORM](https://gorm.io/)
- Security Framework：[JWT](https://jwt.io/)
- DataBase：[MySQL](https://www.mysql.com/cn/)
- Cache：[Redis](https://redis.io/)
- Message Queue：[RabbitMQ](https://www.rabbitmq.com/)

## System Structure

项目的整体架构采用经典的 MVC 三层架构，客户端层访问路由层，不同的接口请求路由到不同的业务层，期间经过鉴权层对用户信息进行鉴权认证，通过通过验证的用户请求方可进入到业务层，业务层主要包含了视频流，视频发布，点赞，关注，评论，消息和用户等业务。业务层中不同的业务通过调用服务层中的各项服务来实现接口定义的 Response，服务层主要和数据层以及 Redis，RabbitMQ 等中间件进行交互，获取到数据层的原始数据模型，再做相应的封装，返回给上层服务。数据层主要和数据存储层，其中包括像 MySQL 数据库，OSS 对象存储进行交互。最底层则是基础设施层，包含项目部署所要用到的 Docker，Go Runtime，云服务器等。
