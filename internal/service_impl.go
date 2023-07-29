package internal

import (
	"context"
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/module-email/mails"
)

// SenderService ...
type SenderService struct {

	//starter:component
	_as func(mails.Service) //starter:as("#")

	Regs              []mails.DispatcherRegistry //starter:inject(".")
	DefaultSenderAddr string                     //starter:inject("${mails.default-sender-address}")

	cached []mails.Dispatcher
}

func (inst *SenderService) _impl() mails.Service {
	return inst
}

// Life ...
func (inst *SenderService) Life() *application.Life {
	return &application.Life{OnCreate: inst.init}
}

func (inst *SenderService) init() error {
	_, err := inst.tryGetDispatcherList()
	return err
}

func (inst *SenderService) getDefaultSenderAddress() mails.Address {
	str := inst.DefaultSenderAddr
	return mails.Address(str)
}

// Send ...
func (inst *SenderService) Send(c context.Context, msg *mails.Message) error {

	if msg.FromAddress == "" {
		msg.FromAddress = inst.getDefaultSenderAddress()
	}

	err := fmt.Errorf("no dispatcher accept the mails.Message")
	list := inst.getDispatcherList()
	for _, item := range list {
		if item.Accept(c, msg) {
			e2 := item.Send(c, msg)
			if e2 == nil {
				return nil
			}
			err = e2
		}
	}
	return err
}

func (inst *SenderService) getDispatcherList() []mails.Dispatcher {
	list, err := inst.tryGetDispatcherList()
	if err != nil {
		panic(err)
	}
	return list
}

func (inst *SenderService) tryGetDispatcherList() ([]mails.Dispatcher, error) {
	list := inst.cached
	if list == nil {
		li, err := inst.loadDispatcherList()
		if err != nil {
			return nil, err
		}
		list = li
		inst.cached = li
	}
	return list, nil
}

func (inst *SenderService) loadDispatcherList() ([]mails.Dispatcher, error) {
	src := inst.Regs
	dst := make([]mails.Dispatcher, 0)
	for _, r1 := range src {
		list1 := r1.ListRegistrations()
		for _, r2 := range list1 {
			d := inst.getDispatcher(r2)
			if d == nil {
				continue
			}
			dst = append(dst, d)
		}
	}
	return dst, nil
}

func (inst *SenderService) getDispatcher(r *mails.DispatcherRegistration) mails.Dispatcher {
	if r == nil {
		return nil
	}
	if !r.Enabled {
		return nil
	}
	return r.Dispatcher
}
