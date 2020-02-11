package issue

import (
	"fmt"
	"github.com/google/go-github/github"
	"strings"
)

type IssueCommentEvent github.IssueCommentEvent

// if command is /pay fanux 10 ,type is : pay Command is : fanux 10
type Command struct {
	Type string // like pay
	Command string  // like 10, full command not contain type
}

type IssueEvent struct {
	*IssueCommentEvent
	Command *Command
	Client *github.Client
}

var robot map[string]Robot

type Robot interface {
	Process(event IssueEvent) error
}

// github user config
type Config struct {
	UserName string
	Password string
	Token string
}

func NewConfig(user string, passwd string) Config {
	return Config{UserName:user,Password:passwd}
}

func Process(config Config, event IssueCommentEvent) error{
	tp := github.BasicAuthTransport{
		Username:config.UserName,
		Password:config.Password,
	}
	client := github.NewClient(tp.Client())
	//TODO decode commands
	commands := decodeFromBody(event.Comment.Body)

	for _,command := range commands {
		issueEvent := IssueEvent{
			&event,
			command,
			client,
		}
		if v,ok := robot[command.Type];ok {
			v.Process(issueEvent)
		}
	}

	return nil
}

// Regist user robot
func Regist(command string, r Robot) {
	if robot == nil {
		robot = make(map[string]Robot)
	}
	robot[command] = r
}

func decodeFromBody(body *string) []*Command {
	var res []*Command
	lines := strings.Split(*body,"\n")
	for _,line := range lines {
		if !validCommand(line) {
			continue
		}
		res = append(res, decodeCommand(line))
	}
	return res
}

func validCommand(s string) bool {
	for _,b := range s {
		t := byte(b)
		if t != ' ' && t != '/' {
			return false
		}
		if t == '/' {
			return true
		}
		if t == ' ' {
			continue
		}
	}
	return false
}

// decode /pay 10 like command
func decodeCommand(s string) *Command {
	var command *Command
	var i,j int
	for i := range s {
		if byte(s[i]) == '/' {
			break
		}
	}
	var flag bool
	for j := i;j<len(s);j++ {
		if byte(s[j]) == ' ' {
			flag = true
			command.Type = s[i:j]
		}
		if flag && byte(s[j]) != ' ' {
			command.Command = s[j:]
			break
		}
	}
	fmt.Printf("decode command: %s, [%s][%s]",s,command.Type,command.Command)
	return command
}
