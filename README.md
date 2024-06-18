[中文](README_ZH.md)

# H UI

Just the panel for Hysteria2

# Recommended OS

OS version: CentOS Stream 8/Ubuntu 20.04 LTS/Debian 11.1

CPU Architecture: x86_64/amd64

# Install

- Docker(recommended)

```bash
docker run -d --name hui --restart always \
  --network=host \
  -v /hui/bin:/bin \
  -v /hui/data:/data \
  -v /hui/export:/export \
  -v /hui/logs:/logs \
  hui

```

Custom web port

```bash
docker run -d --name hui --restart always \
  --network=host \
  -v /hui/bin:/bin \
  -v /hui/data:/data \
  -v /hui/export:/export \
  -v /hui/logs:/logs \
  hui \
  ./hui -p [端口]
```

# compile

```shell
go build -trimpath -ldflags -o h-ui -ldflags="-s -w"
```

# Other