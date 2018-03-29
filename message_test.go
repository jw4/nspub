package nspub

import (
	"net"
	"testing"
)

func TestPacking(t *testing.T) {
	for ix, pair := range map[string]struct {
		ip   net.IP
		data string
	}{
		"ipv6":         {ip: net.ParseIP("ffee::9:1"), data: "hello world!"},
		"ipv4":         {ip: net.ParseIP("127.0.0.1"), data: "hello world!"},
		"missing ip":   {data: "it's a great day"},
		"missing data": {ip: net.ParseIP("192.168.199.188")},
	} {
		expect := &Message{ClientIP: pair.ip, Data: []byte(pair.data)}
		got := &Message{}
		got.Unpack(expect.Pack())
		if !expect.ClientIP.Equal(got.ClientIP) {
			t.Errorf("%s: expected ClientIP %q, got %q", ix, expect.ClientIP.String(), got.ClientIP.String())
		}
		if string(expect.Data) != string(got.Data) {
			t.Errorf("%s: expected Data %q, got %q", ix, string(expect.Data), string(got.Data))
		}
	}
}
