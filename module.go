package moduleemail

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/module-email/gen/gen4email"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/module-email"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 导出模块
func Module() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4email.ExportComForModuleEmail)

	mb.Depend(starter.Module())

	return mb.Create()
}
