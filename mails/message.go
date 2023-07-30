package mails

// Address 邮件地址, 格式为 'user@host.domain'
type Address string

// Message 表示一封邮件
type Message struct {
	Title string

	FromAddress Address
	FromUser    string

	ToAddresses []Address
	ToUser      string

	ContentType string
	Content     []byte
}

////////////////////////////////////////////////////////////////////////////////

func (addr Address) String() string {
	return string(addr)
}
