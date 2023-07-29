package mails

import "context"

// Sender 是简单而且抽象的邮件发送接口
type Sender interface {

	// 发邮件
	Send(c context.Context, msg *Message) error
}

// Dispatcher 是跟具体发送地址绑定的邮件发送接口
type Dispatcher interface {
	Sender

	// 判断是否支持发送指定的邮件
	Accept(c context.Context, msg *Message) bool
}

// DispatcherRegistration ...
type DispatcherRegistration struct {
	Enabled    bool
	Priority   int
	Name       string
	Dispatcher Dispatcher
}

// DispatcherRegistry ...
type DispatcherRegistry interface {
	ListRegistrations() []*DispatcherRegistration
}
