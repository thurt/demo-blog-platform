package emailer

import (
	"net/smtp"

	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

type emailer struct {
	smtpAddr string
	smtpAuth smtp.Auth
}

func New(smtpAddr string, smtpAuth smtp.Auth) pb.EmailerServer {
	return &emailer{smtpAddr, smtpAuth}
}

func (e *emailer) Send(ctx context.Context, email *pb.Email) (*empty.Empty, error) {
	msg := "From: " + email.GetFrom() + "\n" +
		"To: " + email.GetTo() + "\n" +
		"Subject: " + email.GetSubject() + "\n\n" +
		email.GetBody()

	err := smtp.SendMail(e.smtpAddr, e.smtpAuth, email.GetFrom(), []string{email.GetTo()}, []byte(msg))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
