package issue

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func (event IssueEvent) CloseIssue() error {
	state := "closed"
	owner := *event.IssueCommentEvent.Repo.Owner.Login
	repo := *event.IssueCommentEvent.Repo.Name
	num := *event.IssueCommentEvent.Issue.Number
	ctx := context.Background()
	req := &github.IssueRequest{
		State: &state,
	}
	_,_,err := event.Client.Issues.Edit(ctx,owner,repo,num,req)
	return err
}

// 回复issue
func (event IssueEvent) CommentBody(body string) error {
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
		return fmt.Errorf("comment issue failed %s", err)
	}
	return nil
}

// 给issue贴标签
func (event IssueEvent) Label(label []string) error {
	owner := *event.IssueCommentEvent.Repo.Owner.Login
	repo := *event.IssueCommentEvent.Repo.Name
	num := *event.IssueCommentEvent.Issue.Number

	ctx := context.Background()
	fmt.Printf("Add label to issue")
	_, _, err := event.Client.Issues.AddLabelsToIssue(ctx, owner, repo, num, label)
	if err != nil {
		return fmt.Errorf("issue label failed %s", err)
	}

	return nil
}

// 获取issue分配人
func (event IssueEvent) GetAssign() ([]string, error) {
	users := event.Issue.Assignees
	var res []string
	for _,u := range users {
		res = append(res,*u.Login)
	}
	fmt.Printf("issue assgin user: %s", res)
	return res,nil
}

// 分配任务
func (event IssueEvent) SetAssign(user string) error {
	owner := *event.IssueCommentEvent.Repo.Owner.Login
	repo := *event.IssueCommentEvent.Repo.Name
	num := *event.IssueCommentEvent.Issue.Number

	_,_,err := event.Client.Issues.AddAssignees(context.Background(),owner,repo,num,[]string{user})
	return err
}
