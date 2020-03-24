package pay

import (
	"fmt"
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
	err := event.CommentBody(body)
	if err != nil {
		return err
	}
	err = event.Label([]string{"paid"})

	return err
}

type PayTo struct {
	Account string
	Money   int
	PayClient *utils.Alipay
}

var UserAliaccountMap = map[string]string{"fanux":"15805691422","PatHoo":"13926139093","cuisongliu":"912387319@qq.com","zhangguanzhang":"zhangguanzhang@qq.com"}

func GetAlipayAccount(user string) string {
	if v,ok := UserAliaccountMap[user];ok {
		return v
	}
	return ""
}

func (p *PayTo) Process(event issue.IssueEvent) error {
	s := utils.SplitMultiBlank(event.Command.Command)
	if len(s) != 2 {
		return fmt.Errorf("pay to command error %s", event.Command.Command)
	}
	var body string
	if *event.Repo.Owner.Login != *event.Comment.User.Login {
		body = fmt.Sprintf("你不是项目owner [%s] 不能进行付款操作！",*event.Repo.Owner.Login)
		return event.CommentBody(body)
	}
	account := GetAlipayAccount(s[0])
	if account == "" {
		body = fmt.Sprintf("@%s 请[使用/alipay xxx提供你的支付宝号，如：\n/alipay 15805691422",s[0])
		return event.CommentBody(body)
	}
	remark := fmt.Sprintf("感谢您完成了%s issue %d, 希望能继续参与我们的伟业",*event.Repo.FullName,*event.Issue.Number)
	err := p.PayClient.PayTo(account,s[1],remark)
	if err != nil {
		body = fmt.Sprintf("支付失败，联系管理员处理 %s", err)
		return event.CommentBody(body)
	}
	body = fmt.Sprintf("已经转账[%s]元到[%s]的支付宝账户,注意查收", s[1], s[0])
	return event.CommentBody(body)
}
