# nspub

Package nspub provides a [CoreDNS](https://github.com/coredns/coredns/) plugin to publish successful DNS lookups to [NSQ](https://github.com/nsqio/nsq).

To use this plugin, CoreDNS must be compiled with this plugin by adding `nspub:jw4.us/nspub` to the plugins.cfg file, at the desired level.
If in doubt, put it right before the line that has `log:log`.

The plugin is configured in the Corefile, inside the desired definition block.
The topic and address arguments are required.

Example Corefile:

```
  . {
    whoami
    nspub <topic> <address>
  }
```

Where `<topic>` is whatever the NSQ topic name should be, and `<address>` is the NSQ TCP address, like `10.0.0.1:4150`.
