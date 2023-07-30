package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/module-email/mails"
)

// MainDispatcherRegistry ...
type MainDispatcherRegistry struct {
	//starter:component
	_as func(mails.DispatcherRegistry) //starter:as(".")

	AppContext         application.Context //starter:inject("context")
	Drivers            mails.DriverManager //starter:inject("#")
	DispatcherNameList string              //starter:inject("${mails.dispatcher-name-list}")

	dispatcherList []*mails.DispatcherRegistration // the cache
}

func (inst *MainDispatcherRegistry) _impl() (mails.DispatcherRegistry, application.Lifecycle) {
	return inst, inst
}

// Life ...
func (inst *MainDispatcherRegistry) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *MainDispatcherRegistry) init() error {
	_, err := inst.tryGetDispatcherList()
	return err
}

// ListRegistrations ...
func (inst *MainDispatcherRegistry) ListRegistrations() []*mails.DispatcherRegistration {
	return inst.getDispatcherList()
}

func (inst *MainDispatcherRegistry) getDispatcherList() []*mails.DispatcherRegistration {
	list, err := inst.tryGetDispatcherList()
	if err != nil {
		panic(err)
	}
	return list
}

func (inst *MainDispatcherRegistry) tryGetDispatcherList() ([]*mails.DispatcherRegistration, error) {
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

func (inst *MainDispatcherRegistry) loadDispatcherList() ([]*mails.DispatcherRegistration, error) {
	dst := make([]*mails.DispatcherRegistration, 0)
	namelist := inst.getDispatcherNameList()
	for _, name := range namelist {
		d, err := inst.loadDispatcher(name)
		if err != nil {
			return nil, err
		}
		if d.Enabled {
			dst = append(dst, d)
		}
	}
	return dst, nil
}

func (inst *MainDispatcherRegistry) getDispatcherNameList() []string {
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

func (inst *MainDispatcherRegistry) loadDispatcher(name string) (*mails.DispatcherRegistration, error) {

	prefix := "mails." + name + "."
	props := inst.AppContext.GetProperties()
	g := &pGetter{props: props}
	cfg := &mails.Configuration{}

	cfg.Driver = g.getString(prefix + "driver")
	cfg.WorkAround = g.getBool(prefix + "workaround")
	cfg.Enabled = g.getBool(prefix + "enabled")
	cfg.Priority = g.getInt(prefix + "priority")

	cfg.SenderAddress = g.getAddress(prefix + "sender-address")
	cfg.SenderName = g.getString(prefix + "sender-name")

	cfg.Host = g.getString(prefix + "host")
	cfg.Port = g.getInt(prefix + "port")
	cfg.UserName = g.getString(prefix + "username")
	cfg.Password = g.getString(prefix + "password")

	if g.err != nil {
		return nil, g.err
	}
	disp, err := inst.Drivers.CreateDispatcher(cfg)
	if err != nil {
		return nil, err
	}
	reg := &mails.DispatcherRegistration{
		Name:       name,
		Enabled:    cfg.Enabled,
		Priority:   cfg.Priority,
		Dispatcher: disp,
	}
	return reg, nil
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

func (inst *pGetter) getInt(name string) int {
	str := inst.getString(name)
	n, err := strconv.Atoi(str)
	if err != nil && inst.err == nil {
		inst.err = err
	}
	return n
}

func (inst *pGetter) getBool(name string) bool {
	str := inst.getString(name)
	return (str == "1") || (str == "true") || (str == "yes")
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
