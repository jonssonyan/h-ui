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
    listen: "Прослушивание",
    tls: "TLS",
    obfs: "Обфускация",
    quic: "Параметры QUIC",
    bandwidth: "Пропускная способность",
    speedTest: "Тест скорости",
    udp: "UDP",
    resolver: "Резолвер",
    sniff: "Определение протокола",
    acl: "ACL",
    outbounds: "Outbounds",
    http: "API статистики трафика (HTTP)",
    masquerade: "Маскировка",
    config: {
      enable: "Включить/отключить",
      remark: "Примечание",
      portHopping:
        "Переключение портов. Несколько отдельных портов: 1234,5678,9012; диапазон портов: 20000-50000; комбинация: 1234,5000-6000,7044,8000-9000",
      clashExtension: "Расширение подписки Clash",
      listen:
        "Если IP-адрес не указан, сервер будет слушать на всех интерфейсах IPv4 и IPv6. Чтобы слушать только IPv4, используйте 0.0.0.0:443. Чтобы слушать только IPv6, используйте [::]:443. Hysteria >= 2.8.0 поддерживает переключение портов.",
      tlsType: "Тип TLS",
      tls: {
        cert: "Путь к файлу сертификата.",
        key: "Путь к файлу ключа.",
        sniGuard:
          'Проверять SNI, переданный клиентом. Соединение принимается только если он совпадает с указанным в сертификате. В противном случае TLS-рукопожатие будет прервано. Установите strict, чтобы принудительно включить это поведение. Установите disable, чтобы полностью отключить. Значение по умолчанию dns-san включает эту функцию только если сертификат содержит расширение "Subject Alternative Name" с доменным именем.',
      },
      acme: {
        domains: "Домены",
        email: "Email",
        ca: "Центр сертификации. Может быть letsencrypt или zerossl.",
        listenHost:
          "Адрес хоста (без порта) для ACME challenge. Если не указан, сервер будет слушать на всех интерфейсах.",
        dir: "Каталог для хранения ключа ACME-аккаунта и сертификатов.",
        type: "Тип ACME challenge. Может быть http, tls или dns.",
        http: {
          altPort:
            "Порт прослушивания для HTTP challenge. Примечание: если это не порт 80, потребуется проброс порта или HTTP reverse proxy, иначе проверка завершится ошибкой.",
        },
        tls: {
          altPort:
            "Порт прослушивания для TLS-ALPN challenge. Примечание: если это не порт 443, потребуется проброс порта или TLS reverse proxy, иначе проверка завершится ошибкой.",
        },
        dns: {
          name: "DNS-провайдер. Подробнее см. в настройке ACME DNS.",
          config: "Конфигурация ACME DNS",
        },
        disableHTTP: "Отключить HTTP challenge.",
        disableTLSALPN: "Отключить TLS-ALPN challenge.",
        altHTTPPort:
          "Альтернативный порт HTTP challenge. Примечание: если используется не 80, нужно настроить проброс порта или HTTP reverse proxy с 80 на этот порт, иначе ACME не сможет выпустить сертификат.",
        altTLSALPNPort:
          "Альтернативный порт TLS-ALPN challenge. Примечание: если используется не 443, нужно настроить проброс порта или SNI proxy с 443 на этот порт, иначе ACME не сможет выпустить сертификат.",
      },
      obfs: {
        type: "Тип",
        salamander: {
          password: "Замените на надежный пароль по вашему выбору.",
        },
      },
      quic: {
        initStreamReceiveWindow: "Начальный размер окна приема потока QUIC.",
        maxStreamReceiveWindow: "Максимальный размер окна приема потока QUIC.",
        initConnReceiveWindow:
          "Начальный размер окна приема соединения QUIC.",
        maxConnReceiveWindow:
          "Максимальный размер окна приема соединения QUIC.",
        maxIdleTimeout:
          "Максимальный тайм-аут простоя. Как долго сервер будет считать клиента подключенным без активности.",
        maxIncomingStreams:
          "Максимальное количество одновременных входящих потоков.",
        disablePathMTUDiscovery: "Отключить обнаружение Path MTU для QUIC.",
      },
      bandwidth: {
        up: "Вверх",
        down: "Вниз",
      },
      ignoreClientBandwidth:
        "Если включено, сервер будет игнорировать любые подсказки по пропускной способности, заданные клиентами",
      congestion: {
        type: "Тип",
        bbrProfile:
          "Это поле применяется только при значении type = bbr. По умолчанию используется standard.",
      },
      speedTest:
        "speedTest включает встроенный сервер тестирования скорости. После включения клиенты смогут проверять скорость загрузки и отдачи через сервер. Подробнее см. документацию по Speed Test.",
      disableUDP:
        "disableUDP отключает переадресацию UDP, разрешая только TCP-соединения.",
      udpIdleTimeout:
        "udpIdleTimeout задает время, в течение которого сервер будет держать локальный UDP-порт открытым для каждой неактивной UDP-сессии. Концептуально это аналог тайм-аута UDP-сессии NAT.",
      resolver: {
        type: "Тип",
        tcp: {
          addr: "Адрес TCP-резолвера.",
          timeout: "Тайм-аут DNS-запросов.",
        },
        udp: {
          addr: "Адрес UDP-резолвера.",
          timeout: "Тайм-аут DNS-запросов.",
        },
        tls: {
          addr: "Адрес TLS-резолвера.",
          timeout: "Тайм-аут DNS-запросов.",
          sni: "SNI для TLS-резолвера.",
          insecure: "Отключить проверку TLS для TLS-резолвера.",
        },
        https: {
          addr: "Адрес HTTPS-резолвера.",
          timeout: "Тайм-аут DNS-запросов.",
          sni: "SNI для TLS-резолвера.",
          insecure: "Отключить проверку TLS для TLS-резолвера.",
        },
      },
      sniff: {
        enable: "Включить определение протокола.",
        timeout:
          "Тайм-аут определения. Если протокол или домен не удается определить за это время, для соединения будет использован исходный адрес.",
        rewriteDomain:
          "Перезаписывать ли запросы, уже использующие доменное имя. Если включено, такие запросы все равно будут анализироваться.",
        tcpPorts:
          "Список TCP-портов. Анализироваться будут только TCP-запросы на этих портах.",
        udpPorts:
          "Список UDP-портов. Анализироваться будут только UDP-запросы на этих портах.",
      },
      aclType: "Тип ACL",
      acl: {
        file: "Путь к файлу ACL.",
        inline: "Список встроенных правил ACL.",
        geoip:
          "Необязательно. Раскомментируйте для включения. Путь к файлу базы данных GeoIP. Если поле не указано, Hysteria автоматически скачает последнюю базу в рабочий каталог.",
        geosite:
          "Необязательно. Раскомментируйте для включения. Путь к файлу базы данных GeoSite. Если поле не указано, Hysteria автоматически скачает последнюю базу в рабочий каталог.",
        geoUpdateInterval:
          "Необязательно. Интервал обновления баз GeoIP/GeoSite. По умолчанию 168 часов (1 неделя). Применяется только если базы GeoIP/GeoSite скачиваются автоматически.",
      },
      outbounds: {
        name: "Имя outbound. Используется в ACL-правилах.",
        type: "Тип",
        socks5: {
          addr: "Адрес прокси-сервера SOCKS5.",
          username:
            "Необязательно. Имя пользователя для SOCKS5-прокси, если требуется аутентификация.",
          password:
            "Необязательно. Пароль для SOCKS5-прокси, если требуется аутентификация.",
        },
        http: {
          url: "URL-адрес прокси-сервера HTTP/HTTPS. (Может быть http:// или https://)",
          insecure:
            "Необязательно. Отключить ли проверку TLS. Применяется только к HTTPS-прокси.",
        },
        direct: {
          mode: "Тип",
          bindIPv4: "Локальный IPv4-адрес для привязки.",
          bindIPv6: "Локальный IPv6-адрес для привязки.",
          bindDevice: "Локальный сетевой интерфейс для привязки.",
          fastOpen: "Включить TCP Fast Open.",
        },
      },
      trafficStats: {
        listen: "Адрес для прослушивания.",
      },
      masquerade: {
        type: "Тип",
        file: {
          dir: "Каталог, из которого будут раздаваться файлы.",
        },
        proxy: {
          url: "URL сайта для проксирования.",
          rewriteHost:
            "Перезаписывать ли заголовок Host в соответствии с проксируемым сайтом. Это необходимо, если целевой веб-сервер использует Host для определения обслуживаемого сайта.",
          insecure: "Отключить проверку TLS для проксируемого сайта.",
          xForwarded:
            "Необязательно. Устанавливать ли заголовки X-Forwarded-For, X-Forwarded-Host и X-Forwarded-Proto при проксировании. По умолчанию отключено.",
        },
        string: {
          content: "Строка, которая будет возвращена.",
          headers: "Необязательно. Заголовки, которые будут возвращены.",
          statusCode: "Необязательно. HTTP-статус ответа. По умолчанию 200.",
        },
        listenHTTP: "Адрес прослушивания HTTP (TCP).",
        listenHTTPS: "Адрес прослушивания HTTPS (TCP).",
        forceHTTPS:
          "Принудительно использовать HTTPS. Если включено, все HTTP-запросы будут перенаправляться на HTTPS.",
      },
    },
  },
};
