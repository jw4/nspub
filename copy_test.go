package nspub

import "github.com/miekg/dns"

var _ dns.ResponseWriter = (*copyWriter)(nil)
