export interface ConfigDto {
  key: string;
}

export interface ConfigsDto {
  keys: Array<string>;
}

export interface ConfigVo {
  key: string;
  value: string;
}

export interface ConfigsUpdateDto {
  key: string;
  value: string;
}

export interface ConfigUpdateDto {
  configUpdateDtos: Array<ConfigsUpdateDto>;
}

export interface Hysteria2ServerConfig {
  listen: string;
  tls: {
    cert: string;
    key: string;
  };
  acme: {
    domains: string[];
    email: string;
    ca: string;
    disableHTTP: string;
    disableTLSALPN: string;
    altHTTPPort: string;
    altTLSALPNPort: string;
    dir: string;
    listenHost: string;
  };
  obfs: {
    type: string;
    salamander: {
      password: string;
    };
  };
  quic: {
    initStreamReceiveWindow: string;
    maxStreamReceiveWindow: string;
    initConnReceiveWindow: string;
    maxConnReceiveWindow: string;
    maxIdleTimeout: string;
    maxIncomingStreams: string;
    disablePathMTUDiscovery: string;
  };
  bandwidth: {
    up: string;
    down: string;
  };
  ignoreClientBandwidth: string;
  speedTest: string;
  disableUDP: string;
  udpIdleTimeout: string;
  auth: {
    type: string;
    password: string;
    userpass: { [key: string]: string };
    http: {
      url: string;
      insecure: string;
    };
    command: string;
  };
  resolver: {
    type: string;
    tcp: {
      addr: string;
      timeout: string;
    };
    udp: {
      addr: string;
      timeout: string;
    };
    tls: {
      addr: string;
      timeout: string;
      sni: string;
      insecure: string;
    };
    https: {
      addr: string;
      timeout: string;
      sni: string;
      insecure: string;
    };
  };
  acl: {
    file: string;
    inline: string[];
    geoip: string;
    geosite: string;
    geoUpdateInterval: string;
  };
  outbounds: {
    name: string;
    type: string;
    socks5: {
      addr: string;
      username: string;
      password: string;
    };
    http: {
      url: string;
      insecure: string;
    };
    direct: {
      mode: string;
      bindIPv4: string;
      bindIPv6: string;
      bindDevice: string;
    };
  }[];
  trafficStats: {
    listen: string;
    secret: string;
  };
  masquerade: {
    type: string;
    file: {
      dir: string;
    };
    proxy: {
      url: string;
      rewriteHost: string;
    };
    string: {
      content: string;
      headers: { [key: string]: string };
      statusCode: string;
    };
    listenHTTP: string;
    listenHTTPS: string;
    forceHTTPS: string;
  };
}

export const defaultHysteria2ServerConfig: Hysteria2ServerConfig = {
  listen: "",
  tls: {
    cert: "",
    key: "",
  },
  acme: {
    domains: [],
    email: "",
    ca: "",
    disableHTTP: "",
    disableTLSALPN: "",
    altHTTPPort: "",
    altTLSALPNPort: "",
    dir: "",
    listenHost: "",
  },
  obfs: {
    type: "",
    salamander: {
      password: "",
    },
  },
  quic: {
    initStreamReceiveWindow: "",
    maxStreamReceiveWindow: "",
    initConnReceiveWindow: "",
    maxConnReceiveWindow: "",
    maxIdleTimeout: "",
    maxIncomingStreams: "",
    disablePathMTUDiscovery: "",
  },
  bandwidth: {
    up: "",
    down: "",
  },
  ignoreClientBandwidth: "",
  speedTest: "",
  disableUDP: "",
  udpIdleTimeout: "",
  auth: {
    type: "",
    password: "",
    userpass: {},
    http: {
      url: "",
      insecure: "",
    },
    command: "",
  },
  resolver: {
    type: "",
    tcp: {
      addr: "",
      timeout: "",
    },
    udp: {
      addr: "",
      timeout: "",
    },
    tls: {
      addr: "",
      timeout: "",
      sni: "",
      insecure: "",
    },
    https: {
      addr: "",
      timeout: "",
      sni: "",
      insecure: "",
    },
  },
  acl: {
    file: "",
    inline: [],
    geoip: "",
    geosite: "",
    geoUpdateInterval: "",
  },
  outbounds: [
    {
      name: "",
      type: "",
      socks5: {
        addr: "",
        username: "",
        password: "",
      },
      http: {
        url: "",
        insecure: "",
      },
      direct: {
        mode: "",
        bindIPv4: "",
        bindIPv6: "",
        bindDevice: "",
      },
    },
  ],
  trafficStats: {
    listen: "",
    secret: "",
  },
  masquerade: {
    type: "",
    file: {
      dir: "",
    },
    proxy: {
      url: "",
      rewriteHost: "",
    },
    string: {
      content: "",
      headers: {},
      statusCode: "",
    },
    listenHTTP: "",
    listenHTTPS: "",
    forceHTTPS: "",
  },
};
