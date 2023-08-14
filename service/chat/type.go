package chat

type inputChat struct {
	Text string `json:"text"`
}

type outputChat struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
