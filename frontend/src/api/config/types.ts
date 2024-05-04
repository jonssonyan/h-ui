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
  obfs: {
    type: string;
    salamander: {
      password: string;
    };
  };
  tls: {
    cert: string;
    key: string;
  };
  acme: {
    domains: string[];
    email: string;
    ca: string;
    disableHTTP: boolean;
    disableTLSALPN: boolean;
    listenHost: string;
    altHTTPPort: number;
    altTLSALPNPort: number;
    dir: string;
  };
  quic: {
    initStreamReceiveWindow: number;
    maxStreamReceiveWindow: number;
    initConnectionReceiveWindow: number;
    maxConnectionReceiveWindow: number;
    maxIdleTimeout: number;
    maxIncomingStreams: number;
    disablePathMTUDiscovery: boolean;
  };
  bandwidth: {
    up: string;
    down: string;
  };
  ignoreClientBandwidth: boolean;
  speedTest: boolean;
  disableUDP: boolean;
  udpIdleTimeout: number;
  auth: {
    type: string;
    password: string;
    userpass: { [key: string]: string };
    http: {
      url: string;
      insecure: boolean;
    };
    command: string;
  };
  resolver: {
    type: string;
    tcp: {
      addr: string;
      timeout: number;
    };
    udp: {
      addr: string;
      timeout: number;
    };
    tls: {
      addr: string;
      timeout: number;
      sni: string;
      insecure: boolean;
    };
    https: {
      addr: string;
      timeout: number;
      sni: string;
      insecure: boolean;
    };
  };
  acl: {
    file: string;
    inline: string[];
    geoip: string;
    geosite: string;
    geoUpdateInterval: number;
  };
  outbounds: {
    name: string;
    type: string;
    direct: {
      mode: string;
      bindIPv4: string;
      bindIPv6: string;
      bindDevice: string;
    };
    socks5: {
      addr: string;
      username: string;
      password: string;
    };
    http: {
      url: string;
      insecure: boolean;
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
      rewriteHost: boolean;
    };
    string: {
      content: string;
      headers: { [key: string]: string };
      statusCode: number;
    };
    listenHTTP: string;
    listenHTTPS: string;
    forceHTTPS: boolean;
  };
}

export const defaultHysteria2ServerConfig: Hysteria2ServerConfig = {
  listen: "443",
  obfs: {
    type: "",
    salamander: {
      password: "",
    },
  },
  tls: {
    cert: "",
    key: "",
  },
  acme: {
    domains: [],
    email: "",
    ca: "",
    disableHTTP: false,
    disableTLSALPN: false,
    listenHost: "",
    altHTTPPort: 0,
    altTLSALPNPort: 0,
    dir: "",
  },
  quic: {
    initStreamReceiveWindow: 0,
    maxStreamReceiveWindow: 0,
    initConnectionReceiveWindow: 0,
    maxConnectionReceiveWindow: 0,
    maxIdleTimeout: 0,
    maxIncomingStreams: 0,
    disablePathMTUDiscovery: false,
  },
  bandwidth: {
    up: "",
    down: "",
  },
  ignoreClientBandwidth: false,
  speedTest: false,
  disableUDP: false,
  udpIdleTimeout: 0,
  auth: {
    type: "",
    password: "",
    userpass: {},
    http: {
      url: "",
      insecure: false,
    },
    command: "",
  },
  resolver: {
    type: "",
    tcp: {
      addr: "",
      timeout: 0,
    },
    udp: {
      addr: "",
      timeout: 0,
    },
    tls: {
      addr: "",
      timeout: 0,
      sni: "",
      insecure: false,
    },
    https: {
      addr: "",
      timeout: 0,
      sni: "",
      insecure: false,
    },
  },
  acl: {
    file: "",
    inline: [],
    geoip: "",
    geosite: "",
    geoUpdateInterval: 0,
  },
  outbounds: [],
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
      rewriteHost: false,
    },
    string: {
      content: "",
      headers: {},
      statusCode: 0,
    },
    listenHTTP: "",
    listenHTTPS: "",
    forceHTTPS: false,
  },
};
