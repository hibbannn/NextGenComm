package contract

type MailContract interface {
	Send(to []string, subject, body string) error
}
