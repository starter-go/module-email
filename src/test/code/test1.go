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

	ToAddr string //starter:inject("${mails.test.to-addr}")
}

// Life ...
func (inst *Test1) Life() *application.Life {
	return &application.Life{OnStartPost: inst.run}
}

func (inst *Test1) run() error {

	ctx := context.Background()
	msg := &mails.Message{}

	text := "hello, world"
	toAddr := inst.ToAddr

	msg.ToAddresses = []mails.Address{mails.Address(toAddr)}
	msg.Title = "hello"
	msg.ContentType = "text/plain"
	msg.Content = []byte(text)

	return inst.Sender.Send(ctx, msg)
}
