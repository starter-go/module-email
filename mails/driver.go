package mails

// Configuration ...
type Configuration struct {
	FromAddress Address
	FromUser    string
	DriverName  string
}

// Driver 表示邮件发送驱动
type Driver interface {
	Accept(cfg *Configuration) bool
	CreateDispatcher(cfg *Configuration) (Dispatcher, error)
}

// DriverRegistration ...
type DriverRegistration struct {
	Name    string
	Enabled bool
	Driver  Driver
}

// DriverRegistry ...
type DriverRegistry interface {
	ListRegistrations() []*DriverRegistration
}

// DriverManager 驱动管理...
type DriverManager interface {
	FindDriver(cfg *Configuration) (Driver, error)
	CreateDispatcher(cfg *Configuration) (Dispatcher, error)
}
