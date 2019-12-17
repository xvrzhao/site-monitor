# Site Monitor
Availability Monitor for websites or HTTP services.

## Features
- Support multiple parameter settings to meet monitoring requirements
- Real-time status information
- Email notification when the website is unhealthy

## Installation
```shell script
$ go get github.com/xvrzhao/site-monitor
```

## Usage

```shell script
$ site-monitor --monitor-url=http://example.com --monitor-interval=3000 \
--mail-smtp-auth-host=smtp.example.com --mail-smtp-server-addr=smtp.example.com:25 \
--mail-from-addr=site@example.com --mail-from-name=site-monitor --mail-from-password=password \
--mail-to-addr=xvrzhao@gmail.com \
--debug-mode
```

Run `site-monitor --help` for more usage information.