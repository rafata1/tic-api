package errors

type DomainError struct {
	Code    int
	Message string
}

func (d DomainError) Error() string {
	return d.Message
}

func New(code int, message string) error {
	return DomainError{
		Code:    code,
		Message: message,
	}
}
