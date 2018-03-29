/*
Package nspub provides a CoreDNS plugin to publish successful DNS
lookups to NSQ.

To use this plugin, CoreDNS must be compiled with this plugin by adding
'nspub:jw4.us/nspub' to the plugins.cfg file, at the desired level. If
in doubt, put it right before the line 'log:log'

The plugin is configured in the Corefile, inside the desired definition
block. The topic and address arguments are required.

For example:

    . {
      whoami
      nspub <topic> <address>
    }

Where <topic> is whatever the NSQ topic name should be, and <address> is
the NSQ TCP address, like '10.0.0.1:4150'
*/
package nspub

import (
	"context"
	"log"
	"net"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

// CoreDNSPluginName is the canonical name for this plugin.
const CoreDNSPluginName = "nspub"

type publisher struct {
	next plugin.Handler
	cfg  *config
}

func (p *publisher) Name() string { return CoreDNSPluginName }
func (p *publisher) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	cw := &copyWriter{inner: w}
	rcode, err := plugin.NextOrFailure(p.Name(), p.next, ctx, cw, r)
	switch rcode {
	case dns.RcodeSuccess:
		if cw.msg != nil {
			if err = p.publish(w.RemoteAddr().String(), cw.msg); err != nil {
				log.Printf("error publishing to nsq")
			}
		}
	}
	return rcode, err
}

func (p *publisher) publish(clientAddress string, msg *dns.Msg) error {
	prod, err := p.cfg.producer()
	switch err {
	case nil: // no error
	case errNoNSQConfig: // unconfigured publisher
		return nil
	default:
		return err
	}

	data, err := msg.Pack()
	if err != nil {
		return err
	}
	host, _, err := net.SplitHostPort(clientAddress)
	if err != nil {
		return err
	}
	send := &Message{ClientIP: net.ParseIP(host), Data: data}
	return prod.PublishAsync(p.cfg.topic, send.Pack(), nil)
}
