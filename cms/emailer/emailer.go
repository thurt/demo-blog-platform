package emailer

import (
	"bytes"
	"net/smtp"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

type emailer struct {
	c *smtp.Client
}

func New(c *smtp.Client) pb.EmailerServer {
	return &emailer{c}
}

func (e *emailer) Send(ctx context.Context, email *pb.Email) (*empty.Empty, error) {
	// Set the sender and recipient.
	e.c.Mail(email.GetFrom())
	e.c.Rcpt(email.GetTo())
	// Send the email body.
	// TODO: I think there should be a mutex on m.c because m.c.Data() writer is supposed to be Closed before more commands are issued to m.c
	wc, err := e.c.Data()
	defer wc.Close()
	if err != nil {
		return nil, err
	}
	msg := "From: " + email.GetFrom() + "\n" +
		"To: " + email.GetTo() + "\n" +
		"Subject: " + email.GetSubject() + "\n\n" +
		email.GetBody()

	buf := bytes.NewBufferString(msg)
	if _, err = buf.WriteTo(wc); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

//	err := smtp.SendMail("mail:25",
//		smtp.PlainAuth("", from, pass, "mail"),
//		from, []string{to}, []byte(msg))
