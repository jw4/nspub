package nspub

import (
	"encoding/json"
	"net"
	"time"

	"github.com/miekg/dns"
)

// Message represents a client ip address and some data to be sent in
// a binary format.
type Message struct {
	ClientIP net.IP
	Time     time.Time
	Msg      *dns.Msg
}

func (m *Message) String() string {
	data, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
