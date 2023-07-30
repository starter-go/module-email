package internal

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/starter-go/module-email/mails"
)

// SMTPSenderDriver ...
type SMTPSenderDriver struct {

	//starter:component
	_as func(mails.DriverRegistry) //starter:as(".")

}

func (inst *SMTPSenderDriver) _Impl() (mails.Driver, mails.DriverRegistry) {
	return inst, inst
}

func (inst *SMTPSenderDriver) name() string {
	return "smtp"
}

// ListRegistrations 。。。
func (inst *SMTPSenderDriver) ListRegistrations() []*mails.DriverRegistration {
	name := inst.name()
	r1 := &mails.DriverRegistration{
		Name:    name,
		Enabled: true,
		Driver:  inst,
	}
	return []*mails.DriverRegistration{r1}
}

// Accept ...
func (inst *SMTPSenderDriver) Accept(cfg *mails.Configuration) bool {
	return (cfg.Driver == inst.name())
}

// CreateDispatcher ...
func (inst *SMTPSenderDriver) CreateDispatcher(cfg *mails.Configuration) (mails.Dispatcher, error) {
	if cfg == nil {
		return nil, fmt.Errorf("mails.Configuration is nil")
	}
	sender := &smtpSender{}
	sender.config = *cfg
	return sender, nil
}

////////////////////////////////////////////////////////////////////////////////

// smtpSender 默认的邮件发送组件
type smtpSender struct {
	config mails.Configuration
}

func (inst *smtpSender) _Impl() mails.Dispatcher {
	return inst
}

// Accept ...
func (inst *smtpSender) Accept(c context.Context, msg *mails.Message) bool {
	a1 := msg.FromAddress
	a2 := inst.config.SenderAddress
	return a1 == a2
}

// Send 发送邮件
func (inst *smtpSender) Send(c context.Context, m *mails.Message) error {

	const defaultPort = 25
	cfg := inst.config
	host := cfg.Host
	port := cfg.Port
	from := m.FromAddress.String()
	to := inst.stringListForAddresses(m.ToAddresses)

	if port < 1 {
		port = defaultPort
	}

	hostPort := host + ":" + strconv.Itoa(port)
	auth := inst.prepareAuth(m)
	msg := inst.prepareMessage(m)

	if cfg.WorkAround {
		workaround := smtpSendMailWorkaround{host: host}
		return workaround.SendMail(hostPort, auth, from, to, msg)
	}

	return smtp.SendMail(hostPort, auth, from, to, msg)
}

func (inst *smtpSender) stringListForAddresses(src []mails.Address) []string {
	dst := make([]string, 0)
	for _, addr := range src {
		str := addr.String()
		dst = append(dst, str)
	}
	return dst
}

func (inst *smtpSender) prepareAuth(m *mails.Message) smtp.Auth {
	ident := ""
	cfg := inst.config
	username := cfg.UserName
	password := cfg.Password
	host := cfg.Host
	return smtp.PlainAuth(ident, username, password, host)
}

func (inst *smtpSender) prepareMessage(m *mails.Message) []byte {

	const nl = "\r\n"
	contentType := inst.prepareContentType(m)
	tolist := inst.stringifyToAddrList(m)
	builder := bytes.Buffer{}

	builder.WriteString("To: ")
	builder.WriteString(tolist)
	builder.WriteString(nl)

	builder.WriteString("From: ")
	builder.WriteString(m.FromAddress.String())
	builder.WriteString(nl)

	builder.WriteString("Subject: ")
	builder.WriteString(m.Title)
	builder.WriteString(nl)

	builder.WriteString("Content-Type: ")
	builder.WriteString(contentType)
	builder.WriteString(nl)

	builder.WriteString(nl)
	builder.Write(m.Content)
	return builder.Bytes()
}

func (inst *smtpSender) prepareContentType(m *mails.Message) string {
	t := strings.ToLower(m.ContentType)
	if strings.Contains(t, "html") {
		t = "text/html"
	} else {
		t = "text/plain"
	}
	return t + "; charset=UTF-8"
}

func (inst *smtpSender) stringifyToAddrList(m *mails.Message) string {
	sep := ""
	builder := strings.Builder{}
	list := m.ToAddresses
	for _, addr := range list {
		item := strings.TrimSpace(addr.String())
		if item == "" {
			continue
		}
		builder.WriteString(sep)
		builder.WriteString(item)
		sep = ";"
	}
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////

// HookStartTLS nil, except for tests
type HookStartTLS func(*tls.Config)

type smtpSendMailWorkaround struct {
	testHookStartTLS HookStartTLS
	host             string
}

func (inst *smtpSendMailWorkaround) _Impl() {}

func (inst *smtpSendMailWorkaround) validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("smtp: A line must not contain CR or LF")
	}
	return nil
}

func (inst *smtpSendMailWorkaround) Dial(addr string, cfg *tls.Config) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, cfg)
	if err != nil {
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func (inst *smtpSendMailWorkaround) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {

	if err := inst.validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err := inst.validateLine(recp); err != nil {
			return err
		}
	}
	config := &tls.Config{ServerName: inst.serverName()}
	c, err := inst.Dial(addr, config)
	if err != nil {
		return err
	}
	defer c.Close()

	if a != nil {
		if err = c.Auth(a); err != nil {
			return err
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// -> c.serverName
func (inst *smtpSendMailWorkaround) serverName() string {
	return inst.host
}
