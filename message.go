package nspub

import (
	"encoding/json"
	"net"
	"time"

	"github.com/miekg/dns"
)

// Header is the simplified structure for deserialization.
type Header struct {
	Name string
}

// Question is the simplified structure for deserialization.
type Question struct {
	Name string
}

// Answer is the simplified structure for deserialization.
type Answer struct {
	Hdr  Header
	A    string
	AAAA string
}

// Msg is the simplified structure for deserialization.
type Msg struct {
	ID       int
	Question []Question
	Answer   []Answer
}

// Message is the simplified structure for deserialization.
type Message struct {
	ClientIP string
	Time     time.Time
	Msg      Msg
}

// SerializeMessage is the structure serialized to send.
type SerializeMessage struct {
	ClientIP net.IP
	Time     time.Time
	Msg      *dns.Msg
}

func (m *SerializeMessage) String() string {
	data, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
