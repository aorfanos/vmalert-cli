# vmalertcli
![Release binary](https://github.com/aorfanos/vmalert-cli/workflows/Release%20binary/badge.svg)


A CLI application to manage VictoriaMetrics vmalert.

All actions return JSON objects which can be pretty-printed using the `-pretty` flag (`metrics` action cannot be pretty-printed).

Host/port are by default `localhost:8880`, you can change the configuration by using the `-host` and `-port` flags.

### Help menu

```
Usage of ./vmalertcli:
  -action="groups": VMAlert action to take {groups|alerts|metrics|reload}
  -host="localhost": Host where VMAlert responds
  -schema="http": Schema to use when accessing VMAlert
  -port=8880: VMAlert port
  -pretty=false: Pretty print {false|true}
```

### Installation 
 
Put binary in $PATH (release will be updated).

### Usage

View alert groups:
```
vmalertcli -schema http -action groups -pretty
```

View VMAlert metrics:
```
vmalertcli -action metrics
```

Hot-reload VMAlert configuration:
```
vmalertcli -action reload
```

View active (firing) alerts:
```
vmalertcli -action alerts
```