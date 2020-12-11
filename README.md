# vmalertcli

A CLI application to manage VictoriaMetrics vmalert.

All actions return JSON objects (prettyprint TBA), so usage of `jq` is good for formatting output.

### Help menu

```
vmalertcli -h
Usage of vmalertcli:
  -action="groups": VMAlert action to take {groups|alerts|metrics|reload}
  -host="localhost": Host where VMAlert responds
  -port=8880: VMAlert port
```

### Installation 
 
Put binary in $PATH (release will be updated).

### Usage

View alert groups:
```
vmalertcli -action groups | jq
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