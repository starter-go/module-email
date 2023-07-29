package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	moduleemail "github.com/starter-go/module-email"
	"github.com/starter-go/starter"
)

//go:embed "resources"
var theModuleResFS embed.FS

func module() application.Module {

	parent := moduleemail.Module()
	mb := &application.ModuleBuilder{}

	mb.Name(parent.Name())
	mb.Version(parent.Version())
	mb.Revision(parent.Revision())

	mb.EmbedResources(theModuleResFS, "resources")

	mb.Components(nil)

	mb.Depend(parent)
	return mb.Create()
}

func main() {
	i := starter.Init(os.Args)
	i.MainModule(module())
	i.WithPanic(true).Run()
}
