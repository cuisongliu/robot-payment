package pay

import (
	"context"
	"fmt"
	"github.com/fanux/robot/issue"
	"github.com/google/go-github/github"
)

type Pay struct {
}

func (p *Pay)Process(event issue.IssueEvent) error {
	if *event.Action != "created" {
		fmt.Printf("not create isuue skip")
		return nil
	}

	cmd := event.Command.Command

	body := fmt.Sprintf("完成这个issue中的任务，代码成功合并，%s 老板就会支付你 %s 元, \n" +
		"请在issue中回复:\n" +
		"/alipay [你的支付宝号] （如 /alipay 15281817171）\n" +
		"以让我知道您的支付宝",*event.Comment.User.Login,cmd)
	comment := &github.IssueComment{
		Body: &body,
	}
	fmt.Printf("comment issue %s %s %d %s",*event.IssueCommentEvent.Repo.Name,
		*event.IssueCommentEvent.Repo.Owner.Login,*event.IssueCommentEvent.Issue.Number,comment)

	owner :=*event.IssueCommentEvent.Repo.Owner.Login
	repo :=*event.IssueCommentEvent.Repo.Name
		num:=*event.IssueCommentEvent.Issue.Number
	ctx := context.Background()
	_,_,err := event.Client.Issues.CreateComment(ctx,owner,repo,num,comment)
	if err != nil {
		fmt.Printf("comment issue failed %s",err)
	}

	ctx = context.Background()
	fmt.Printf("Add label to issue")
	_,_,err = event.Client.Issues.AddLabelsToIssue(ctx,owner,repo,num,[]string{"paid"})
	if err != nil {
		fmt.Printf("issue label failed %s",err)
	}

	return err
}