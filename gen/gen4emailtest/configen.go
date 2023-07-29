package gen4emailtest

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForModuleEmailTest ...
func ExportComForModuleEmailTest(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
