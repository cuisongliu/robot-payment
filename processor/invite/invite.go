package invite

import (
	"fmt"
	"github.com/fanux/robot/issue"
)

type Invite struct {
}

func (p *Invite) Process(event issue.IssueEvent) error {
	who := event.Command.Command
	return event.CommentBody(fmt.Sprintf("@%s 请加QQ群：98488045 详细探讨一下，或者加微信：sealnux 拉你进微信群", who))
}
