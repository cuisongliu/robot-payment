package pay

import (
	"context"
	"fmt"
	"github.com/fanux/robot/issue"
	"github.com/google/go-github/github"
	"time"
)

type Pay struct {
}

func (p *Pay)Process(event issue.IssueEvent) error {
	cmd := event.Command.Command

	body := fmt.Sprintf("完成这个issue中的任务，代码成功合并，fanux老板就会支付你 %s 元, \n" +
		"请在issue中回复 \n" +
		"/alipay [你的支付宝号] （如 /alipay 15281817171）\n" +
		"以让我知道您的支付宝",cmd)
	comment := &github.IssueComment{
		Body: &body,
	}
	fmt.Printf("comment issue %s %s %d %s",*event.IssueCommentEvent.Repo.Name,
		*event.IssueCommentEvent.Repo.Owner.Login,*event.IssueCommentEvent.Issue.Number,comment)

	ctx := context.Background()
	_,_,err := event.Client.Issues.CreateComment(ctx,*event.IssueCommentEvent.Repo.Owner.Login,*event.IssueCommentEvent.Repo.Name,*event.IssueCommentEvent.Issue.Number,comment)

	if err != nil {
		fmt.Printf("comment issue failed #{err}")
	}

	time.Sleep(time.Second * 5)
	return err
}