package ticket

type inputTicket struct {
	SupportType    string   `json:"support_type"`
	Description    string   `json:"description"`
	AttachmentURLs []string `json:"attachment_urls"`
	CustomerEmail  string   `json:"customer_email"`
}

type outputTicket struct {
	Key string `json:"key"`
}
