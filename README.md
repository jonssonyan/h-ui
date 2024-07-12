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
- Monitor system status and Hysteria2 status
- Limit user traffic, user online status, force users to log off, number of online users, reset user traffic
- Limit the number of users' online devices at the same time, the number of online devices
- User subscription link, node URL, import and export users
- Managing Hysteria2 configurations and Hysteria2 versions
- Change the Web port, modify the Hysteria2 traffic multiplier
- View, import, and export system logs and Hysteria2 logs
- I18n: English, 简体中文
- Page adaptation, support night mode, custom page themes
- More features waiting for you to discover

## Recommended OS

OS: CentOS 8+/Ubuntu 20+/Debian 11+

CPU: x86_64/amd64 arm64/aarch64

Memory: ≥ 256MB

## Deployment

### Quick Install (Recommended)

```bash
bash <(curl -fsSL https://raw.githubusercontent.com/jonssonyan/h-ui/main/install.sh)
```

### Docker

1. Install Docker

   https://docs.docker.com/engine/install/

   ```bash
   bash <(curl -fsSL https://get.docker.com)
   ```

2. Start a container

   ```bash
   docker pull jonssonyan/h-ui

   docker run -d --cap-add=NET_ADMIN \
     --name h-ui --restart always \
     --network=host \
     -v /h-ui/bin:/h-ui/bin \
     -v /h-ui/data:/h-ui/data \
     -v /h-ui/export:/h-ui/export \
     -v /h-ui/logs:/h-ui/logs \
     jonssonyan/h-ui
   ```

   Custom web port, default is 8081

   ```bash
   docker run -d --cap-add=NET_ADMIN \
     --name h-ui --restart always \
     --network=host \
     -v /h-ui/bin:/h-ui/bin \
     -v /h-ui/data:/h-ui/data \
     -v /h-ui/export:/h-ui/export \
     -v /h-ui/logs:/h-ui/logs \
     jonssonyan/h-ui \
     ./h-ui -p 8081
   ```

   Set the time zone, default is Asia/Shanghai

   ```bash
   docker run -d --cap-add=NET_ADMIN \
     --name h-ui --restart always \
     --network=host \
     -e TZ=Asia/Shanghai \
     -v /h-ui/bin:/h-ui/bin \
     -v /h-ui/data:/h-ui/data \
     -v /h-ui/export:/h-ui/export \
     -v /h-ui/logs:/h-ui/logs \
     jonssonyan/h-ui
   ```

Uninstall

```bash
docker rm -f h-ui
docker rmi jonssonyan/h-ui
rm -rf /h-ui
```

### systemd

Executable files: https://github.com/jonssonyan/h-ui/releases

```bash
mkdir -p /usr/local/h-ui/
curl -fsSL https://github.com/jonssonyan/h-ui/releases/latest/download/h-ui-linux-amd64 -o /usr/local/h-ui/h-ui && chmod +x /usr/local/h-ui/h-ui
curl -fsSL https://raw.githubusercontent.com/jonssonyan/h-ui/main/h-ui.service -o /etc/systemd/system/h-ui.service
# Custom web port, default is 8081
# sed -i "s|^ExecStart=.*|ExecStart=/usr/local/h-ui/h-ui -p 8081|" "/etc/systemd/system/h-ui.service"
systemctl daemon-reload
systemctl enable h-ui
systemctl restart h-ui
```

Uninstall

```bash
systemctl stop h-ui
rm -rf /etc/systemd/system/h-ui.service /usr/local/h-ui/
```

## Default Installation Information

- Panel Port: 8081
- Username/Password: sysadmin

## System Upgrade

Export the user, system configuration, and Hysteria2 configuration in the management background, redeploy the latest
version of h-ui, and import the data into the management background after the deployment is complete.

## FAQ

[English > FAQ](./docs/FAQ.md)

## Performance Optimization

- Scheduled server restart

    ```bash
    0 4 * * * /sbin/reboot
    ```

- Install Network Accelerator
    - [TCP Brutal](https://github.com/apernet/tcp-brutal) (Recommended)
    - [teddysun/across#bbrsh](https://github.com/teddysun/across#bbrsh)
    - [Chikage0o0/Linux-NetSpeed](https://github.com/ylx2016/Linux-NetSpeed)
    - [ylx2016/Linux-NetSpeed](https://github.com/ylx2016/Linux-NetSpeed)

## Client

https://v2.hysteria.network/docs/getting-started/3rd-party-apps/

## Development

Go >= 1.20, Node.js >= 18.12.0

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

Contact me: https://t.me/jonssonyan_channel

If you want to chat on Telegram, you can see here: https://t.me/jonssonyan_chat

You can subscribe to my channel on YouTube: https://www.youtube.com/@jonssonyan

## Contributors

Thanks to everyone who contributed to this project.

<a href="https://github.com/jonssonyan/h-ui/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=jonssonyan/h-ui" />
</a>

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=jonssonyan/h-ui&type=Date)](https://star-history.com/#jonssonyan/h-ui&Date)

## License

[GPL-3.0](LICENSE)