package bo

type Hysteria2ServerConfig struct {
	Listen                *string                     `yaml:"listen,omitempty" json:"listen" validate:"required"`
	Obfs                  *serverConfigObfs           `yaml:"obfs,omitempty" json:"obfs" validate:"omitempty"`
	TLS                   *serverConfigTLS            `yaml:"tls,omitempty" json:"tls" validate:"omitempty"`
	ACME                  *serverConfigACME           `yaml:"acme,omitempty" json:"acme" validate:"omitempty"`
	QUIC                  *serverConfigQUIC           `yaml:"quic,omitempty" json:"quic" validate:"omitempty"`
	Bandwidth             *serverConfigBandwidth      `yaml:"bandwidth,omitempty" json:"bandwidth" validate:"omitempty"`
	IgnoreClientBandwidth *bool                       `yaml:"ignoreClientBandwidth,omitempty" json:"ignoreClientBandwidth" validate:"omitempty"`
	SpeedTest             *bool                       `yaml:"speedTest,omitempty" json:"speedTest" validate:"omitempty"`
	DisableUDP            *bool                       `yaml:"disableUDP,omitempty" json:"disableUDP" validate:"omitempty"`
	UDPIdleTimeout        *string                     `yaml:"udpIdleTimeout,omitempty" json:"udpIdleTimeout" validate:"omitempty"`
	Auth                  *ServerConfigAuth           `yaml:"auth,omitempty" json:"-" validate:"omitempty"`
	Resolver              *serverConfigResolver       `yaml:"resolver,omitempty" json:"resolver" validate:"omitempty"`
	Sniff                 *serverConfigSniff          `yaml:"sniff,omitempty" json:"sniff" validate:"omitempty"`
	ACL                   *serverConfigACL            `yaml:"acl,omitempty" json:"acl" validate:"omitempty"`
	Outbounds             []serverConfigOutboundEntry `yaml:"outbounds,omitempty" json:"outbounds" validate:"omitempty"`
	TrafficStats          *ServerConfigTrafficStats   `yaml:"trafficStats,omitempty" json:"trafficStats" validate:"required"`
	Masquerade            *serverConfigMasquerade     `yaml:"masquerade,omitempty" json:"masquerade" validate:"omitempty"`
}

type serverConfigObfsSalamander struct {
	Password *string `yaml:"password,omitempty" json:"password" validate:"required"`
}

type serverConfigObfs struct {
	Type       *string                     `yaml:"type,omitempty" json:"type" validate:"required"`
	Salamander *serverConfigObfsSalamander `yaml:"salamander,omitempty" json:"salamander" validate:"required"`
}

type serverConfigTLS struct {
	Cert *string `yaml:"cert,omitempty" json:"cert" validate:"required"`
	Key  *string `yaml:"key,omitempty"  json:"key" validate:"required"`
}

type serverConfigACME struct {
	// Common fields
	Domains    []string `yaml:"domains,omitempty" json:"domains" validate:"required"`
	Email      *string  `yaml:"email,omitempty" json:"email" validate:"required"`
	CA         *string  `yaml:"ca,omitempty" json:"ca" validate:"required"`
	ListenHost *string  `yaml:"listenHost,omitempty" json:"listenHost" validate:"required"`
	Dir        *string  `yaml:"dir,omitempty" json:"dir" validate:"required"`

	// Type selection
	Type *string               `yaml:"type,omitempty" json:"type" validate:"omitempty"`
	HTTP *serverConfigACMEHTTP `yaml:"http,omitempty" json:"http" validate:"omitempty"`
	TLS  *serverConfigACMETLS  `yaml:"tls,omitempty" json:"tls" validate:"omitempty"`
	DNS  *serverConfigACMEDNS  `yaml:"dns,omitempty" json:"dns" validate:"omitempty"`

	// Legacy fields for backwards compatibility
	// Only applicable when Type is empty
	DisableHTTP    *bool `yaml:"disableHTTP,omitempty" json:"disableHTTP" validate:"required"`
	DisableTLSALPN *bool `yaml:"disableTLSALPN,omitempty" json:"disableTLSALPN" validate:"required"`
	AltHTTPPort    *int  `yaml:"altHTTPPort,omitempty" json:"altHTTPPort" validate:"required"`
	AltTLSALPNPort *int  `yaml:"altTLSALPNPort,omitempty" json:"altTLSALPNPort" validate:"required"`
}

type serverConfigACMEHTTP struct {
	AltPort *int `yaml:"altPort,omitempty" json:"altPort" validate:"required"`
}

type serverConfigACMETLS struct {
	AltPort *int `yaml:"altPort,omitempty" json:"altPort" validate:"required"`
}

type serverConfigACMEDNS struct {
	Name   *string           `yaml:"name,omitempty" json:"name" validate:"required"`
	Config map[string]string `yaml:"config,omitempty" json:"config" validate:"required"`
}

type serverConfigQUIC struct {
	InitStreamReceiveWindow     *uint64 `yaml:"initStreamReceiveWindow,omitempty" json:"initStreamReceiveWindow" validate:"omitempty"`
	MaxStreamReceiveWindow      *uint64 `yaml:"maxStreamReceiveWindow,omitempty" json:"maxStreamReceiveWindow" validate:"omitempty"`
	InitConnectionReceiveWindow *uint64 `yaml:"initConnReceiveWindow,omitempty" json:"initConnReceiveWindow" validate:"omitempty"`
	MaxConnectionReceiveWindow  *uint64 `yaml:"maxConnReceiveWindow,omitempty" json:"maxConnReceiveWindow" validate:"omitempty"`
	MaxIdleTimeout              *string `yaml:"maxIdleTimeout,omitempty" json:"maxIdleTimeout" validate:"omitempty"`
	MaxIncomingStreams          *int64  `yaml:"maxIncomingStreams,omitempty" json:"maxIncomingStreams" validate:"omitempty"`
	DisablePathMTUDiscovery     *bool   `yaml:"disablePathMTUDiscovery,omitempty" json:"disablePathMTUDiscovery" validate:"omitempty"`
}

type serverConfigBandwidth struct {
	Up   *string `yaml:"up,omitempty" json:"up" validate:"required"`
	Down *string `yaml:"down,omitempty" json:"down" validate:"required"`
}

type ServerConfigAuthHTTP struct {
	URL      *string `yaml:"url,omitempty" json:"url" validate:"required"`
	Insecure *bool   `yaml:"insecure,omitempty" json:"insecure" validate:"required"`
}

type ServerConfigAuth struct {
	Type     *string               `yaml:"type,omitempty" json:"type" validate:"omitempty"`
	Password *string               `yaml:"password,omitempty" json:"password" validate:"omitempty"`
	UserPass map[string]string     `yaml:"userpass,omitempty" json:"userpass" validate:"omitempty"`
	HTTP     *ServerConfigAuthHTTP `yaml:"http,omitempty" json:"http" validate:"omitempty"`
	Command  *string               `yaml:"command,omitempty" json:"command" validate:"omitempty"`
}

type serverConfigResolverTCP struct {
	Addr    *string `yaml:"addr,omitempty" json:"addr" validate:"required"`
	Timeout *string `yaml:"timeout,omitempty" json:"timeout" validate:"required"`
}

type serverConfigResolverUDP struct {
	Addr    *string `yaml:"addr,omitempty" json:"addr" validate:"required"`
	Timeout *string `yaml:"timeout,omitempty" json:"timeout" validate:"required"`
}

type serverConfigResolverTLS struct {
	Addr     *string `yaml:"addr,omitempty" json:"addr" validate:"required"`
	Timeout  *string `yaml:"timeout,omitempty" json:"timeout" validate:"required"`
	SNI      *string `yaml:"sni,omitempty" json:"sni" validate:"required"`
	Insecure *bool   `yaml:"insecure,omitempty" json:"insecure" validate:"required"`
}

type serverConfigResolverHTTPS struct {
	Addr     *string `yaml:"addr,omitempty" json:"addr" validate:"required"`
	Timeout  *string `yaml:"timeout,omitempty" json:"timeout" validate:"required"`
	SNI      *string `yaml:"sni,omitempty" json:"sni"  validate:"required"`
	Insecure *bool   `yaml:"insecure,omitempty" json:"insecure" validate:"required"`
}

type serverConfigResolver struct {
	Type  *string                    `yaml:"type,omitempty" json:"type" validate:"omitempty"`
	TCP   *serverConfigResolverTCP   `yaml:"tcp,omitempty" json:"tcp" validate:"omitempty"`
	UDP   *serverConfigResolverUDP   `yaml:"udp,omitempty" json:"udp" validate:"omitempty"`
	TLS   *serverConfigResolverTLS   `yaml:"tls,omitempty" json:"tls" validate:"omitempty"`
	HTTPS *serverConfigResolverHTTPS `yaml:"https,omitempty" json:"https" validate:"omitempty"`
}

type serverConfigSniff struct {
	Enable        *bool   `yaml:"enable,omitempty" json:"enable" validate:"required"`
	Timeout       *string `yaml:"timeout,omitempty" json:"timeout" validate:"required"`
	RewriteDomain *bool   `yaml:"rewriteDomain,omitempty" json:"rewriteDomain" validate:"required"`
	TCPPorts      *string `yaml:"tcpPorts,omitempty" json:"tcpPorts" validate:"omitempty"`
	UDPPorts      *string `yaml:"udpPorts,omitempty" json:"udpPorts" validate:"omitempty"`
}

type serverConfigACL struct {
	File              *string  `yaml:"file,omitempty" json:"file" validate:"omitempty"`
	Inline            []string `yaml:"inline,omitempty" json:"inline" validate:"omitempty"`
	GeoIP             *string  `yaml:"geoip,omitempty" json:"geoip" validate:"omitempty"`
	GeoSite           *string  `yaml:"geosite,omitempty" json:"geosite" validate:"omitempty"`
	GeoUpdateInterval *string  `yaml:"geoUpdateInterval,omitempty" json:"geoUpdateInterval" validate:"omitempty"`
}

type serverConfigOutboundDirect struct {
	Mode       *string `yaml:"mode,omitempty" json:"mode" validate:"required"`
	BindIPv4   *string `yaml:"bindIPv4,omitempty" json:"bindIPv4" validate:"required"`
	BindIPv6   *string `yaml:"bindIPv6,omitempty" json:"bindIPv6" validate:"required"`
	BindDevice *string `yaml:"bindDevice,omitempty" json:"bindDevice" validate:"required"`
}

type serverConfigOutboundSOCKS5 struct {
	Addr     *string `yaml:"addr,omitempty" json:"addr" validate:"required"`
	Username *string `yaml:"username,omitempty" json:"username" validate:"omitempty"`
	Password *string `yaml:"password,omitempty" json:"password" validate:"omitempty"`
}

type serverConfigOutboundHTTP struct {
	URL      *string `yaml:"url,omitempty" json:"url" validate:"required"`
	Insecure *bool   `yaml:"insecure,omitempty" json:"insecure" validate:"required"`
}

type serverConfigOutboundEntry struct {
	Name   *string                     `yaml:"name,omitempty" json:"name" validate:"required"`
	Type   *string                     `yaml:"type,omitempty" json:"type" validate:"omitempty"`
	Direct *serverConfigOutboundDirect `yaml:"direct,omitempty" json:"direct" validate:"omitempty"`
	SOCKS5 *serverConfigOutboundSOCKS5 `yaml:"socks5,omitempty" json:"socks5" validate:"omitempty"`
	HTTP   *serverConfigOutboundHTTP   `yaml:"http,omitempty" json:"http" validate:"omitempty"`
}

type ServerConfigTrafficStats struct {
	Listen *string `yaml:"listen,omitempty" json:"listen" validate:"required"`
	Secret *string `yaml:"secret,omitempty" json:"-" validate:"omitempty"`
}

type serverConfigMasqueradeFile struct {
	Dir *string `yaml:"dir,omitempty" json:"dir" validate:"required"`
}

type serverConfigMasqueradeProxy struct {
	URL         *string `yaml:"url,omitempty" json:"url" validate:"required"`
	RewriteHost *bool   `yaml:"rewriteHost,omitempty" json:"rewriteHost" validate:"required"`
}

type serverConfigMasqueradeString struct {
	Content    *string           `yaml:"content,omitempty" json:"content" validate:"required"`
	Headers    map[string]string `yaml:"headers,omitempty" json:"headers" validate:"omitempty"`
	StatusCode *int              `yaml:"statusCode,omitempty" json:"statusCode" validate:"omitempty"`
}

type serverConfigMasquerade struct {
	Type        *string                       `yaml:"type,omitempty" json:"type" validate:"omitempty"`
	File        *serverConfigMasqueradeFile   `yaml:"file,omitempty" json:"file" validate:"omitempty"`
	Proxy       *serverConfigMasqueradeProxy  `yaml:"proxy,omitempty" json:"proxy" validate:"omitempty"`
	String      *serverConfigMasqueradeString `yaml:"string,omitempty" json:"string" validate:"omitempty"`
	ListenHTTP  *string                       `yaml:"listenHTTP,omitempty" json:"listenHTTP" validate:"omitempty"`
	ListenHTTPS *string                       `yaml:"listenHTTPS,omitempty" json:"listenHTTPS" validate:"omitempty"`
	ForceHTTPS  *bool                         `yaml:"forceHTTPS,omitempty" json:"forceHTTPS" validate:"omitempty"`
}
