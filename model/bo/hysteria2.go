package bo

import "time"

type Hysteria2ServerConfig struct {
	Listen                string                      `yaml:"listen" json:"listen"`
	Obfs                  serverConfigObfs            `yaml:"obfs" json:"obfs"`
	TLS                   serverConfigTLS             `yaml:"tls" json:"tls"`
	ACME                  serverConfigACME            `yaml:"acme" json:"acme"`
	QUIC                  serverConfigQUIC            `yaml:"quic" json:"quic"`
	Bandwidth             serverConfigBandwidth       `yaml:"bandwidth" json:"bandwidth"`
	IgnoreClientBandwidth bool                        `yaml:"ignoreClientBandwidth" json:"ignoreClientBandwidth"`
	SpeedTest             bool                        `yaml:"speedTest" json:"speedTest"`
	DisableUDP            bool                        `yaml:"disableUDP" json:"disableUDP"`
	UDPIdleTimeout        time.Duration               `yaml:"udpIdleTimeout" json:"udpIdleTimeout"`
	Auth                  serverConfigAuth            `yaml:"auth" json:"auth"`
	Resolver              serverConfigResolver        `yaml:"resolver" json:"resolver"`
	ACL                   serverConfigACL             `yaml:"acl" json:"acl"`
	Outbounds             []serverConfigOutboundEntry `yaml:"outbounds" json:"outbounds"`
	TrafficStats          serverConfigTrafficStats    `yaml:"trafficStats" json:"trafficStats"`
	Masquerade            serverConfigMasquerade      `yaml:"masquerade" json:"masquerade"`
}

type serverConfigObfsSalamander struct {
	Password string `yaml:"password" json:"password"`
}

type serverConfigObfs struct {
	Type       string                     `yaml:"type" json:"type"`
	Salamander serverConfigObfsSalamander `yaml:"salamander" json:"salamander"`
}

type serverConfigTLS struct {
	Cert string `yaml:"cert" json:"cert"`
	Key  string `yaml:"key"  json:"key"`
}

type serverConfigACME struct {
	Domains        []string `yaml:"domains" json:"domains"`
	Email          string   `yaml:"email" json:"email"`
	CA             string   `yaml:"ca" json:"ca"`
	DisableHTTP    bool     `yaml:"disableHTTP" json:"disableHTTP"`
	DisableTLSALPN bool     `yaml:"disableTLSALPN" json:"disableTLSALPN"`
	ListenHost     string   `yaml:"listenHost" json:"listenHost"`
	AltHTTPPort    int      `yaml:"altHTTPPort" json:"altHTTPPort"`
	AltTLSALPNPort int      `yaml:"altTLSALPNPort" json:"altTLSALPNPort"`
	Dir            string   `yaml:"dir" json:"dir"`
}

type serverConfigQUIC struct {
	InitStreamReceiveWindow     uint64        `yaml:"initStreamReceiveWindow" json:"initStreamReceiveWindow"`
	MaxStreamReceiveWindow      uint64        `yaml:"maxStreamReceiveWindow" json:"maxStreamReceiveWindow"`
	InitConnectionReceiveWindow uint64        `yaml:"initConnReceiveWindow" json:"initConnReceiveWindow"`
	MaxConnectionReceiveWindow  uint64        `yaml:"maxConnReceiveWindow" json:"maxConnReceiveWindow"`
	MaxIdleTimeout              time.Duration `yaml:"maxIdleTimeout" json:"maxIdleTimeout"`
	MaxIncomingStreams          int64         `yaml:"maxIncomingStreams" json:"maxIncomingStreams"`
	DisablePathMTUDiscovery     bool          `yaml:"disablePathMTUDiscovery" json:"disablePathMTUDiscovery"`
}

type serverConfigBandwidth struct {
	Up   string `yaml:"up" json:"up"`
	Down string `yaml:"down" json:"down"`
}

type serverConfigAuthHTTP struct {
	URL      string `yaml:"url" json:"url"`
	Insecure bool   `yaml:"insecure" json:"insecure"`
}

type serverConfigAuth struct {
	Type     string               `yaml:"type" json:"type"`
	Password string               `yaml:"password" json:"password"`
	UserPass map[string]string    `yaml:"userpass" json:"userpass"`
	HTTP     serverConfigAuthHTTP `yaml:"http" json:"http"`
	Command  string               `yaml:"command" json:"command"`
}

type serverConfigResolverTCP struct {
	Addr    string        `yaml:"addr" json:"addr"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
}

type serverConfigResolverUDP struct {
	Addr    string        `yaml:"addr" json:"addr"`
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
}

type serverConfigResolverTLS struct {
	Addr     string        `yaml:"addr" json:"addr"`
	Timeout  time.Duration `yaml:"timeout" json:"timeout"`
	SNI      string        `yaml:"sni" json:"sni"`
	Insecure bool          `yaml:"insecure" json:"insecure"`
}

type serverConfigResolverHTTPS struct {
	Addr     string        `yaml:"addr" json:"addr"`
	Timeout  time.Duration `yaml:"timeout" json:"timeout"`
	SNI      string        `yaml:"sni" json:"sni"`
	Insecure bool          `yaml:"insecure" json:"insecure"`
}

type serverConfigResolver struct {
	Type  string                    `yaml:"type" json:"type"`
	TCP   serverConfigResolverTCP   `yaml:"tcp" json:"tcp"`
	UDP   serverConfigResolverUDP   `yaml:"udp" json:"udp"`
	TLS   serverConfigResolverTLS   `yaml:"tls" json:"tls"`
	HTTPS serverConfigResolverHTTPS `yaml:"https" json:"https"`
}

type serverConfigACL struct {
	File              string        `yaml:"file" json:"file"`
	Inline            []string      `yaml:"inline" json:"inline"`
	GeoIP             string        `yaml:"geoip" json:"geoip"`
	GeoSite           string        `yaml:"geosite" json:"geosite"`
	GeoUpdateInterval time.Duration `yaml:"geoUpdateInterval" json:"geoUpdateInterval"`
}

type serverConfigOutboundDirect struct {
	Mode       string `yaml:"mode" json:"mode"`
	BindIPv4   string `yaml:"bindIPv4" json:"bindIPv4"`
	BindIPv6   string `yaml:"bindIPv6" json:"bindIPv6"`
	BindDevice string `yaml:"bindDevice" json:"bindDevice"`
}

type serverConfigOutboundSOCKS5 struct {
	Addr     string `yaml:"addr" json:"addr"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type serverConfigOutboundHTTP struct {
	URL      string `yaml:"url" json:"url"`
	Insecure bool   `yaml:"insecure" json:"insecure"`
}

type serverConfigOutboundEntry struct {
	Name   string                     `yaml:"name" json:"name"`
	Type   string                     `yaml:"type" json:"type"`
	Direct serverConfigOutboundDirect `yaml:"direct" json:"direct"`
	SOCKS5 serverConfigOutboundSOCKS5 `yaml:"socks5" json:"socks5"`
	HTTP   serverConfigOutboundHTTP   `yaml:"http" json:"http"`
}

type serverConfigTrafficStats struct {
	Listen string `yaml:"listen" json:"listen"`
	Secret string `yaml:"secret" json:"secret"`
}

type serverConfigMasqueradeFile struct {
	Dir string `yaml:"dir" json:"dir"`
}

type serverConfigMasqueradeProxy struct {
	URL         string `yaml:"url" json:"url"`
	RewriteHost bool   `yaml:"rewriteHost" json:"rewriteHost"`
}

type serverConfigMasqueradeString struct {
	Content    string            `yaml:"content" json:"content"`
	Headers    map[string]string `yaml:"headers" json:"headers"`
	StatusCode int               `yaml:"statusCode" json:"statusCode"`
}

type serverConfigMasquerade struct {
	Type        string                       `yaml:"type" json:"type"`
	File        serverConfigMasqueradeFile   `yaml:"file" json:"file"`
	Proxy       serverConfigMasqueradeProxy  `yaml:"proxy" json:"proxy"`
	String      serverConfigMasqueradeString `yaml:"string" json:"string"`
	ListenHTTP  string                       `yaml:"listenHTTP" json:"listenHTTP"`
	ListenHTTPS string                       `yaml:"listenHTTPS" json:"listenHTTPS"`
	ForceHTTPS  bool                         `yaml:"forceHTTPS" json:"forceHTTPS"`
}
