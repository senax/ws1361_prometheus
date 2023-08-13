# ws1361_prometheus
Small program to present WS1361 sound pressure meter reading on a prometheus endpoint.

# rasbian
```
apt-get install git libusb-1.0
# golang in rasbian is very old.
wget https://go.dev/dl/go1.21.0.linux-armv6l.tar.gz
tar -C /usr/local/ -xzf go1.21*
export PATH=$PATH:/usr/local/go/bin

mkdir -p /root/go/src/github.com/senax
cd !$
git clone https://github.com/senax/ws1361_prometheus.git
cd ws1361_prometheus/
go get
go build -ldflags="-s -w"
```

# Example output
```
$ curl -s 192.168.128.217:1361/metrics |grep ws1361
# HELP ws1361_decibel sound pressure
# TYPE ws1361_decibel gauge
ws1361_decibel 47.8
# HELP ws1361_fast Fast: 1=fast 0=slow
# TYPE ws1361_fast gauge
ws1361_fast 1
# HELP ws1361_max Max: 1=on 0=off
# TYPE ws1361_max gauge
ws1361_max 1
# HELP ws1361_mode Mode: 1=dba 0=dbc
# TYPE ws1361_mode gauge
ws1361_mode 1
# HELP ws1361_range Range: 0: 30-80, 1: 40-90, 2: 50-100, 3: 60-110, 4: 70-120, 5: 80-130, 7: 30-130
# TYPE ws1361_range gauge
ws1361_range 0
```
