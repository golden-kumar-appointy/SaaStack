package email

type EmailInterface interface {
	Send(from string, to string, message string) error
}
