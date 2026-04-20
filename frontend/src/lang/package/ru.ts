export default {
  // 路由国际化
  route: {
    account: "Аккаунт",
    accountList: "Управление аккаунтами",
    hysteria: "Hysteria",
    hysteriaList: "Управление Hysteria",
    config: "Система",
    configList: "Системные настройки",
    monitor: "Мониторинг",
    monitorSystem: "Мониторинг системы",
    log: "Логи",
    logSystem: "Системные логи",
    logHysteria: "Логи Hysteria",
    info: "Информация",
    infoAccount: "Информация об аккаунте",
    telegram: "Telegram",
    telegramList: "Управление Telegram",
  },
  // 登录页面国际化
  login: {
    title: "H UI",
    username: "Имя пользователя",
    password: "Пароль",
    login: "Войти",
  },
  // 导航栏国际化
  navbar: {
    logout: "Выйти",
  },
  common: {
    id: "ID",
    createTime: "Время создания",
    operate: "Действия",
    edit: "Редактировать",
    delete: "Удалить",
    deleted: "Статус",
    all: "Все",
    enable: "Включить",
    disable: "Отключить",
    search: "Поиск",
    reset: "Сброс",
    add: "Добавить",
    confirm: "Подтвердить",
    cancel: "Отмена",
    copySuccess: "Скопировано",
    subscribe: "Подписка",
    subscribeQrCode: "QR-код подписки",
    nodeUrl: "URL узла",
    nodeQrCode: "QR-код узла",
    resetTraffic: "Сбросить трафик",
    import: "Импорт",
    export: "Экспорт",
    save: "Сохранить",
    update: "Обновить",
    downloadSuccess: "Загрузка выполнена",
    wait: "Версия изменяется, пожалуйста, подождите",
    enableSuccess: "Hysteria2 успешно запущена",
    disableSuccess: "Hysteria2 успешно остановлена",
    success: "Успешно",
    refresh: "Обновить",
    yes: "Да",
    no: "Нет",
    securityRisk: "Риски безопасности",
    defaultPassTip: `Пожалуйста, как можно скорее смените пароль входа по умолчанию. Рекомендуется установить сложный пароль для защиты вашей учетной записи. <a href="/#/account/list?focus=change-pass" style="color: #00BFFF">Нажмите здесь</a>, чтобы изменить`,
    noHttpsTip: `Ваш сайт не использует HTTPS, из-за чего передача данных небезопасна. Пожалуйста, как можно скорее включите HTTPS для защиты пользовательских данных. <a href="/#/config/list?focus=huiHttps" style="color: #00BFFF">Нажмите здесь</a>, чтобы включить`,
  },
  info: {
    expireTime: "y-M-d H:m:s",
    greeting1:
      "Прохладный и свежий воздух наполняет вас энергией на весь день 🌅！",
    greeting2: "Доброе утро，",
    greeting3: "Добрый день，",
    greeting4: "Добрый вечер，",
    greeting5:
      "Хочу стать падающей звездой, прорезающей тьму, лишь бы осветить ваши сны. Спокойной ночи 🌛！",
  },
  account: {
    remark: "Примечание",
    username: "Имя пользователя",
    pass: "Пароль",
    conPass: "Подтверждение пароля",
    quota: "Квота",
    download: "Скачано",
    upload: "Загружено",
    expireTime: "Срок действия",
    kickUtilTimeLast: "Оставшееся время офлайна",
    kickUtilTime: "Время офлайна",
    deviceNo: "Лимит устройств",
    onlineStatus: "Статус онлайн",
    online: "Онлайн",
    offline: "Офлайн",
    device: "Устройства онлайн",
    role: "Роль",
    unit: "Единица",
    loginAt: "Время последнего входа",
    conAt: "Время последнего подключения",
    createTime: "Время создания",
    releaseSuccess: "Разблокировка выполнена",
    kick: "Отключить",
    kickTip: "Принудительно отключить пользователя",
    releaseKick: "Снять блокировку",
    releaseKickTip: "Убрать статус офлайн",
  },
  config: {
    huiWebPort: "Веб-порт H UI",
    huiWebContext: "Веб-контекст H UI",
    hysteria2TrafficTime: "Время учета трафика Hysteria2",
    huiCrtPath: "Путь к CRT-файлу H UI",
    huiKeyPath: "Путь к KEY-файлу H UI",
    uploadCrtFile: "Загрузить CRT-файл",
    uploadKeyFile: "Загрузить KEY-файл",
    restartServer: "Перезапустить панель",
    restartTip: "Выполняется перезапуск, пожалуйста, обновите страницу",
    useHysteria2Cert: "Использовать сертификат Hysteria2",
    huiHttps: "Включить HTTPS для панели",
    resetTrafficCron: "Задача по расписанию для сброса трафика",
    resetTrafficCronTip:
      "Выражение cron для планировщика, справка: https://pkg.go.dev/github.com/robfig/cron/v3",
    resetTrafficMonth: "Запуск раз в месяц, в полночь первого числа",
    resetTrafficWeek:
      "Запуск раз в неделю, в полночь между субботой и воскресеньем",
  },
  monitor: {
    huiVersion: "Версия H UI",
    cpuPercent: "Использование CPU",
    memPercent: "Использование памяти",
    diskPercent: "Использование диска",
    hysteria2UserTotal: "Количество пользователей онлайн",
    hysteria2DeviceTotal: "Количество устройств онлайн",
    hysteria2Version: "Версия Hysteria2",
    hysteria2Running: "Статус Hysteria2",
    hysteria2RunningTrue: "Запущена",
    hysteria2RunningFalse: "Остановлена",
  },
  log: {
    numLine: "Количество строк",
  },
  telegram: {
    placeholder: "Плейсхолдер",
    enable: "Включить",
    disable: "Отключить",
    telegramEnable: "Включить/отключить",
    telegramToken: "Токен Telegram",
    telegramChatId: "Chat ID Telegram",
    telegramJob: "Список задач",
    telegramLoginJob: "Уведомление о входе",
    telegramLoginJobEnable: "Включить/отключить",
    telegramLoginJobText: "Шаблон содержимого",
  },
  hysteria: {
    enable: "Включить",
    disable: "Отключить",
    addConfigItem: "Добавить элемент конфигурации",
    hysteria2Version: "Версия Hysteria2",
    hysteria2Running: "Статус Hysteria2",
    hysteria2ChangeVersion: "Изменить",
    addOutbound: "Добавить outbound",
    extension: "Расширение",
    listen: "Listen",
    tls: "TLS",
    obfs: "Обфускация",
    quic: "Параметры QUIC",
    bandwidth: "Полоса пропускания",
    speedTest: "Тест скорости",
    udp: "UDP",
    resolver: "Резолвер",
    sniff: "Анализ протоколов",
    acl: "ACL",
    outbounds: "Исходящие каналы",
    http: "API статистики трафика (HTTP)",
    masquerade: "Маскировка",
    config: {
      enable: "Enable/Disable",
      remark: "Remark",
      portHopping:
        "Переключение между портами, несколько отдельных портов: 1234,5678,9012; Диапазон портов: 20000-50000; Комбинация обоих: 1234,5000-6000,7044,8000-9000",
      clashExtension: "Продление подписки на Clash",
      listen:
        "Если IP-адрес опущен, сервер будет слушать на всех интерфейсах, как IPv4, так и IPv6. Для прослушивания только IPv4 используйте 0.0.0.0:443. Для только IPv6 — [::]:443. Hysteria >= 2.8.0 поддерживает переключение портов.",
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
        initConnReceiveWindow: "Размер окна начального QUIC-соединения.",
        maxConnReceiveWindow: "Максимальный размер окна QUIC-соединения.",
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
