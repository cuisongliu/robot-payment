// 使用腾讯faas跑robot
package main

import (
	"context"
	"fmt"
	"github.com/fanux/robot/processor/drone_promote"
	"time"

	"github.com/fanux/robot/issue"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func hello(ctx context.Context, event issue.IssueCommentEvent) (string, error) {
	// or using env: GITHUB_USER GITHUB_PASSWD
	config := issue.NewConfig("sealrobot", "xxx")
	// regist what robot your need, and the robot config
	issue.Regist("promote", &drone_promote.DronePromote{"https://cloud.drone.io", "QSp93SmhZVpJAmb7tWPuWIOh3qs6BhuI"})
	err := issue.Process(config, event)
	return fmt.Sprintf("goversionecho %s", err), nil
}

func main() {
	t1 := time.Now()
	d := t1.Sub(t1)
	d-time.Hour
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(hello)
}
