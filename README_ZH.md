[English](README.md)

# H UI

仅仅是 Hysteria2 的面板

# 系统要求

系统: CentOS Stream 8/Ubuntu 20.04 LTS/Debian 11.1

CPU: x86_64/amd64

# 安装

- Docker(推荐)

```bash
docker run -d --name h-ui --restart always \
  --network=host \
  -v /h-ui/bin:/bin \
  -v /h-ui/data:/data \
  -v /h-ui/export:/export \
  -v /h-ui/logs:/logs \
  jonssonyan/h-ui
```

自定义 Web 端口

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

设置时区，默认 Asia/Shanghai

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

# 编译

[build.bat](build.bat) / [build.sh](build.sh)

```bash
go build -o h-ui -trimpath -ldflags "-s -w"
```

# 其他
