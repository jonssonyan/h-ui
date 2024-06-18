<div align="center">

<a href="https://github.com/jonssonyan/h-ui"><img src="./docs/images/head-cover.png" alt="H UI" width="150" /></a>

<h1 align="center">H UI</h1>

English / [简体中文](README_ZH.md)

Just the panel for Hysteria2

仅仅是 Hysteria2 的面板

<p>
<a href="https://www.gnu.org/licenses/gpl-3.0.html"><img src="https://img.shields.io/github/license/jonssonyan/h-ui" alt="License: GPL-3.0"></a>
<a href="https://github.com/jonssonyan/h-ui/stargazers"><img src="https://img.shields.io/github/stars/jonssonyan/h-ui" alt="GitHub stars"></a>
<a href="https://github.com/jonssonyan/h-ui/forks"><img src="https://img.shields.io/github/forks/jonssonyan/h-ui" alt="GitHub forks"></a>
<a href="https://github.com/jonssonyan/h-ui/releases"><img src="https://img.shields.io/github/v/release/jonssonyan/h-ui" alt="GitHub release"></a>
<a href="https://hub.docker.com/r/jonssonyan/h-ui"><img src="https://img.shields.io/docker/pulls/jonssonyan/h-ui" alt="Docker pulls"></a>
</p>

![cover](./docs/images/cover.png)

</div>

## Features

- Lightweight, low resource usage, easy to deploy
- Monitor system status and Hysterai2 status
- Limit user traffic, user online status, force users to log off, number of online users
- Limit the number of users' online devices at the same time, the number of online devices
- User subscription link, node URL, import and export users
- Managing Hysterai2 configurations and Hysterai2 versions
- Change the Web port, modify the Hysteria2 traffic multiplier
- View, import, and export system logs and Hysteria2 logs
- I18n: English, 简体中文
- Page adaptation, support night mode, custom page themes

## Recommended OS

OS version: CentOS Stream 8/Ubuntu 20.04 LTS/Debian 11.1

CPU Architecture: x86_64/amd64

Memory: ≥ 256MB

## Deployment

### Docker (Recommended)

1. Install Docker

    ```bash
    source <(curl -L https://github.com/jonssonyan/install-script/raw/main/docker-install.sh)
    ```

2. Start a container

   ```bash
   docker pull jonssonyan/h-ui

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

### Manual

## Development

- frontend

   ```bash
   cd frontend
   pnpm install
   npm run dev
   ```

- backend

   ```bash
   go run main.go
   ```

## Build

- frontend

   ```bash
   npm run build:prod
   ```

- backend

  Windows: [build.bat](build.bat)

  Mac/Linux: [build.sh](build.sh)

## Other

you can contact me at YouTube: https://www.youtube.com/@jonssonyan

## Contributors

<a href="https://github.com/jonssonyan/h-ui/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=jonssonyan/h-ui" />
</a>

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=jonssonyan/h-ui&type=Date)](https://star-history.com/#jonssonyan/h-ui&Date)

## License

[GPL-3.0](LICENSE)