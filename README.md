# trade - 电商服务
trade为电商单机版本，目的是为了对项目需求探底并通过呀测获取监控指标，为集群版本的开发铺路。

# 目录结构
* assets：静态文件(存放配置等)
* cmd: 运行文件(存放服务启动文件)
* core: 核心代码
* logs: 日志文件
* pkg: 辅助插件
* router: 路由分配

# 技术栈
* 版本控制：git/github
* 开发语言: golang
* web 框架: gin
* grpc/secret_key
* 数据库: mongodb
* 消息队列: kafka
* 容器: docker/docker-compose/k8s  