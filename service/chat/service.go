package chat

import (
	"context"
	"fmt"
	"regexp"
	"strings"
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

	satisfactionQues := "Bạn có hài lòng với câu trả lời này không?"
	if strings.Contains(strings.ToLower(userInput.Text), "segment") {
		return outputChat{
			Text: fmt.Sprintf("%s\n\n%s", segmentAnswer, satisfactionQues),
			Type: "satisfaction",
		}, nil
	}

	if strings.Contains(strings.ToLower(userInput.Text), "ott") {
		return outputChat{
			Text: fmt.Sprintf("%s\n\n%s", ottAnswer, satisfactionQues),
			Type: "satisfaction",
		}, nil
	}

	return outputChat{
		Text: "Câu hỏi của bạn không phổ biến, vui lòng tạo yêu cầu hỗ trợ",
		Type: "create_ticket",
	}, nil
}

var regex = regexp.MustCompile(`TEK-\d+`)

func isTicketKey(s string) bool {
	return regex.MatchString(s)
}

func NewService() IService {
	return &service{}
}
