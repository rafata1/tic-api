package chat

import (
	"context"
	"regexp"
)

type IService interface {
	Answer(ctx context.Context, userInput inputChat) (outputChat, error)
}

type service struct{}

func (s service) Answer(ctx context.Context, userInput inputChat) (outputChat, error) {
	if isTicketKey(userInput.Text) {
		return outputChat{
			Text: "Yêu cầu của bạn đã được tiếp nhận bởi Nguyên và sẽ được xử lý trong thời gian sớm nhất!",
			Type: "issue_tracking",
		}, nil
	}
	return outputChat{
		Text: "Bạn có hài lòng không hả?",
		Type: "satisfaction",
	}, nil
}

var regex = regexp.MustCompile(`TEK-\d+`)

func isTicketKey(s string) bool {
	return regex.MatchString(s)
}

func NewService() IService {
	return &service{}
}
