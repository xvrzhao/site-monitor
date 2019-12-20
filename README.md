# Site Monitor
Availability Monitor for websites or HTTP services.

## Features
- Support multiple flag settings to meet monitoring requirements
- Real-time website status information
- Email notification when the website is unhealthy

## Installation
```shell script
$ go get github.com/xvrzhao/site-monitor
```

## Usage
```shell script
$ site-monitor -url=http://example.com -cycle=3000 -mail-faddr=sender@qq.com -mail-fpwd=senderpassword -mail-taddr=recipient@gmail.com -mail-auth=smtp.qq.com -mail-server=smtp.qq.com:25 -debug
```
Run `site-monitor -h` for more usage information.