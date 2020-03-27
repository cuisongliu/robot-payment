package apply

import (
	"fmt"
	"github.com/fanux/robot/issue"
)

type Apply struct {
}

func (p *Apply) Process(event issue.IssueEvent) error {
	var as []string
	var err error
	if as,err = event.GetAssign();err != nil {
		return err
	}
	if len(as) != 0 {
		event.CommentBody(fmt.Sprintf("@%s 任务已经分配给了%s, 重复申请没用",*event.Comment.User.Login,as[0]))
		return nil
	}

	err = event.SetAssign(*event.Comment.User.Login)
	if err != nil {
		event.CommentBody(fmt.Sprintf("任务分配失败，%s",err))
		return err
	}
	return event.CommentBody(fmt.Sprintf("@%s 你成功申请了任务，任务完成后使用/alipay xxx 提供支付宝号",*event.Comment.User.Login))
}

