# github robot
根据github的一些事件去触发处理一些事

处理issue中的一些诸如`/pay 100`的指令，可以自定义指令和注册指令的处理器。

使用场景：

> 自动回复issue

比如issue一天没回复，自动回复以让提issue的人莫着急

> 自动提醒

如kubernetes中/sig 指令会把消息推到对应的兴趣小组

> 自动化测试

`/test e2e test/test.sh`

> 触发drone promote事件

`/promote 42 test` 在issue里回复这个，就会触发CI/CD里面下面定义的一个pipeline

```
# 42 is the build number 
# drone build promote fanux/sealos 42 test
- name: e2etest
  image: golang:1.12
  commands:
    - echo "this is sealos test"
  when:
    event:
    - promote
    target:
    - test
```

```
    github     robot     drone
 issue |         |         |
------>| event   |         |
       |-------->|         |
       |         | promote |
       |         |-------->|
       |         |         | do what ever you want
       |         |         |
       V         V         V
```

> 自动merge代码

`/merge` 指令可以自动merge代码，还可以在merge之前之后做一些事，比如记录下PR的作者，发邮件，等等

> 付款

[sealos](https://github.com/fanux/sealos)的开发者是会有一定酬劳的，maintainer会把任务分解写成issue, 然后加个 `/pay 100`指令
机器会首先会自动给这个issue打上`paid`标签，然后开发者开发代码PR，一旦被merge就会自动把钱转入该开发者的支付宝账户。

> 其它

打标签，关闭超时issue等，

## 使用事例

这里使用一个faas跑一个机器人的例子帮助大家理解。 这个机器人是监听issue然后去触发CI的一些pipeline的功能。

比如我们为一个开发任务写了一个issue,开发人员开发完了PR过来，那么我们肯定希望跑一下测试用例再决定merge不merge。

此时就可以在issue下回复：

/promote 42 test key=value

然后机器人就会去触发drone的事件，执行drone pipeline中目标为test的步骤，来完成测试用例运行，当然具体怎么测试会由pipeline自己决定

```golang
// hello是个http handler, github 把事件数据以json格式发送过来，已经被解析到event结构体中
func promote(ctx context.Context, event issue.IssueCommentEvent) (string, error) {
    // or using env: GITHUB_USER GITHUB_PASSWD
    // github 账户名和密码，因为机器人可能还要回复issue什么的操作，这里建议单独给机器人申请个账号
    // 不传参数就会从环境变量中读取
    config := issue.NewConfig("sealrobot", "xxx")
    // regist what robot your need, and the robot config
    // 注册一下你希望哪个机器人处理，因为一条issue中可能会有很做指令，我们只关心/promote即可
    // Drone的处理器需要知道drone的地址和token是什么
    issue.Regist("promote", &drone_promote.DronePromote{"https://cloud.drone.io", "QSp93SmhZVpJAmb7tWPuWIOh3qs6BhuI"})
    // 处理issue
    err := issue.Process(config, event)
    return fmt.Sprintf("goversionecho %s", err), nil
}
```

## 扩展处理器

处理器就是为了处理issue中有自己感兴趣的指令。只要实现如下接口即可
```golang
type Robot interface {
	Process(event IssueEvent) error
}
```
注意要想Processor生效必须要注册处理器 issue.Regist(命令,处理器)

可以参考[drone-promote的实现](https://github.com/fanux/robot/blob/master/processor/drone_promote/drone_promote.go)

## 友情链接
[sealos-一键安装kubernetes HA集群](https://github.com/fanux/sealos)
