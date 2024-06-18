# H UI

仅仅是 Hysteria2 的面板

# 系统要求

系统: CentOS Stream 8/Ubuntu 20.04 LTS/Debian 11.1

CPU: x86_64/amd64

# 安装

- Docker(推荐)

```bash
docker run -d --name hui --restart always \
  --network=host \
  -v /hui/bin:/bin \
  -v /hui/data:/data \
  -v /hui/export:/export \
  -v /hui/logs:/logs \
  hui

```

自定义 Web 端口

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

# 编译

```shell
go build -trimpath -ldflags -o h-ui -ldflags="-s -w"
```

# 其他
