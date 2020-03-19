package pay

import (
	"fmt"
	"github.com/fanux/robot/client_utils"
	"github.com/fanux/robot/issue"
	"github.com/fanux/robot/utils"
)

type Pay struct {
}

func (p *Pay) Process(event issue.IssueEvent) error {
	if *event.Action != "created" {
		fmt.Printf("not create isuue skip")
		return nil
	}

	cmd := event.Command.Command

	body := fmt.Sprintf("完成这个issue中的任务，代码成功合并，%s 老板就会支付你 %s 元, \n"+
		"请在issue中回复:\n"+
		"/alipay [你的支付宝号] （如 /alipay 15281817171）\n"+
		"以让我知道您的支付宝", *event.Comment.User.Login, cmd)
	/*
	comment := &github.IssueComment{
		Body: &body,
	}
	fmt.Printf("comment issue %s %s %d %s", *event.IssueCommentEvent.Repo.Name,
		*event.IssueCommentEvent.Repo.Owner.Login, *event.IssueCommentEvent.Issue.Number, comment)

	owner := *event.IssueCommentEvent.Repo.Owner.Login
	repo := *event.IssueCommentEvent.Repo.Name
	num := *event.IssueCommentEvent.Issue.Number
	ctx := context.Background()
	_, _, err := event.Client.Issues.CreateComment(ctx, owner, repo, num, comment)
	if err != nil {
		fmt.Printf("comment issue failed %s", err)
	}
	*/
	err := event.Comment(body)
	if err != nil {
		return err
	}
	err = event.Label([]string{"paid"})

	return err
}

type PayTo struct {
	Account string
	Money int
}

func (p *PayTo)Process(event issue.IssueEvent) error {
	s := utils.SplitMultiBlank(event.Command.Command)
	if len(s) != 2 {
		return fmt.Errorf("pay to command error %s", event.Command.Command)
	}
	body := fmt.Sprintf("已经转账[%s]元到支付宝账户[%s],注意查收",s[1],s[0])
	return event.Comment(body)
}
