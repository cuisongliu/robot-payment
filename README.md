# github payment robot
Trigger and process some things according to some events on github

To process some instructions in the issue, such as `/pay 100`, you can customize the instructions and register the instructions of the processor.

scenes to be used:

> Auto reply to issue

For example, if there is no reply to the issue in one day, the automatic reply will make the person who raised the issue not anxious.

> Automatic reminder

For example, the /sig command in kubernetes will push the message to the corresponding interest group

> Automated Testing

`/test e2e test/test.sh`

> Trigger drone promote event

If `/promote 42 test` responds to this in the issue, it will trigger a pipeline defined below in CI/CD

```
# 42 is the build number
# drone build promote fanux/sealos 42 test
-name: e2etest
  image: golang:1.12
  commands:
    -echo "this is sealos test"
  when:
    event:
    -promote
    target:
    -test
```

```
    github robot drone
 issue | | |
------>| event | |
       |-------->| |
       | | promote |
       | |-------->|
       | | | do what ever you want
       | | |
       V V V
```

> Automatic merge code

The `/merge` command can automatically merge the code, and can also do some things before and after the merge, such as recording the author of the PR, sending an email, etc.

> Payment

[sealos](https://github.com/fanux/sealos) developers will be paid a certain amount, the maintainer will break down the task into issues, and then add a `/pay 100` command
The machine will first automatically tag the issue with the `paid` label, and then the developer develops the code PR, and once it is merged, it will automatically transfer the money to the developer's Alipay account.

> Other

Tagging, closing timeout issues, etc.,

## Use case

Here is an example of faas running a robot to help everyone understand. This robot monitors the issue and then triggers some pipeline functions of CI.

For example, if we write an issue for a development task, and the developer has finished developing the PR, we definitely hope to run the test case before deciding whether to merge or merge.

At this point, you can reply under the issue:

/promote 42 test key=value

Then the robot will trigger the drone event and execute the step of the drone pipeline with the goal of test to complete the test case operation. Of course, the specific test will be determined by the pipeline itself.

```golang
// hello is an http handler, github sends the event data in json format, which has been parsed into the event structure
func promote(ctx context.Context, event issue.IssueCommentEvent) (string, error) {
    // or using env: GITHUB_USER GITHUB_PASSWD
    // github account name and password, because the robot may have to reply to issues or other operations, it is recommended to apply for an account for the robot separately
    // Without passing parameters, it will be read from environment variables
    config := issue.NewConfig("sealrobot", "xxx")
    // regist what robot your need, and the robot config
    // Register which robot you want to handle, because there may be a lot of instructions in an issue, we only care about /promote
    // Drone's processor needs to know the address and token of the drone
    issue.Regist("promote", &drone_promote.DronePromote{"https://cloud.drone.io", "QSp93SmhZVpJAmb7tWPuWIOh3qs6BhuI"})
    // deal with issue
    err := issue.Process(config, event)
    return fmt.Sprintf("goversionecho %s", err), nil
}
```

## Extended processor

The processor is to process the instructions of interest in the issue. Just implement the following interface
```golang
type Robot interface {
Process(event IssueEvent) error
}
```
Note that for Processor to take effect, you must register the processor issue.Regist (command, processor)

You can refer to [Implementation of drone-promote](https://github.com/fanux/robot/blob/master/processor/drone_promote/drone_promote.go)
