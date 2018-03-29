package nspub

import (
	"net"

	"github.com/miekg/dns"
)

type copyWriter struct {
	inner dns.ResponseWriter
	msg   *dns.Msg
}

func (w *copyWriter) WriteMsg(m *dns.Msg) error {
	w.msg = m.Copy()
	return w.inner.WriteMsg(m)
}

func (w *copyWriter) LocalAddr() net.Addr         { return w.inner.LocalAddr() }
func (w *copyWriter) RemoteAddr() net.Addr        { return w.inner.RemoteAddr() }
func (w *copyWriter) Write(d []byte) (int, error) { return w.inner.Write(d) }
func (w *copyWriter) Close() error                { return w.inner.Close() }
func (w *copyWriter) TsigStatus() error           { return w.inner.TsigStatus() }
func (w *copyWriter) TsigTimersOnly(b bool)       { w.inner.TsigTimersOnly(b) }
func (w *copyWriter) Hijack()                     { w.inner.Hijack() }
