# GO WEB服务模板

简单mvc web服务
路由、控制器、数据服务

基于redis实现队列

自定义命令行

## 运行条件

go安装

gf安装 https://goframe.org/pages/viewpage.action?pageId=1115782

## 运行

```bash
gf run main.go
```

## 命令行执行示例

```bash
gf gen dao --tables=user 生成数据层dao文件

./main test:work 测试自定义命令行

./main queue:work --name=sms 自定义队列消费命令, name可选email
```