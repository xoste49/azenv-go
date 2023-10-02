# azenv-go
Entire application used to check proxy server statuses.

## Installation
#### Docker
`docker run --privileged --ulimit nofile=65536:65536 -d --name azenv-go --restart unless-stopped -p 8080:8080 xoste49/azenv-go:latest`

#### or Systemd

```bash
# Install Go
wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
# Get azenv-go
git clone https://github.com/xoste49/azenv-go
cd azenv-go
# build
CGO_ENABLED=0 go build main.go
cp ./main /usr/local/bin/azenv-go
# install service
cp azenv-go.service /etc/systemd/system/azenv-go.service
systemctl daemon-reload
systemctl enable azenv-go
systemctl start azenv-go
systemctl status azenv-go
```

## Benchmarks

`h2load --h1 -n100000 -c10000 -t6 http://127.0.0.1`

Go
`finished in 699.62ms, 142934.33 req/s, 25.49MB/s`

Go in Docker
`finished in 1.17s, 85318.26 req/s, 15.30MB/s`

PHP Nginx in Docker
`finished in 10.94s, 9142.11 req/s, 0B/s`

Python Flask in Docker
`finished in 33.36s, 2879.40 req/s, 2.31MB/s`

Python Fastapi in Docker
`finished in 23.23s, 4304.22 req/s, 899.51KB/s`

---
`h2load --h1 -n1000000 -c1000 -t6 http://127.0.0.1`

Go 
`finished in 4.28s, 233496.52 req/s, 41.64MB/s`


## Scaling

ulimit -n

now shows 65536 instead of 1024 = and no problems now

- https://djangoadventures.com/how-to-increase-the-open-files-limit-on-ubuntu/