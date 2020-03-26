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
		event.CommentBody(fmt.Sprintf("任务已经分配给了%s, 重复申请没用",as[0]))
		return nil
	}

	err = event.SetAssign(*event.Comment.User.Login)
	if err != nil {
		event.CommentBody(fmt.Sprintf("任务分配失败，%s",err))
		return err
	}

	return nil
}

