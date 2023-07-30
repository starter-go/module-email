package mails

// Configuration ...
type Configuration struct {
	Driver string // name of driver

	WorkAround bool
	Enabled    bool
	Priority   int

	SenderAddress Address
	SenderName    string

	Host     string
	Port     int
	UserName string
	Password string
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
