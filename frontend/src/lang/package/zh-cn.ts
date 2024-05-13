export default {
  // 路由国际化
  route: {
    account: "账户",
    accountList: "账户管理",
    hysteria: "Hysteria",
    hysteriaList: "Hysteria 管理",
    config: "系统",
    configList: "系统设置",
    monitor: "监控",
    monitorSystem: "系统监控",
    log: "日志",
    logSystem: "系统日志",
    logHysteria: "Hysteria 日志",
    info: "信息",
    infoAccount: "账户信息",
    telegram: "Telegram",
    telegramList: "Telegram 管理",
  },
  // 登录页面国际化
  login: {
    title: "H UI",
    username: "用户名",
    password: "密码",
    login: "登 录",
  },
  // 导航栏国际化
  navbar: {
    logout: "注销",
  },
  common: {
    id: "编号",
    createTime: "创建时间",
    operate: "操作",
    edit: "编辑",
    delete: "删除",
    deleted: "状态",
    all: "全部",
    enable: "正常",
    disable: "禁用",
    search: "搜索",
    reset: "重设",
    add: "新增",
    confirm: "确定",
    cancel: "取消",
  },
  info: {
    expireTime: "年-月-日 时:分:秒",
  },
  account: {
    username: "用户名",
    pass: "登录密码",
    conPass: "连接密码",
    quota: "配额",
    download: "下载",
    upload: "上传",
    expireTime: "过期时间",
    role: "角色",
    unit: "单位",
    createTime: "注册时间",
  },
  config: {
    huiWebPort: "H UI Web 端口",
    hysteria2TrafficTime: "Hysteria2 流量倍数",
  },
  hysteria: {
    enable: "开启",
    listen: "监听地址",
    tls: "TLS",
    obfs: "混淆",
    quic: "QUIC 参数",
    bandwidth: "带宽",
    speedTest: "速度测试",
    udp: "UDP",
    resolver: "DNS 解析",
    acl: "ACL",
    outbounds: "出站规则",
    http: "流量统计 API",
    masquerade: "伪装",
    config: {
      enable: "开启/关闭",
      listen:
        "当只有端口没有 IP 地址时，服务器将监听所有可用的 IPv4 和 IPv6 地址。要仅监听 IPv4，可以使用 0.0.0.0:443。要仅监听 IPv6，可以使用 [::]:443。",
      tls: {
        cert: "CERT 路径",
        key: "KEY 路径",
      },
      acme: {
        domains: "域名",
        email: "邮箱",
        ca: "要使用的 CA。可以是 letsencrypt 或 zerossl。",
        disableHTTP: "禁用 HTTP 挑战。",
        disableTLSALPN: "禁用 TLS-ALPN 挑战。",
        altHTTPPort:
          "用于 HTTP 挑战的监听端口。 （注意： 改为非 80 需要另行配置端口转发或者 HTTP 反向代理，否则证书会签署失败！）",
        altTLSALPNPort:
          "用于 TLS-ALPN 挑战的监听端口。 （注意： 改为非 443 需要另行配置端口转发或者 SNI Proxy，否则证书会签署失败！）",
        dir: "存储 ACME 账户密钥和证书的目录。",
        listenHost:
          "用于 ACME 服务器验证的监听地址（不含端口）。默认监听所有可用的地址。",
      },
      obfs: {
        type: "类型",
        salamander: {
          password: "替换为你的混淆密码。",
        },
      },
      quic: {
        initStreamReceiveWindow: "初始的 QUIC 流接收窗口大小。",
        maxStreamReceiveWindow: "最大的 QUIC 流接收窗口大小。",
        initConnReceiveWindow: "初始的 QUIC 连接接收窗口大小。",
        maxConnReceiveWindow: "最大的 QUIC 连接接收窗口大小。",
        maxIdleTimeout:
          "最长空闲超时时间。服务器会在多长时间没有收到任何客户端数据后关闭连接。",
        maxIncomingStreams: "最大并发传入流的数量。",
        disablePathMTUDiscovery: "禁用 MTU 探测。",
      },
      bandwidth: {
        up: "上传",
        down: "下载",
      },
      ignoreClientBandwidth: "忽略客户端带宽设置",
      speedTest: "启用后，服务端将允许客户端进行下载和上传速度测试。",
      disableUDP: "disableUDP 启用后服务端禁用 UDP 转发，只支持 TCP。",
      udpIdleTimeout:
        "udpIdleTimeout 用于指定服务器对于每个 UDP 会话，在没有流量时保持本地 UDP 端口的时间长度。概念上与 NAT 的 UDP 会话超时时间相似。",
      resolver: {
        type: "类型",
        tcp: {
          addr: "TCP DNS 服务器地址。",
          timeout: "DNS 查询超时时间。",
        },
        udp: {
          addr: "UDP DNS 服务器地址。",
          timeout: "DNS 查询超时时间。",
        },
        tls: {
          addr: "DNS over TLS 服务器地址。",
          timeout: "DNS 查询超时时间。",
          sni: "DNS over TLS 服务器的 SNI。",
          insecure: "禁用 TLS 证书验证。",
        },
        https: {
          addr: "DNS over HTTPS 服务器地址。",
          timeout: "DNS 查询超时时间。",
          sni: "DNS over TLS 服务器的 SNI。",
          insecure: "禁用 TLS 证书验证。",
        },
      },
      acl: {
        file: "ACL 文件的路径。",
        inline: "内联 ACL 规则的列表。",
        geoip:
          "可选。取消注释以启用。GeoIP 数据库文件的路径。如果省略这个字段，Hysteria 会自动下载最新的数据库到工作目录。",
        geosite:
          "可选。取消注释以启用。GeoSite 数据库文件的路径。如果省略这个字段，Hysteria 会自动下载最新的数据库到工作目录。",
        geoUpdateInterval:
          "可选。GeoIP/GeoSite 数据库刷新的间隔。默认为 168 小时（1 周）。仅在 GeoIP/GeoSite 数据库是自动下载的情况下生效。",
      },
      outbounds: {
        name: "出站规则的名称。在 ACL 中使用。",
        type: "类型",
        socks5: {
          addr: "SOCKS5 代理地址。",
          username: "可选。SOCKS5 代理用户名。",
          password: "可选。SOCKS5 代理密码。",
        },
        http: {
          url: "HTTP/HTTPS 代理 URL。(可以是 http:// 或 https:// 开头)",
          insecure: "可选。禁用 TLS 证书验证。仅适用于 HTTPS 代理。",
        },
        direct: {
          mode: "类型",
          bindIPv4: "要绑定的本地 IPv4 地址。",
          bindIPv6: "要绑定的本地 IPv6 地址。",
          bindDevice: "要绑定的本地网卡。",
        },
      },
      trafficStats: {
        listen: "监听地址。",
      },
      masquerade: {
        type: "类型",
        file: {
          dir: "用于提供文件的目录。",
        },
        proxy: {
          url: "要代理的网站的 URL。",
          rewriteHost:
            "是否重写 Host 头以匹配被代理的网站。如果目标网站通过 Host 识别请求的网站，这个选项是必须的。",
        },
        string: {
          content: "要返回的字符串。",
          headers: "可选。要返回的 HTTP 头列表。",
          statusCode: "可选。要返回的 HTTP 状态码。默认为 200。",
        },
        listenHTTP: "HTTP (TCP) 监听地址。",
        listenHTTPS: "HTTPS (TCP) 监听地址。",
        forceHTTPS:
          "是否强制使用 HTTPS。如果启用，HTTP 请求将被重定向到 HTTPS。",
      },
    },
  },
};
