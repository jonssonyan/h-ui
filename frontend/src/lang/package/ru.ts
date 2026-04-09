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
      "Выражение запланированной задачи, пример: https://pkg.go.dev/github.com/robfig/cron/v3",
    resetTrafficMonth: "Запуск раз в месяц, в полночь, первого числа",
    resetTrafficWeek: "Запуск раз в неделю, в полночь между Сб и Вс",
  },
  monitor: {
    huiVersion: "H UI Version",
    cpuPercent: "CPU Usage",
    memPercent: "Memory Usage",
    diskPercent: "Disk Usage",
    hysteria2UserTotal: "Колич польз в сети",
    hysteria2DeviceTotal: "Колич устр в сети",
    hysteria2Version: "Hysteria2 Версия",
    hysteria2Running: "Hysteria2 Статус",
    hysteria2RunningTrue: "Запущено",
    hysteria2RunningFalse: "Остановлено",
  },
  log: {
    numLine: "Количество строк",
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
    hysteria2Version: "Hysteria2 Версия",
    hysteria2Running: "Hysteria2 Статус",
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
        "Переключение между портами, несколько отдельных портов: 1234,5678,9012; Диапазон портов: 20000-50000; Комбинация обоих: 1234,5000-6000,7044,8000-9000",
      clashExtension: "Продление подписки на Clash",
      listen:
        "Если IP-адрес не указан, сервер будет прослушивать все интерфейсы, как IPv4, так и IPv6. Чтобы прослушивать только IPv4, вы можете использовать 0.0.0.0:443. Для прослушивания только по протоколу IPv6 вы можете использовать [::]:443.",
      tlsType: "TLS type",
      tls: {
        cert: "Путь к Cert file.",
        key: "Путь к Key file.",
        sniGuard:
          'Проверьте SNI, предоставленный клиентом. Принимайте соединение только в том случае, если оно соответствует тому, что указано в сертификате. В противном случае завершите подтверждение TLS. Установите значение strict, чтобы обеспечить соблюдение этого правила. Установите значение disable, чтобы полностью отключить это. По умолчанию используется dns-san, который включает эту функцию только в том случае, если сертификат содержит расширение "Альтернативное имя субъекта" с доменным именем в нем.',
      },
      acme: {
        domains: "Domains",
        email: "Email",
        ca: "The CA to use. Can be letsencrypt or zerossl.",
        listenHost:
          "Адрес хоста (не включая порт) для прослушивания запроса ACME. Если этот параметр не указан, сервер будет прослушивать все интерфейсы.",
        dir: "Каталог для хранения ключа учетной записи ACME и сертификатов.",
        type: "Тип запроса ACME. Может быть http, tls или dns.",
        http: {
          altPort:
            "Порт прослушивания для запросов HTTP. (Примечание: для перехода на порт, отличный от 80, требуется переадресация портов или обратный прокси-сервер HTTP, иначе вызов завершится неудачей!)",
        },
        tls: {
          altPort:
            "Порт прослушивания для вызовов TLS-ALPN. (Примечание: для перехода на порт, отличный от 443, требуется переадресация портов или обратный прокси-сервер TLS, в противном случае вызов завершится неудачей!)",
        },
        dns: {
          name: "DNS поставщик. Дополнительные сведения см. в разделе Настройка ACME DNS.",
          config: "ACME DNS настройка",
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
          password: "Замени на надежный пароль.",
        },
      },
      quic: {
        initStreamReceiveWindow: "Начальный размер окна приема потока QUIC.",
        maxStreamReceiveWindow: "Максимальный размер окна приема QUIC stream.",
        initConnReceiveWindow:
          "Размер окна начального QUIC-соединения.",
        maxConnReceiveWindow:
          "Максимальный размер окна QUIC-соединения.",
        maxIdleTimeout:
          "Максимальное время ожидания в режиме ожидания. Как долго сервер будет считать, что клиент все еще подключен без каких-либо действий.",
        maxIncomingStreams:
          "Максимальное количество одновременных входящих потоков.",
        disablePathMTUDiscovery: "Отключить обнаружение MTU QUIC path.",
      },
      bandwidth: {
        up: "Вверх",
        down: "Вниз",
      },
      ignoreClientBandwidth:
        "Когда этот параметр включен, сервер игнорирует любые указания по пропускной способности, установленные клиентами",
      speedTest:
        "speedTest включает встроенный сервер тестирования скорости. Если этот параметр включен, клиенты могут тестировать скорость загрузки с помощью сервера. Дополнительные сведения см. в документации по тестированию скорости.",
      disableUDP:
        "disableUDP отключает переадресацию UDP, разрешая только TCP-соединения.",
      udpIdleTimeout:
        "udpIdleTimeout определяет время, в течение которого сервер будет поддерживать локальный UDP-порт открытым для каждого сеанса UDP, в котором нет активности. Это концептуально аналогично тайм-ауту сеанса NAT UDP.",
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
        file: "Путь к файлу ACL.",
        inline: "Список встроенных правил ACL.",
        geoip:
          "Необязательно. Раскомментируйте, чтобы включить. Путь к файлу базы данных GeoIP. Если это поле не указано, Hysteria автоматически загрузит последнюю версию базы данных в ваш рабочий каталог.",
        geosite:
          "Необязательно. Раскомментируйте, чтобы включить. Путь к файлу базы данных GeoSite. Если это поле не указано, Hysteria автоматически загрузит последнюю версию базы данных в ваш рабочий каталог.",
        geoUpdateInterval:
          "Необязательно. Интервал обновления баз данных GeoIP/GeoSite. По умолчанию - 168 часов (1 неделя). Применяется только в том случае, если базы данных GeoIP/GeoSite загружаются автоматически. (Для получения дополнительной информации ознакомьтесь с примечанием ниже).",
      },
      outbounds: {
        name: "Имя исх соединения (outbound). Это используется в правилах ACL.",
        type: "Type",
        socks5: {
          addr: "Адрес прокси-сервера SOCKS5.",
          username:
            "Рекомендуется. Имя пользователя для прокси-сервера SOCKS5, если требуется аутентификация.",
          password:
            "Рекомендуется. Пароль для прокси-сервера SOCKS5, если требуется аутентификация.",
        },
        http: {
          url: "URL-адрес прокси-сервера HTTP/HTTPS. (Может быть http:// или https://)",
          insecure:
            "Необязательно. Следует ли отключить проверку по протоколу TLS. Применяется только к HTTPS-прокси.",
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
          "Если этот параметр включен, все HTTP-запросы будут перенаправляться на HTTPS.",
      },
    },
  },
};
