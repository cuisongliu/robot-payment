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

[sealos](https://github.com/fanux/sealos) 的开发者是会有一定酬劳的，maintainer会把任务分解写成issue, 然后加个 `/pay 100`指令
机器会首先会自动给这个issue打上`paid`标签，然后开发者开发代码PR，一旦被merge就会自动把钱转入该开发者的支付宝账户。

> 其它

打标签，关闭超时issue等，

## 开发教程

Event 中会存放事件的一些信息以及一个用于访问和操作github的client,还有触发事件的指令
```
type Event struct {
    EventInfo
    Client 
    Command string //如 /test e2e
}
```

以开发一个处理/test指令的处理器为例，用户只需要实现一个处理器并注册即可

```
type TestRobot struct{
    //你需要的信息
}

func (t *TestRobot)Processor(event Event){
    // 处理任务,监听到github事件此函数就会被回调
}

Regist("test", &TestRobot) // test是指令名字 `/test e2e` 这样这个处理器不会处理别的指令如 `/pay 8`
```
