package chat

import "context"

type IService interface {
	Answer(ctx context.Context, userInput inputChat) (outputChat, error)
}

type service struct{}

func (s service) Answer(ctx context.Context, userInput inputChat) (outputChat, error) {
	return outputChat{
		Text: "Bạn có hài lòng không hả?",
		Type: "satisfaction",
	}, nil
}

func NewService() IService {
	return &service{}
}
