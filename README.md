[![amd64 binary](https://github.com/aorfanos/vmalert-cli/actions/workflows/build-bin.yml/badge.svg)](https://github.com/aorfanos/vmalert-cli/actions/workflows/build-bin.yml)
# vmalert-cli

A CLI application to manage VictoriaMetrics vmalert.

All actions return JSON objects which can be pretty-printed using the `-pretty` flag (`metrics` action cannot be pretty-printed).

Host/port are by default `localhost:8880`, you can change the configuration by using the `-host` and `-port` flags.

### Help menu

```console
Usage of ./vmalert-cli:
  -action="groups": VMAlert action to take {groups|alerts|metrics|reload|status <alertName>}
  -host="localhost": Host where VMAlert responds
  -port=8880: VMAlert port
  -pretty=false: Pretty print {false|true}
  -schema="http": Use http|https
```

### Installation 
 
Quick install (example for linux/amd64, change it for your arch):

```console
curl --silent https://api.github.com/repos/aorfanos/vmalert-cli/releases/latest \
  | grep browser_download_url \
  | grep amd64 \
  | grep linux \
  | cut -d '"' -f 4 \
  | wget -qi - && \
  mv vmalertcli-linux-amd64-latest /usr/local/bin/vmalert-cli
```

### Usage

View alert groups:
```
vmalert-cli -schema http -action groups -pretty
```

View VMAlert metrics:
```
vmalert-cli -action metrics
```

Hot-reload VMAlert configuration:
```
vmalert-cli -action reload
```

View active (firing) alerts:
```
vmalert-cli -action alerts
```

View status for specific alert:
```
vmalert-cli -pretty -action status KubeDeploymentFailed
```
