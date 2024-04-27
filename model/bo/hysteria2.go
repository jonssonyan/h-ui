package bo

import "time"

type Hysteria2ServerConfig struct {
	Listen                string                      `mapstructure:"listen"`
	Obfs                  serverConfigObfs            `mapstructure:"obfs"`
	TLS                   *serverConfigTLS            `mapstructure:"tls"`
	ACME                  *serverConfigACME           `mapstructure:"acme"`
	QUIC                  serverConfigQUIC            `mapstructure:"quic"`
	Bandwidth             serverConfigBandwidth       `mapstructure:"bandwidth"`
	IgnoreClientBandwidth bool                        `mapstructure:"ignoreClientBandwidth"`
	SpeedTest             bool                        `mapstructure:"speedTest"`
	DisableUDP            bool                        `mapstructure:"disableUDP"`
	UDPIdleTimeout        time.Duration               `mapstructure:"udpIdleTimeout"`
	Auth                  serverConfigAuth            `mapstructure:"auth"`
	Resolver              serverConfigResolver        `mapstructure:"resolver"`
	ACL                   serverConfigACL             `mapstructure:"acl"`
	Outbounds             []serverConfigOutboundEntry `mapstructure:"outbounds"`
	TrafficStats          serverConfigTrafficStats    `mapstructure:"trafficStats"`
	Masquerade            serverConfigMasquerade      `mapstructure:"masquerade"`
}

type serverConfigObfsSalamander struct {
	Password string `mapstructure:"password"`
}

type serverConfigObfs struct {
	Type       string                     `mapstructure:"type"`
	Salamander serverConfigObfsSalamander `mapstructure:"salamander"`
}

type serverConfigTLS struct {
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
}

type serverConfigACME struct {
	Domains        []string `mapstructure:"domains"`
	Email          string   `mapstructure:"email"`
	CA             string   `mapstructure:"ca"`
	DisableHTTP    bool     `mapstructure:"disableHTTP"`
	DisableTLSALPN bool     `mapstructure:"disableTLSALPN"`
	ListenHost     string   `mapstructure:"listenHost"`
	AltHTTPPort    int      `mapstructure:"altHTTPPort"`
	AltTLSALPNPort int      `mapstructure:"altTLSALPNPort"`
	Dir            string   `mapstructure:"dir"`
}

type serverConfigQUIC struct {
	InitStreamReceiveWindow     uint64        `mapstructure:"initStreamReceiveWindow"`
	MaxStreamReceiveWindow      uint64        `mapstructure:"maxStreamReceiveWindow"`
	InitConnectionReceiveWindow uint64        `mapstructure:"initConnReceiveWindow"`
	MaxConnectionReceiveWindow  uint64        `mapstructure:"maxConnReceiveWindow"`
	MaxIdleTimeout              time.Duration `mapstructure:"maxIdleTimeout"`
	MaxIncomingStreams          int64         `mapstructure:"maxIncomingStreams"`
	DisablePathMTUDiscovery     bool          `mapstructure:"disablePathMTUDiscovery"`
}

type serverConfigBandwidth struct {
	Up   string `mapstructure:"up"`
	Down string `mapstructure:"down"`
}

type serverConfigAuthHTTP struct {
	URL      string `mapstructure:"url"`
	Insecure bool   `mapstructure:"insecure"`
}

type serverConfigAuth struct {
	Type     string               `mapstructure:"type"`
	Password string               `mapstructure:"password"`
	UserPass map[string]string    `mapstructure:"userpass"`
	HTTP     serverConfigAuthHTTP `mapstructure:"http"`
	Command  string               `mapstructure:"command"`
}

type serverConfigResolverTCP struct {
	Addr    string        `mapstructure:"addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type serverConfigResolverUDP struct {
	Addr    string        `mapstructure:"addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type serverConfigResolverTLS struct {
	Addr     string        `mapstructure:"addr"`
	Timeout  time.Duration `mapstructure:"timeout"`
	SNI      string        `mapstructure:"sni"`
	Insecure bool          `mapstructure:"insecure"`
}

type serverConfigResolverHTTPS struct {
	Addr     string        `mapstructure:"addr"`
	Timeout  time.Duration `mapstructure:"timeout"`
	SNI      string        `mapstructure:"sni"`
	Insecure bool          `mapstructure:"insecure"`
}

type serverConfigResolver struct {
	Type  string                    `mapstructure:"type"`
	TCP   serverConfigResolverTCP   `mapstructure:"tcp"`
	UDP   serverConfigResolverUDP   `mapstructure:"udp"`
	TLS   serverConfigResolverTLS   `mapstructure:"tls"`
	HTTPS serverConfigResolverHTTPS `mapstructure:"https"`
}

type serverConfigACL struct {
	File              string        `mapstructure:"file"`
	Inline            []string      `mapstructure:"inline"`
	GeoIP             string        `mapstructure:"geoip"`
	GeoSite           string        `mapstructure:"geosite"`
	GeoUpdateInterval time.Duration `mapstructure:"geoUpdateInterval"`
}

type serverConfigOutboundDirect struct {
	Mode       string `mapstructure:"mode"`
	BindIPv4   string `mapstructure:"bindIPv4"`
	BindIPv6   string `mapstructure:"bindIPv6"`
	BindDevice string `mapstructure:"bindDevice"`
}

type serverConfigOutboundSOCKS5 struct {
	Addr     string `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type serverConfigOutboundHTTP struct {
	URL      string `mapstructure:"url"`
	Insecure bool   `mapstructure:"insecure"`
}

type serverConfigOutboundEntry struct {
	Name   string                     `mapstructure:"name"`
	Type   string                     `mapstructure:"type"`
	Direct serverConfigOutboundDirect `mapstructure:"direct"`
	SOCKS5 serverConfigOutboundSOCKS5 `mapstructure:"socks5"`
	HTTP   serverConfigOutboundHTTP   `mapstructure:"http"`
}

type serverConfigTrafficStats struct {
	Listen string `mapstructure:"listen"`
	Secret string `mapstructure:"secret"`
}

type serverConfigMasqueradeFile struct {
	Dir string `mapstructure:"dir"`
}

type serverConfigMasqueradeProxy struct {
	URL         string `mapstructure:"url"`
	RewriteHost bool   `mapstructure:"rewriteHost"`
}

type serverConfigMasqueradeString struct {
	Content    string            `mapstructure:"content"`
	Headers    map[string]string `mapstructure:"headers"`
	StatusCode int               `mapstructure:"statusCode"`
}

type serverConfigMasquerade struct {
	Type        string                       `mapstructure:"type"`
	File        serverConfigMasqueradeFile   `mapstructure:"file"`
	Proxy       serverConfigMasqueradeProxy  `mapstructure:"proxy"`
	String      serverConfigMasqueradeString `mapstructure:"string"`
	ListenHTTP  string                       `mapstructure:"listenHTTP"`
	ListenHTTPS string                       `mapstructure:"listenHTTPS"`
	ForceHTTPS  bool                         `mapstructure:"forceHTTPS"`
}
