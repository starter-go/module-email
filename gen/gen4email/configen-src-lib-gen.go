package gen4email
import (
    p0ef6f2938 "github.com/starter-go/application"
    pecb1d470d "github.com/starter-go/module-email/internal"
    p6a34f6f22 "github.com/starter-go/module-email/mails"
     "github.com/starter-go/application"
)

// type pecb1d470d.DriverManagerImpl in package:github.com/starter-go/module-email/internal
//
// id:com-ecb1d470d2f4f72d-internal-DriverManagerImpl
// class:
// alias:alias-6a34f6f2249275109e9baea3c805a883-DriverManager
// scope:singleton
//
type pecb1d470d2_internal_DriverManagerImpl struct {
}

func (inst* pecb1d470d2_internal_DriverManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecb1d470d2f4f72d-internal-DriverManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-6a34f6f2249275109e9baea3c805a883-DriverManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecb1d470d2_internal_DriverManagerImpl) new() any {
    return &pecb1d470d.DriverManagerImpl{}
}

func (inst* pecb1d470d2_internal_DriverManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecb1d470d.DriverManagerImpl)
	nop(ie, com)

	
    com.Regs = inst.getRegs(ie)


    return nil
}


func (inst*pecb1d470d2_internal_DriverManagerImpl) getRegs(ie application.InjectionExt)[]p6a34f6f22.DriverRegistry{
    dst := make([]p6a34f6f22.DriverRegistry, 0)
    src := ie.ListComponents(".class-6a34f6f2249275109e9baea3c805a883-DriverRegistry")
    for _, item1 := range src {
        item2 := item1.(p6a34f6f22.DriverRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type pecb1d470d.MainDispatcher in package:github.com/starter-go/module-email/internal
//
// id:com-ecb1d470d2f4f72d-internal-MainDispatcher
// class:class-6a34f6f2249275109e9baea3c805a883-Dispatcher class-6a34f6f2249275109e9baea3c805a883-DispatcherRegistry
// alias:
// scope:singleton
//
type pecb1d470d2_internal_MainDispatcher struct {
}

func (inst* pecb1d470d2_internal_MainDispatcher) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecb1d470d2f4f72d-internal-MainDispatcher"
	r.Classes = "class-6a34f6f2249275109e9baea3c805a883-Dispatcher class-6a34f6f2249275109e9baea3c805a883-DispatcherRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecb1d470d2_internal_MainDispatcher) new() any {
    return &pecb1d470d.MainDispatcher{}
}

func (inst* pecb1d470d2_internal_MainDispatcher) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecb1d470d.MainDispatcher)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)
    com.Drivers = inst.getDrivers(ie)
    com.DispatcherNameList = inst.getDispatcherNameList(ie)


    return nil
}


func (inst*pecb1d470d2_internal_MainDispatcher) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pecb1d470d2_internal_MainDispatcher) getDrivers(ie application.InjectionExt)p6a34f6f22.DriverManager{
    return ie.GetComponent("#alias-6a34f6f2249275109e9baea3c805a883-DriverManager").(p6a34f6f22.DriverManager)
}


func (inst*pecb1d470d2_internal_MainDispatcher) getDispatcherNameList(ie application.InjectionExt)string{
    return ie.GetString("${mails.dispatcher-name-list}")
}



// type pecb1d470d.MockDriver in package:github.com/starter-go/module-email/internal
//
// id:com-ecb1d470d2f4f72d-internal-MockDriver
// class:class-6a34f6f2249275109e9baea3c805a883-DriverRegistry
// alias:
// scope:singleton
//
type pecb1d470d2_internal_MockDriver struct {
}

func (inst* pecb1d470d2_internal_MockDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecb1d470d2f4f72d-internal-MockDriver"
	r.Classes = "class-6a34f6f2249275109e9baea3c805a883-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecb1d470d2_internal_MockDriver) new() any {
    return &pecb1d470d.MockDriver{}
}

func (inst* pecb1d470d2_internal_MockDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecb1d470d.MockDriver)
	nop(ie, com)

	


    return nil
}



// type pecb1d470d.SenderService in package:github.com/starter-go/module-email/internal
//
// id:com-ecb1d470d2f4f72d-internal-SenderService
// class:
// alias:alias-6a34f6f2249275109e9baea3c805a883-Service
// scope:singleton
//
type pecb1d470d2_internal_SenderService struct {
}

func (inst* pecb1d470d2_internal_SenderService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecb1d470d2f4f72d-internal-SenderService"
	r.Classes = ""
	r.Aliases = "alias-6a34f6f2249275109e9baea3c805a883-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecb1d470d2_internal_SenderService) new() any {
    return &pecb1d470d.SenderService{}
}

func (inst* pecb1d470d2_internal_SenderService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecb1d470d.SenderService)
	nop(ie, com)

	
    com.Regs = inst.getRegs(ie)
    com.DefaultSenderAddr = inst.getDefaultSenderAddr(ie)


    return nil
}


func (inst*pecb1d470d2_internal_SenderService) getRegs(ie application.InjectionExt)[]p6a34f6f22.DispatcherRegistry{
    dst := make([]p6a34f6f22.DispatcherRegistry, 0)
    src := ie.ListComponents(".class-6a34f6f2249275109e9baea3c805a883-DispatcherRegistry")
    for _, item1 := range src {
        item2 := item1.(p6a34f6f22.DispatcherRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pecb1d470d2_internal_SenderService) getDefaultSenderAddr(ie application.InjectionExt)string{
    return ie.GetString("${mails.default-sender-address}")
}


