# Myip

Myip is a simple command tool to obtain the users public IP address using the ipify.org API

### Prerequisites

A system with Go installed

### Installing

Compile the tool as follows

```
go build myip.go
```

Run the tool from current folder or copy to location in your path, such as /usr/local/bin/

### Usage
```
$ myip
20.30.40.50

$ myip -output=json
{"ip":"82.23.60.241"}

$ myip -version
myip, version:0.1
```