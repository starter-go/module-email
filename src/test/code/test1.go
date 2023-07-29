package code

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/module-email/mails"
)

// Test1 ...
type Test1 struct {
	//starter:component

	Sender mails.Service //starter:inject("#")
}

// Life ...
func (inst *Test1) Life() *application.Life {
	return &application.Life{OnStartPost: inst.run}
}

func (inst *Test1) run() error {

	ctx := context.Background()
	msg := &mails.Message{}

	text := "hello, world"

	msg.ToAddress = "foo@bar"
	msg.Title = "hello"
	msg.ContentType = "text/plain"
	msg.Content = []byte(text)

	return inst.Sender.Send(ctx, msg)
}
