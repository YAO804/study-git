package review

/*
今日内容：
依赖管理 go module
go.mod 的使用
go mod init [包名]   // 初始化项目
go mod tidy   // 检查代码里的依赖取更新go.mod文件中的依赖
go get
go mod download


context
非常重要！
如何优雅的控制 子goroutine 退出？

解决 goroutine 的超时退出问题
专门用来简化对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能设计多个API调用

两个默认值：
context.Background()
context.TODO()

with系列函数：
	WithCancel(context.Background())
	WithDeadline(context.Background(), time.Time)
	WithTimeout(context.Background(), time.Duration)
	WithValue(context.Background(), key, value)


服务端Agent开发
zookeeper、kafka部署
项目架构设计
Kafka 和 zookeeper
tailf介绍




*/
