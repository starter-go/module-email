package internal

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/module-email/mails"
)

// DriverManagerImpl ...
type DriverManagerImpl struct {

	//starter:component
	_as func(mails.DriverManager) //starter:as("#")

	Regs []mails.DriverRegistry //starter:inject(".")

	cached []mails.Driver
}

func (inst *DriverManagerImpl) _impl() mails.DriverManager {
	return inst
}

// Life ...
func (inst *DriverManagerImpl) Life() *application.Life {
	return &application.Life{OnCreate: inst.init}
}

func (inst *DriverManagerImpl) init() error {
	_, err := inst.tryGetDriverList()
	return err
}

func (inst *DriverManagerImpl) getDriverList() []mails.Driver {
	list, err := inst.tryGetDriverList()
	if err != nil {
		panic(err)
	}
	return list
}

func (inst *DriverManagerImpl) tryGetDriverList() ([]mails.Driver, error) {
	list := inst.cached
	if list == nil {
		li, err := inst.loadDriverList()
		if err != nil {
			return nil, err
		}
		list = li
		inst.cached = li
	}
	return list, nil
}

func (inst *DriverManagerImpl) loadDriverList() ([]mails.Driver, error) {
	src := inst.Regs
	dst := make([]mails.Driver, 0)
	for _, r1 := range src {
		list1 := r1.ListRegistrations()
		for _, r2 := range list1 {
			d := inst.getDriver(r2)
			if d == nil {
				continue
			}
			dst = append(dst, d)
		}
	}
	return dst, nil
}

func (inst *DriverManagerImpl) getDriver(r *mails.DriverRegistration) mails.Driver {
	if r == nil {
		return nil
	}
	if !r.Enabled {
		return nil
	}
	return r.Driver
}

// FindDriver ...
func (inst *DriverManagerImpl) FindDriver(cfg *mails.Configuration) (mails.Driver, error) {
	all := inst.getDriverList()
	for _, drv := range all {
		if drv.Accept(cfg) {
			return drv, nil
		}
	}
	return nil, fmt.Errorf("no mails.Driver with name: %s", cfg.Driver)
}

// CreateDispatcher ...
func (inst *DriverManagerImpl) CreateDispatcher(cfg *mails.Configuration) (mails.Dispatcher, error) {
	driver, err := inst.FindDriver(cfg)
	if err != nil {
		return nil, err
	}
	return driver.CreateDispatcher(cfg)
}
