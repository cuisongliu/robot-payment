package issue

import (
"context"
"fmt"

"github.com/google/go-github/github"
)

// 回复issue
func (event IssueEvent)CommentBody(body string) error{
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
func (event IssueEvent)Label(label []string) error {
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
