package internal

import (
	"context"
	"encoding/json"

	"github.com/starter-go/module-email/mails"
	"github.com/starter-go/vlog"
)

// SMSDriver ...
type SMSDriver struct {
	//starter:component
	_as func(mails.DriverRegistry) //starter:as(".")
}

func (inst *SMSDriver) _impl() (mails.Driver, mails.DriverRegistry) {
	return inst, inst
}

func (inst *SMSDriver) name() string {
	return "sms"
}

// ListRegistrations ...
func (inst *SMSDriver) ListRegistrations() []*mails.DriverRegistration {
	name := inst.name()
	r1 := &mails.DriverRegistration{
		Name:    name,
		Enabled: true,
		Driver:  inst,
	}
	return []*mails.DriverRegistration{r1}
}

// Accept ...
func (inst *SMSDriver) Accept(cfg *mails.Configuration) bool {
	return cfg.Driver == inst.name()
}

// CreateDispatcher ...
func (inst *SMSDriver) CreateDispatcher(cfg *mails.Configuration) (mails.Dispatcher, error) {
	d := &mySMSDriverDispatcher{}
	d.config = *cfg
	return d, nil
}

////////////////////////////////////////////////////////////////////////////////

type mySMSDriverDispatcher struct {
	config mails.Configuration
}

func (inst *mySMSDriverDispatcher) _impl() mails.Dispatcher {
	return inst
}

func (inst *mySMSDriverDispatcher) Accept(c context.Context, msg *mails.Message) bool {
	addr1 := msg.FromAddress
	addr2 := inst.config.SenderAddress
	return addr1 == addr2
}

func (inst *mySMSDriverDispatcher) Send(c context.Context, msg *mails.Message) error {
	const (
		prefix = ""
		indent = "\t"
	)
	data, err := json.MarshalIndent(msg, prefix, indent)
	if err != nil {
		return err
	}
	vlog.Info("sms_email_dispatcher: send SMS %s", string(data))
	return nil
}
