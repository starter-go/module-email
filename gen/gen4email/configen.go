package gen4email

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForModuleEmail ...
func ExportComForModuleEmail(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
