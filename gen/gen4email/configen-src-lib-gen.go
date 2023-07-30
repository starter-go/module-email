package gen4email
import (
    p0ef6f2938 "github.com/starter-go/application"
    pecb1d470d "github.com/starter-go/module-email/internal"
    p6a34f6f22 "github.com/starter-go/module-email/mails"
     "github.com/starter-go/application"
)

// type pecb1d470d.MainDispatcherRegistry in package:github.com/starter-go/module-email/internal
//
// id:com-ecb1d470d2f4f72d-internal-MainDispatcherRegistry
// class:class-6a34f6f2249275109e9baea3c805a883-DispatcherRegistry
// alias:
// scope:singleton
//
type pecb1d470d2_internal_MainDispatcherRegistry struct {
}

func (inst* pecb1d470d2_internal_MainDispatcherRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecb1d470d2f4f72d-internal-MainDispatcherRegistry"
	r.Classes = "class-6a34f6f2249275109e9baea3c805a883-DispatcherRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecb1d470d2_internal_MainDispatcherRegistry) new() any {
    return &pecb1d470d.MainDispatcherRegistry{}
}

func (inst* pecb1d470d2_internal_MainDispatcherRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecb1d470d.MainDispatcherRegistry)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)
    com.Drivers = inst.getDrivers(ie)
    com.DispatcherNameList = inst.getDispatcherNameList(ie)


    return nil
}


func (inst*pecb1d470d2_internal_MainDispatcherRegistry) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pecb1d470d2_internal_MainDispatcherRegistry) getDrivers(ie application.InjectionExt)p6a34f6f22.DriverManager{
    return ie.GetComponent("#alias-6a34f6f2249275109e9baea3c805a883-DriverManager").(p6a34f6f22.DriverManager)
}


func (inst*pecb1d470d2_internal_MainDispatcherRegistry) getDispatcherNameList(ie application.InjectionExt)string{
    return ie.GetString("${mails.dispatcher-name-list}")
}



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



// type pecb1d470d.DispatcherManager in package:github.com/starter-go/module-email/internal
//
// id:com-ecb1d470d2f4f72d-internal-DispatcherManager
// class:
// alias:alias-6a34f6f2249275109e9baea3c805a883-Service
// scope:singleton
//
type pecb1d470d2_internal_DispatcherManager struct {
}

func (inst* pecb1d470d2_internal_DispatcherManager) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecb1d470d2f4f72d-internal-DispatcherManager"
	r.Classes = ""
	r.Aliases = "alias-6a34f6f2249275109e9baea3c805a883-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecb1d470d2_internal_DispatcherManager) new() any {
    return &pecb1d470d.DispatcherManager{}
}

func (inst* pecb1d470d2_internal_DispatcherManager) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecb1d470d.DispatcherManager)
	nop(ie, com)

	
    com.Regs = inst.getRegs(ie)
    com.DefaultSenderAddr = inst.getDefaultSenderAddr(ie)


    return nil
}


func (inst*pecb1d470d2_internal_DispatcherManager) getRegs(ie application.InjectionExt)[]p6a34f6f22.DispatcherRegistry{
    dst := make([]p6a34f6f22.DispatcherRegistry, 0)
    src := ie.ListComponents(".class-6a34f6f2249275109e9baea3c805a883-DispatcherRegistry")
    for _, item1 := range src {
        item2 := item1.(p6a34f6f22.DispatcherRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pecb1d470d2_internal_DispatcherManager) getDefaultSenderAddr(ie application.InjectionExt)string{
    return ie.GetString("${mails.default-sender-address}")
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



// type pecb1d470d.SMSDriver in package:github.com/starter-go/module-email/internal
//
// id:com-ecb1d470d2f4f72d-internal-SMSDriver
// class:class-6a34f6f2249275109e9baea3c805a883-DriverRegistry
// alias:
// scope:singleton
//
type pecb1d470d2_internal_SMSDriver struct {
}

func (inst* pecb1d470d2_internal_SMSDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecb1d470d2f4f72d-internal-SMSDriver"
	r.Classes = "class-6a34f6f2249275109e9baea3c805a883-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecb1d470d2_internal_SMSDriver) new() any {
    return &pecb1d470d.SMSDriver{}
}

func (inst* pecb1d470d2_internal_SMSDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecb1d470d.SMSDriver)
	nop(ie, com)

	


    return nil
}



// type pecb1d470d.SMTPSenderDriver in package:github.com/starter-go/module-email/internal
//
// id:com-ecb1d470d2f4f72d-internal-SMTPSenderDriver
// class:class-6a34f6f2249275109e9baea3c805a883-DriverRegistry
// alias:
// scope:singleton
//
type pecb1d470d2_internal_SMTPSenderDriver struct {
}

func (inst* pecb1d470d2_internal_SMTPSenderDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ecb1d470d2f4f72d-internal-SMTPSenderDriver"
	r.Classes = "class-6a34f6f2249275109e9baea3c805a883-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pecb1d470d2_internal_SMTPSenderDriver) new() any {
    return &pecb1d470d.SMTPSenderDriver{}
}

func (inst* pecb1d470d2_internal_SMTPSenderDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pecb1d470d.SMTPSenderDriver)
	nop(ie, com)

	


    return nil
}


