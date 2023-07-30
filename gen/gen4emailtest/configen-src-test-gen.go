package gen4emailtest
import (
    p6a34f6f22 "github.com/starter-go/module-email/mails"
    pcedd612ad "github.com/starter-go/module-email/src/test/code"
     "github.com/starter-go/application"
)

// type pcedd612ad.Test1 in package:github.com/starter-go/module-email/src/test/code
//
// id:com-cedd612adf6a490b-code-Test1
// class:
// alias:
// scope:singleton
//
type pcedd612adf_code_Test1 struct {
}

func (inst* pcedd612adf_code_Test1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-cedd612adf6a490b-code-Test1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pcedd612adf_code_Test1) new() any {
    return &pcedd612ad.Test1{}
}

func (inst* pcedd612adf_code_Test1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pcedd612ad.Test1)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.ToAddr = inst.getToAddr(ie)


    return nil
}


func (inst*pcedd612adf_code_Test1) getSender(ie application.InjectionExt)p6a34f6f22.Service{
    return ie.GetComponent("#alias-6a34f6f2249275109e9baea3c805a883-Service").(p6a34f6f22.Service)
}


func (inst*pcedd612adf_code_Test1) getToAddr(ie application.InjectionExt)string{
    return ie.GetString("${mails.test.to-addr}")
}


