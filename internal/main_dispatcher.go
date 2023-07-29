package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/module-email/mails"
)

// MainDispatcher ...
type MainDispatcher struct {
	//starter:component
	_as func(mails.Dispatcher, mails.DispatcherRegistry) //starter:as(".",".")

	AppContext         application.Context //starter:inject("context")
	Drivers            mails.DriverManager //starter:inject("#")
	DispatcherNameList string              //starter:inject("${mails.dispatcher-name-list}")

	dispatcherList []mails.Dispatcher
}

func (inst *MainDispatcher) _impl() (mails.Dispatcher, mails.DispatcherRegistry, application.Lifecycle) {
	return inst, inst, inst
}

func (inst *MainDispatcher) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *MainDispatcher) init() error {
	_, err := inst.tryGetDispatcherList()
	return err
}

func (inst *MainDispatcher) ListRegistrations() []*mails.DispatcherRegistration {
	r1 := &mails.DispatcherRegistration{
		Name:       "main",
		Enabled:    true,
		Dispatcher: inst,
	}
	return []*mails.DispatcherRegistration{r1}
}

func (inst *MainDispatcher) Accept(c context.Context, msg *mails.Message) bool {
	return true
}

func (inst *MainDispatcher) Send(c context.Context, msg *mails.Message) error {
	if msg == nil {
		return fmt.Errorf("mails.Message is nil")
	}
	list := inst.getDispatcherList()
	for _, item := range list {
		if item.Accept(c, msg) {
			return item.Send(c, msg)
		}
	}
	return fmt.Errorf("no dispatcher accept the mails.Message from %s", msg.FromAddress)
}

func (inst *MainDispatcher) getDispatcherList() []mails.Dispatcher {
	list, err := inst.tryGetDispatcherList()
	if err != nil {
		panic(err)
	}
	return list
}

func (inst *MainDispatcher) tryGetDispatcherList() ([]mails.Dispatcher, error) {
	list := inst.dispatcherList
	if list == nil {
		li, err := inst.loadDispatcherList()
		if err != nil {
			return nil, err
		}
		list = li
		inst.dispatcherList = li
	}
	return list, nil
}

func (inst *MainDispatcher) loadDispatcher(name string) (mails.Dispatcher, error) {

	prefix := "mails." + name + "."
	props := inst.AppContext.GetProperties()
	g := &pGetter{props: props}
	cfg := &mails.Configuration{}

	cfg.DriverName = g.getString(prefix + "driver")
	cfg.FromAddress = g.getAddress(prefix + "from-address")
	cfg.FromUser = g.getString(prefix + "from-user")

	if g.err != nil {
		return nil, g.err
	}
	return inst.Drivers.CreateDispatcher(cfg)
}

func (inst *MainDispatcher) loadDispatcherList() ([]mails.Dispatcher, error) {
	dst := make([]mails.Dispatcher, 0)
	namelist := inst.getDispatcherNameList()
	for _, name := range namelist {
		d, err := inst.loadDispatcher(name)
		if err != nil {
			return nil, err
		}
		dst = append(dst, d)
	}
	return dst, nil
}

func (inst *MainDispatcher) getDispatcherNameList() []string {
	const sep = ","
	text := inst.DispatcherNameList
	src := strings.Split(text, sep)
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type pGetter struct {
	props properties.Table
	err   error
}

func (inst *pGetter) getString(name string) string {
	value, err := inst.props.GetPropertyRequired(name)
	if err != nil {
		inst.err = err
	}
	return value
}

func (inst *pGetter) getAddress(name string) mails.Address {
	value := inst.getString(name)
	addr, err := inst.parseAddress(value)
	if err != nil && inst.err == nil {
		inst.err = err
	}
	return addr
}

func (inst *pGetter) parseAddress(str string) (mails.Address, error) {
	parts := strings.Split(str, "@")
	if len(parts) == 2 {
		user1 := parts[0]
		host1 := parts[1]
		user2 := strings.TrimSpace(user1)
		host2 := strings.TrimSpace(host1)
		if user1 == user2 && host1 == host2 && user1 != "" && host1 != "" {
			return mails.Address(str), nil
		}
	}
	return "", fmt.Errorf("bad email address: %s", str)
}
