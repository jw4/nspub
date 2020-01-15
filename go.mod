module jw4.us/nspub

require (
	github.com/caddyserver/caddy v1.0.4
	github.com/coredns/coredns v1.6.6
	github.com/miekg/dns v1.1.25
	github.com/nsqio/go-nsq v1.0.7
)

replace golang.org/x/net v0.0.0-20190813000000-74dc4d7220e7 => golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297

go 1.13
