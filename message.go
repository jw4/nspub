package nspub

import (
	"net"
	"strings"

	"github.com/miekg/dns"
)

// Message represents a client ip address and some data to be sent in
// a binary format.
type Message struct {
	ClientIP net.IP
	Data     []byte
}

func (m *Message) String() string {
	b := &strings.Builder{}
	b.WriteString(m.ClientIP.String())
	b.WriteString("|")

	msg := &dns.Msg{}
	if err := msg.Unpack(m.Data); err != nil {
		return err.Error()
	}

	delim := ""
	for _, q := range msg.Question {
		b.WriteString(delim + q.Name)
		delim = ","
	}
	if len(msg.Question) > 0 {
		b.WriteString("|")
		delim = ""
	}
	for _, a := range msg.Answer {
		b.WriteString(delim + a.String())
		delim = ","
	}
	return b.String()
}

// Pack produces a []byte containaing in the first byte the length of
// ClientIP, then the []byte representation of ClientIP, then the rest
// of the slice contains the Data.
func (m *Message) Pack() []byte {
	data := make([]byte, 1+len(m.ClientIP)+len(m.Data))
	ip := m.ClientIP
	ipsize := len(m.ClientIP)
	data[0] = byte(ipsize)
	copy(data[1:1+ipsize], ip)
	copy(data[1+ipsize:], m.Data)
	return data
}

// Unpack populates the message from the passed in []byte.
func (m *Message) Unpack(data []byte) {
	length := len(data)
	if length < 1 {
		return
	}
	ipsize := int(data[0])
	if length < 1+ipsize {
		return
	}
	ip := make([]byte, ipsize)
	copy(ip, data[1:1+ipsize])
	m.ClientIP = net.IP(ip)
	m.Data = make([]byte, length-ipsize-1)
	copy(m.Data, data[1+ipsize:])
}
