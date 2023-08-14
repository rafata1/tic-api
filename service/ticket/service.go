package ticket

import (
	"context"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/rafata1/tic-api/service/common"
	"log"
	"strings"
)

type IService interface {
	CreateTicket(ctx context.Context, input inputTicket) (outputTicket, error)
}

type service struct {
	jiraCli *jira.Client
}

func (s service) CreateTicket(ctx context.Context, input inputTicket) (outputTicket, error) {
	jiraIssue := toJiraTicket(input)
	issue, _, err := s.jiraCli.Issue.Create(jiraIssue)
	if err != nil {
		log.Printf("error creating jira issue: %s\n", err.Error())
		return outputTicket{}, common.ErrCall3rdParty
	}
	return outputTicket{
		Key: issue.Key,
	}, nil
}

func toJiraTicket(input inputTicket) *jira.Issue {
	return &jira.Issue{
		Fields: &jira.IssueFields{
			Description: fmt.Sprintf(
				"%s \nAttachments: %s",
				input.Description,
				strings.Join(input.AttachmentURLs, "\n"),
			),
			Type: jira.IssueType{Name: "Task"},
			Project: jira.Project{
				Key: "TEK",
			},
			Summary: fmt.Sprintf(
				"[%s] %s",
				input.SupportType,
				input.CustomerEmail,
			),
		},
	}
}

func NewService(jiraCli *jira.Client) IService {
	return &service{
		jiraCli: jiraCli,
	}
}
