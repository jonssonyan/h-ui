[中文](README_ZH.md)

## H UI

Just the panel for Hysteria2

## Recommended OS

OS version: CentOS Stream 8/Ubuntu 20.04 LTS/Debian 11.1

CPU Architecture: x86_64/amd64

## Install

- Docker(Recommended)

```bash
docker run -d --name h-ui --restart always \
  --network=host \
  -v /h-ui/bin:/bin \
  -v /h-ui/data:/data \
  -v /h-ui/export:/export \
  -v /h-ui/logs:/logs \
  jonssonyan/h-ui
```

Custom web port

```bash
docker run -d --name h-ui --restart always \
  --network=host \
  -v /h-ui/bin:/bin \
  -v /h-ui/data:/data \
  -v /h-ui/export:/export \
  -v /h-ui/logs:/logs \
  jonssonyan/h-ui \
  ./h-ui -p [端口]
```

Set the time zone, default is Asia/Shanghai

```bash
docker run -d --name h-ui --restart always \
  --network=host \
  -e TZ=Asia/Shanghai \
  -v /h-ui/bin:/bin \
  -v /h-ui/data:/data \
  -v /h-ui/export:/export \
  -v /h-ui/logs:/logs \
  jonssonyan/h-ui
```

## Compilation

[build.bat](build.bat) / [build.sh](build.sh)

```bash
go build -o h-ui -trimpath -ldflags "-s -w"
```

## Other

## Contributors

<a href="https://github.com/jonssonyan/h-ui/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=jonssonyan/h-ui" />
</a>

## License