package nspub

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin(CoreDNSPluginName, caddy.Plugin{ServerType: "dns", Action: setup})
}

func setup(c *caddy.Controller) error {
	cfg := newConfig(&c.Dispenser)
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler { return &publisher{next: next, cfg: cfg} })
	return nil
}
