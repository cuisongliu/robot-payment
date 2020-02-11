// 使用腾讯faas跑robot
package main

import (
	"context"
	"fmt"

	"github.com/fanux/robot/issue"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func hello(ctx context.Context, event issue.IssueCommentEvent) (string, error) {
	err := issue.Process(event)
	return fmt.Sprintf("goversionecho %s", err), nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(hello)
}
