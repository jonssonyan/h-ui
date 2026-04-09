export default {
  // 路由国际化
  route: {
    account: "Аккаунт",
    accountList: "Настрока Аккаунта",
    hysteria: "Hysteria",
    hysteriaList: "Настрока Hysteria",
    config: "Система",
    configList: "Настройки системы",
    monitor: "Мониторинг",
    monitorSystem: "Мониторинг системы",
    log: "Логи",
    logSystem: "Логи системы",
    logHysteria: "Логи Hysteria",
    info: "Информация",
    infoAccount: "Об аккаунте",
    telegram: "Телеграм",
    telegramList: "Настройки Telegram",
  },
  // 登录页面国际化
  login: {
    title: "H UI",
    username: "Логин",
    password: "Пароль",
    login: "Логин",
  },
  // 导航栏国际化
  navbar: {
    logout: "Выход",
  },
  common: {
    id: "ID",
    createTime: "Create Time",
    operate: "Operate",
    edit: "Редактировать",
    delete: "Удалить",
    deleted: "Статус", #Status
    all: "Все",
    enable: "Включить",
    disable: "Выключить",
    search: "Поиск",
    reset: "Сброс",
    add: "Добавить",
    confirm: "Подтвердить",
    cancel: "Отмена",
    copySuccess: "Copy successful",
    subscribe: "Subscribe",
    subscribeQrCode: "Subscribe QR Code",
    nodeUrl: "Node URL",
    nodeQrCode: "Node QR Code",
    resetTraffic: "Reset traffic",
    import: "Импорт",
    export: "Экспорт",
    save: "Сохранить",
    update: "Обновить",
    downloadSuccess: "Загрузка успешна",
    wait: "Версия меняется, ожидайте",
    enableSuccess: "Hysteria2 запущена успешно",
    disableSuccess: "Hysteria2 остановленна успешно",
    success: "Успешно",
    refresh: "Обновить",
    yes: "Да",
    no: "Нет",
    securityRisk: "Security Risks",
    defaultPassTip: `Пожалуйста, измените пароль для входа в систему по умолчанию как можно скорее, рекомендуется установить надежный пароль для защиты вашей учетной записи. <a href="/#/account/list?focus=change-pass" style="color: #00BFFF">Click here</a> to change`,
    noHttpsTip: `Ваша панель не использует HTTPS, что делает передачу данных небезопасной, пожалуйста, включите HTTPS как можно скорее, чтобы защитить информацию пользователя. <a href="/#/config/list?focus=huiHttps" style="color: #00BFFF">Click here</a> to enable`,
  },
  info: {
    expireTime: "y-M-d H:m:s",
    greeting1: "Воздух чистый, с холодком, пускай кровь разгонит и дух твой пробудит для великих дел 🌅！",
    greeting2: "Доброе утро，",
    greeting3: "Добрый день，",
    greeting4: "Добрый вечер，",
    greeting5:
      "Хочу падующей звездой прошить темень ночную, лишь бы озарить твои грёзы. Доброй ночи🌛！",
  },
  account: {
    remark: "Remark",
    username: "Username",
    pass: "Pass",
    conPass: "ConPass",
    quota: "Quota",
    download: "Download",
    upload: "Upload",
    expireTime: "Expire Time",
    kickUtilTimeLast: "Offline Remaining Time",
    kickUtilTime: "Offline Time",
    deviceNo: "Limit Devices",
    onlineStatus: "Online Status",
    online: "Онлайн",
    offline: "Офлайн",
    device: "Online Devices",
    role: "Role",
    unit: "Unit",
    loginAt: "Last login time",
    conAt: "Last connection time",
    createTime: "Create Time",
    releaseSuccess: "Release successful",
    kick: "Kick",
    kickTip: "Force user to log off",
    releaseKick: "Release",
    releaseKickTip: "Remove offline status",
  },
  config: {
    huiWebPort: "H UI Web Port",
    huiWebContext: "H UI Web Context",
    hysteria2TrafficTime: "Hysteria2 Traffic Time",
    huiCrtPath: "H UI CRT File Path",
    huiKeyPath: "H UI KEY File Path",
    uploadCrtFile: "Upload CRT File",
    uploadKeyFile: "Upload KEY File",
    restartServer: "Restart Panel",
    restartTip: "Restarting, please refresh",
    useHysteria2Cert: "Use Hysteria2 cert",
    huiHttps: "Open https on the panel",
    resetTrafficCron: "Reset traffic schedule task",
    resetTrafficCronTip:
      "Scheduled task expression, reference: https://pkg.go.dev/github.com/robfig/cron/v3",
    resetTrafficMonth: "Run once a month, midnight, first of month",
    resetTrafficWeek: "Run once a week, midnight between Sat/Sun",
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
  telegram: {
    placeholder: "Placeholder",
    enable: "Enable",
    disable: "Disable",
    telegramEnable: "Enable/Disable",
    telegramToken: "Telegram Token",
    telegramChatId: "Telegram ChatId",
    telegramJob: "Job List",
    telegramLoginJob: "Login reminder",
    telegramLoginJobEnable: "Enable/Disable",
    telegramLoginJobText: "Content template",
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
      clashExtension: "Clash subscription extension",
      listen:
        "When the IP address is omitted, the server will listen on all interfaces, both IPv4 and IPv6. To listen on IPv4 only, you can use 0.0.0.0:443. To listen on IPv6 only, you can use [::]:443.",
      tlsType: "TLS type",
      tls: {
        cert: "The path to the Cert file.",
        key: "The path to the Key file.",
        sniGuard:
          'Verify the SNI provided by the client. Accept the connection only when it matches what\'s in the certificate. Terminate the TLS handshake otherwise. Set to strict to enforce this behavior. Set to disable to disable this entirely. The default is dns-san, which enables this feature only when the certificate contains the "Subject Alternative Name" extension with a domain name in it.',
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
          "Альтернативный порт запроса HTTP. (Примечание: Если вы хотите использовать что-либо, отличное от 80, вы должны настроить прямой/обратный HTTP-прокси-сервер с 80 на этот порт, иначе ACME не сможет выдать сертификат.)",
        altTLSALPNPort:
          "Альтернативный порт вызова TLS-ALPN. (Примечание: Если вы хотите использовать что-либо, отличное от 443, вы должны настроить переадресацию портов/SNI-прокси с 443 на этот порт, иначе ACME не сможет выдать сертификат.)",
      },
      obfs: {
        type: "Type",
        salamander: {
          password: "Замени на надежный пароль по вашему выбору.",
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
        up: "Вверх",
        down: "Вниз",
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
        enable: "Следует ли включать protocol sniffing.",
        timeout:
          "Sniffing timeout. Если протокол/домен не могут быть определены в течение этого времени, для установления соединения будет использован исходный адрес.",
        rewriteDomain:
          "Следует ли переписывать запросы, которые уже находятся в форме доменного имени. Если эта опция включена, запросы с целевым адресом, уже указанным в форме доменного имени, по-прежнему будут обрабатываться.",
        tcpPorts:
          "Список TCP-портов. Анализироваться(sniffing) будут только TCP-запросы, поступающие на эти порты.",
        udpPorts:
          "Список UDP-портов. Анализироваться(sniffing) будут только UDP-запросы, поступающие на эти порты.",
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
          fastOpen: "Enable TCP fast open.",
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
          url: "URL-адрес веб-сайта для подключения к прокси-серверу.",
          rewriteHost:
            "Следует ли переписать заголовок Host, чтобы он соответствовал прокси-сайту. Это необходимо, если целевой веб-сервер использует Host для определения того, какой сайт обслуживать.",
          insecure: "Отключите проверку TLS для проксируемого веб-сайта.",
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
