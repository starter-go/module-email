package internal

import (
	"context"

	"github.com/starter-go/module-email/mails"
	"github.com/starter-go/vlog"
)

// MockDriver ...
type MockDriver struct {
	//starter:component
	_as func(mails.DriverRegistry) //starter:as(".")
}

func (inst *MockDriver) _impl() (mails.Driver, mails.DriverRegistry) {
	return inst, inst
}

func (inst *MockDriver) name() string {
	return "mock"
}

func (inst *MockDriver) ListRegistrations() []*mails.DriverRegistration {
	name := inst.name()
	r1 := &mails.DriverRegistration{
		Name:    name,
		Enabled: true,
		Driver:  inst,
	}
	return []*mails.DriverRegistration{r1}
}

func (inst *MockDriver) Accept(cfg *mails.Configuration) bool {
	return cfg.DriverName == inst.name()
}

func (inst *MockDriver) CreateDispatcher(cfg *mails.Configuration) (mails.Dispatcher, error) {
	d := &myMockDriverDispatcher{}
	d.config = *cfg
	return d, nil
}

////////////////////////////////////////////////////////////////////////////////

type myMockDriverDispatcher struct {
	config mails.Configuration
}

func (inst *myMockDriverDispatcher) _impl() mails.Dispatcher {
	return inst
}

func (inst *myMockDriverDispatcher) Accept(c context.Context, msg *mails.Message) bool {
	addr1 := msg.FromAddress
	addr2 := inst.config.FromAddress
	return addr1 == addr2
}

func (inst *myMockDriverDispatcher) Send(c context.Context, msg *mails.Message) error {
	vlog.Info("mock: send mail to %s", msg.ToAddress)
	return nil
}
