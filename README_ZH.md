<div align="center">

<a href="https://github.com/jonssonyan/h-ui"><img src="./docs/images/head-cover.png" alt="H UI" width="150" /></a>

<h1 align="center">H UI</h1>

[English](README.md) / 简体中文

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

## 主要功能

- 轻量级、资源占用低、易于部署
- 监控系统状态和 Hysteria2 状态
- 限制用户流量、用户在线状态、强制用户下线、在线用户数
- 限制用户同时在线设备数、在线设备数量
- 用户订阅链接、节点URL、导入和导出用户
- 管理 Hysteria2 配置和 Hysteria2 版本
- 更改 Web 端口、修改 Hysteria2 流量倍数
- 查看、导入和导出系统日志和 Hysteria2 日志
- 多国语言支持: English, 简体中文
- 页面适配、支持夜间模式、自定义页面主题

## 建议系统

系统: CentOS Stream 8/Ubuntu 20.04 LTS/Debian 11.1

CPU: x86_64/amd64

内存: ≥ 128MB

## 部署

面板地址: http://[your_ip/your_domain]:8081

默认管理员用户名：sysadmin 密码：123456，部署后请及时登录管理后台修改密码

### Docker (推荐)

1. 安装 Docker

   https://docs.docker.com/engine/install/

   ```bash
   bash <(curl -fsSL https://get.docker.com)
   ```

2. 启动容器

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

### 手动

下载可执行文件: https://github.com/jonssonyan/h-ui/releases

```bash
mkdir h-ui && cd h-ui
curl -L -o h-ui https://github.com/jonssonyan/h-ui/releases/download/v0.0.1/h-ui-linux-amd64 && chmod +x ./h-ui && ./h-ui
```

## 系统升级

在管理后台将用户、系统配置、Hysteria2 配置导出，重新部署最新版的 h-ui，部署完成之后在管理后台将数据导入

## 常见问题

[简体中文 > 常见问题](./docs/FAQ_ZH.md)

## 性能优化

- 定时重启服务器

    ```bash
    0 4 * * * /sbin/reboot
    ```

- 安装网络加速
    - [Chikage0o0/Linux-NetSpeed](https://github.com/ylx2016/Linux-NetSpeed)
    - [ylx2016/Linux-NetSpeed](https://github.com/ylx2016/Linux-NetSpeed)
    - [teddysun/across#bbrsh](https://github.com/teddysun/across#bbrsh)

## 客户端

https://v2.hysteria.network/zh/docs/getting-started/3rd-party-apps/

## 开发

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

## 构建

- frontend

   ```bash
   npm run build:prod
   ```

- backend

  Windows: [build.bat](build.bat)

  Mac/Linux: [build.sh](build.sh)

## 其他

你可以在 YouTube 上关注我: https://www.youtube.com/@jonssonyan

## 贡献者

<a href="https://github.com/jonssonyan/h-ui/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=jonssonyan/h-ui" />
</a>

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=jonssonyan/h-ui&type=Date)](https://star-history.com/#jonssonyan/h-ui&Date)

## 开源协议

[GPL-3.0](LICENSE)