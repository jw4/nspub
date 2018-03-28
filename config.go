package nspub

import (
	"errors"
	"log"

	"github.com/mholt/caddy/caddyfile"
	nsq "github.com/nsqio/go-nsq"
)

var errNoNSQConfig = errors.New("no nsq configuration found")

type config struct {
	topic string
	addr  string
	ncfg  *nsq.Config
	prod  *nsq.Producer
}

func (n *config) producer() (*nsq.Producer, error) {
	if n == nil {
		return nil, errNoNSQConfig
	}
	if n.prod == nil {
		prod, err := nsq.NewProducer(n.addr, n.ncfg)
		if err != nil {
			return nil, err
		}
		n.prod = prod
	}
	return n.prod, nil
}

func newConfig(d *caddyfile.Dispenser) *config {
	args := d.RemainingArgs()
	if len(args) != 3 {
		log.Fatalf("ARGS: %v", args)
		log.Panicf("parse error for plugin %q. <topic> and <nsqd addr> required. file %s, line %d", CoreDNSPluginName, d.File(), d.Line())
	}
	ncfg := nsq.NewConfig()
	ncfg.UserAgent = CoreDNSPluginName
	return &config{
		topic: args[1],
		addr:  args[2],
		ncfg:  ncfg,
	}
}
