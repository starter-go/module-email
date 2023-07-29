package gen4email

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&pecb1d470d2_internal_DriverManagerImpl{})
    inst.register(&pecb1d470d2_internal_MainDispatcher{})
    inst.register(&pecb1d470d2_internal_MockDriver{})
    inst.register(&pecb1d470d2_internal_SenderService{})


    return nil
}
