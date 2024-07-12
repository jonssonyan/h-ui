export default {
  // 路由国际化
  route: {
    account: "Account",
    accountList: "Account Manage",
    hysteria: "Hysteria",
    hysteriaList: "Hysteria Manage",
    config: "System",
    configList: "System Config",
    monitor: "Monitor",
    monitorSystem: "System Monitor",
    log: "Log",
    logSystem: "System Log",
    logHysteria: "Hysteria Log",
    info: "Info",
    infoAccount: "Account Info",
    telegram: "Telegram",
    telegramList: "Telegram Manage",
  },
  // 登录页面国际化
  login: {
    title: "H UI",
    username: "Username",
    password: "Password",
    login: "Login",
  },
  // 导航栏国际化
  navbar: {
    logout: "Logout",
  },
  common: {
    id: "ID",
    createTime: "CreateTime",
    operate: "Operate",
    edit: "Edit",
    delete: "Delete",
    deleted: "Status",
    all: "All",
    enable: "Enable",
    disable: "Disable",
    search: "Search",
    reset: "Reset",
    add: "Add",
    confirm: "Confirm",
    cancel: "Cancel",
    copySuccess: "Copy successful",
    subscribe: "Subscribe",
    subscribeQrCode: "Subscribe QR Code",
    nodeUrl: "Node URL",
    nodeQrCode: "Node QR Code",
    resetTraffic: "Reset traffic",
    import: "Import",
    export: "Export",
    save: "Save",
    update: "Update",
    downloadSuccess: "Download successful",
    wait: "The version is being changed, please wait a moment",
    enableSuccess: "Hysteria2 start successful",
    disableSuccess: "Hysteria2 stop successful",
    success: "Success",
    refresh: "Refresh",
    yes: "Yes",
    no: "No",
  },
  info: {
    expireTime: "y-M-d H:m:s",
    greeting1: "The cool and fresh air awakens your energy for the day🌅！",
    greeting2: "Good morning，",
    greeting3: "Good afternoon，",
    greeting4: "Good evening，",
    greeting5:
      "I want to be a shooting star, cutting through the darkness, just to illuminate your dreams, good night🌛！",
  },
  account: {
    username: "Username",
    pass: "Pass",
    conPass: "ConPass",
    quota: "Quota",
    download: "Download",
    upload: "Upload",
    expireTime: "ExpireTime",
    kickUtilTimeLast: "Offline Remaining Time",
    kickUtilTime: "Offline Time",
    deviceNo: "Limit Devices",
    onlineStatus: "Online Status",
    online: "Online",
    offline: "Offline",
    device: "Online Devices",
    role: "Role",
    unit: "Unit",
    createTime: "Create Time",
    releaseSuccess: "Release successful",
    kick: "Kick",
    kickTip: "Force user to log off",
    releaseKick: "Release",
    releaseKickTip: "Remove offline status",
  },
  config: {
    huiWebPort: "H UI Web Port",
    hysteria2TrafficTime: "Hysteria2 Traffic Time",
    huiCrtPath: "H UI CRT File Path",
    huiKeyPath: "H UI KEY File Path",
    restartServer: "Restart Panel",
    restartTip: "Restarting, please refresh",
    useHysteria2Cert: "Use Hysteria2 cert",
    huiHttps: "Open https on the panel",
  },
  monitor: {
    huiVersion: "H UI Version",
    cpuPercent: "CPU Usage",
    memPercent: "Memory Usage",
    diskPercent: "Disk Usage",
    hysteria2UserTotal: "Number of online users",
    hysteria2DeviceTotal: "Number of online devices",
    hysteria2Version: "Hysteria2 Version",
    hysteria2Running: "Hysteria2 Status",
    hysteria2RunningTrue: "Running",
    hysteria2RunningFalse: "Stop",
  },
  log: {
    numLine: "Number of lines",
  },
  hysteria: {
    enable: "Enable",
    disable: "Disable",
    addConfigItem: "Add Config Item",
    hysteria2Version: "Hysteria2 Version",
    hysteria2Running: "Hysteria2 Status",
    hysteria2ChangeVersion: "Change",
    addOutbound: "Add Outbound",
    extension: "Extension",
    listen: "Listen",
    tls: "TLS",
    obfs: "Obfuscation",
    quic: "QUIC parameters",
    bandwidth: "Bandwidth",
    speedTest: "Speed Test",
    udp: "UDP",
    resolver: "Resolver",
    sniff: "Protocol Sniffing",
    acl: "ACL",
    outbounds: "Outbounds",
    http: "Traffic Stats API (HTTP)",
    masquerade: "Masquerade",
    config: {
      enable: "Enable/Disable",
      remark: "Remark",
      portHopping:
        "Port Hopping, Multiple individual ports: 1234,5678,9012; A range of ports: 20000-50000; A combination of both: 1234,5000-6000,7044,8000-9000",
      listen:
        "When the IP address is omitted, the server will listen on all interfaces, both IPv4 and IPv6. To listen on IPv4 only, you can use 0.0.0.0:443. To listen on IPv6 only, you can use [::]:443.",
      tlsType: "TLS type",
      tls: {
        cert: "The path to the Cert file.",
        key: "The path to the Key file.",
      },
      acme: {
        domains: "Domains",
        email: "Email",
        ca: "The CA to use. Can be letsencrypt or zerossl.",
        listenHost:
          "The host address (not including the port) to listen on for the ACME challenge. If omitted, the server will listen on all interfaces.",
        dir: "The directory to store the ACME account key and certificates.",
        type: "ACME challenge type. Can be http, tls, or dns.",
        http: {
          altPort:
            "Listening port for HTTP challenges. (Note: Changing to a port other than 80 requires port forwarding or HTTP reverse proxy, or the challenge will fail!)",
        },
        tls: {
          altPort:
            "Listening port for TLS-ALPN challenges. (Note: Changing to a port other than 443 requires port forwarding or TLS reverse proxy, or the challenge will fail!)",
        },
        dns: {
          name: "DNS provider. For details, refer to ACME DNS Configuration.",
          config: "ACME DNS Configuration",
        },
        disableHTTP: "Disable HTTP challenge.",
        disableTLSALPN: "Disable TLS-ALPN challenge.",
        altHTTPPort:
          "Alternate HTTP challenge port. (Note: If you want to use anything other than 80, you must set up port forward/HTTP reverse proxy from 80 to that port, otherwise ACME will not be able to issue the certificate.)",
        altTLSALPNPort:
          "Alternate TLS-ALPN challenge port. (Note: If you want to use anything other than 443, you must set up port forward/SNI proxy from 443 to that port, otherwise ACME will not be able to issue the certificate.)",
      },
      obfs: {
        type: "Type",
        salamander: {
          password: "Replace with a strong password of your choice.",
        },
      },
      quic: {
        initStreamReceiveWindow: "The initial QUIC stream receive window size.",
        maxStreamReceiveWindow: "The maximum QUIC stream receive window size.",
        initConnReceiveWindow:
          "The initial QUIC connection receive window size.",
        maxConnReceiveWindow:
          "The maximum QUIC connection receive window size.",
        maxIdleTimeout:
          "The maximum idle timeout. How long the server will consider the client still connected without any activity.",
        maxIncomingStreams:
          "The maximum number of concurrent incoming streams.",
        disablePathMTUDiscovery: "Disable QUIC path MTU discovery.",
      },
      bandwidth: {
        up: "Up",
        down: "Down",
      },
      ignoreClientBandwidth:
        "When enabled, makes the server to disregard any bandwidth hints set by clients",
      speedTest:
        "speedTest enables the built-in speed test server. When enabled, clients can test their download and upload speeds with the server. For more information, see the Speed Test documentation.",
      disableUDP:
        "disableUDP disables UDP forwarding, only allowing TCP connections.",
      udpIdleTimeout:
        "udpIdleTimeout specifies the amount of time the server will keep a local UDP port open for each UDP session that has no activity. This is conceptually similar to the NAT UDP session timeout.",
      resolver: {
        type: "Type",
        tcp: {
          addr: "The address of the TCP resolver.",
          timeout: "The timeout for DNS queries.",
        },
        udp: {
          addr: "The address of the UDP resolver.",
          timeout: "The timeout for DNS queries.",
        },
        tls: {
          addr: "The address of the TLS resolver.",
          timeout: "The timeout for DNS queries.",
          sni: "The SNI to use for the TLS resolver.",
          insecure: "Disable TLS verification for the TLS resolver.",
        },
        https: {
          addr: "The address of the HTTPS resolver.",
          timeout: "The timeout for DNS queries.",
          sni: "The SNI to use for the TLS resolver.",
          insecure: "Disable TLS verification for the TLS resolver.",
        },
      },
      sniff: {
        enable: "Whether to enable protocol sniffing.",
        timeout:
          "Sniffing timeout. If the protocol/domain cannot be determined within this time, the original address will be used to initiate the connection.",
        rewriteDomain:
          "Whether to rewrite requests that are already in domain name form. If enabled, requests with the target address already in domain name form will still be sniffed.",
        tcpPorts:
          "List of TCP ports. Only TCP requests on these ports will be sniffed.",
        udpPorts:
          "List of UDP ports. Only UDP requests on these ports will be sniffed.",
      },
      aclType: "ACL type",
      acl: {
        file: "The path to the ACL file.",
        inline: "The list of inline ACL rules.",
        geoip:
          "Optional. Uncomment to enable. The path to the GeoIP database file. If this field is omitted, Hysteria will automatically download the latest database to your working directory.",
        geosite:
          "Optional. Uncomment to enable. The path to the GeoSite database file. If this field is omitted, Hysteria will automatically download the latest database to your working directory.",
        geoUpdateInterval:
          "Optional. The interval at which to refresh the GeoIP/GeoSite databases. 168 hours (1 week) by default. Only applies if the GeoIP/GeoSite databases are automatically downloaded. (Check the note below for more information.)",
      },
      outbounds: {
        name: "The name of the outbound. This is used in ACL rules.",
        type: "Type",
        socks5: {
          addr: "The address of the SOCKS5 proxy.",
          username:
            "Optional. The username for the SOCKS5 proxy, if authentication is required.",
          password:
            "Optional. The password for the SOCKS5 proxy, if authentication is required.",
        },
        http: {
          url: "The URL of the HTTP/HTTPS proxy. (Can be http:// or https://)",
          insecure:
            "Optional. Whether to disable TLS verification. Applies to HTTPS proxies only.",
        },
        direct: {
          mode: "Type",
          bindIPv4: "The local IPv4 address to bind to.",
          bindIPv6: "The local IPv6 address to bind to.",
          bindDevice: "The local network interface to bind to.",
        },
      },
      trafficStats: {
        listen: "The address to listen on.",
      },
      masquerade: {
        type: "Type",
        file: {
          dir: "The directory to serve files from.",
        },
        proxy: {
          url: "The URL of the website to proxy.",
          rewriteHost:
            "Whether to rewrite the Host header to match the proxied website. This is required if the target web server uses Host to determine which site to serve.",
        },
        string: {
          content: "The string to return.",
          headers: "Optional. The headers to return.",
          statusCode: "Optional. The status code to return. 200 by default.",
        },
        listenHTTP: "HTTP (TCP) listen address.",
        listenHTTPS: "HTTPS (TCP) listen address.",
        forceHTTPS:
          "Whether to force HTTPS. If enabled, all HTTP requests will be redirected to HTTPS.",
      },
    },
  },
};
