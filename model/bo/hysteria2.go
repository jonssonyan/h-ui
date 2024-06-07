package bo

type Hysteria2ServerConfig struct {
	Listen                *string                     `yaml:"listen" json:"listen" validate:"required"`
	Obfs                  *serverConfigObfs           `yaml:"obfs" json:"obfs" validate:"omitempty"`
	TLS                   *serverConfigTLS            `yaml:"tls" json:"tls" validate:"omitempty"`
	ACME                  *serverConfigACME           `yaml:"acme" json:"acme" validate:"omitempty"`
	QUIC                  *serverConfigQUIC           `yaml:"quic" json:"quic" validate:"omitempty"`
	Bandwidth             *serverConfigBandwidth      `yaml:"bandwidth" json:"bandwidth" validate:"omitempty"`
	IgnoreClientBandwidth *bool                       `yaml:"ignoreClientBandwidth" json:"ignoreClientBandwidth" validate:"omitempty"`
	SpeedTest             *bool                       `yaml:"speedTest" json:"speedTest" validate:"omitempty"`
	DisableUDP            *bool                       `yaml:"disableUDP" json:"disableUDP" validate:"omitempty"`
	UDPIdleTimeout        *string                     `yaml:"udpIdleTimeout" json:"udpIdleTimeout" validate:"omitempty"`
	Auth                  *ServerConfigAuth           `yaml:"auth" json:"auth" validate:"omitempty"`
	Resolver              *serverConfigResolver       `yaml:"resolver" json:"resolver" validate:"omitempty"`
	ACL                   *serverConfigACL            `yaml:"acl" json:"acl" validate:"omitempty"`
	Outbounds             []serverConfigOutboundEntry `yaml:"outbounds" json:"outbounds" validate:"omitempty"`
	TrafficStats          *ServerConfigTrafficStats   `yaml:"trafficStats" json:"trafficStats" validate:"required"`
	Masquerade            *serverConfigMasquerade     `yaml:"masquerade" json:"masquerade" validate:"omitempty"`
}

type serverConfigObfsSalamander struct {
	Password *string `yaml:"password" json:"password" validate:"required"`
}

type serverConfigObfs struct {
	Type       *string                     `yaml:"type" json:"type" validate:"required"`
	Salamander *serverConfigObfsSalamander `yaml:"salamander" json:"salamander" validate:"required"`
}

type serverConfigTLS struct {
	Cert *string `yaml:"cert" json:"cert" validate:"required"`
	Key  *string `yaml:"key"  json:"key" validate:"required"`
}

type serverConfigACME struct {
	Domains        []string `yaml:"domains" json:"domains" validate:"required"`
	Email          *string  `yaml:"email" json:"email" validate:"required"`
	CA             *string  `yaml:"ca" json:"ca" validate:"required"`
	DisableHTTP    *bool    `yaml:"disableHTTP" json:"disableHTTP" validate:"required"`
	DisableTLSALPN *bool    `yaml:"disableTLSALPN" json:"disableTLSALPN" validate:"required"`
	ListenHost     *string  `yaml:"listenHost" json:"listenHost" validate:"required"`
	AltHTTPPort    *int     `yaml:"altHTTPPort" json:"altHTTPPort" validate:"required"`
	AltTLSALPNPort *int     `yaml:"altTLSALPNPort" json:"altTLSALPNPort" validate:"required"`
	Dir            *string  `yaml:"dir" json:"dir" validate:"required"`
}

type serverConfigQUIC struct {
	InitStreamReceiveWindow     *uint64 `yaml:"initStreamReceiveWindow" json:"initStreamReceiveWindow" validate:"omitempty"`
	MaxStreamReceiveWindow      *uint64 `yaml:"maxStreamReceiveWindow" json:"maxStreamReceiveWindow" validate:"omitempty"`
	InitConnectionReceiveWindow *uint64 `yaml:"initConnReceiveWindow" json:"initConnReceiveWindow" validate:"omitempty"`
	MaxConnectionReceiveWindow  *uint64 `yaml:"maxConnReceiveWindow" json:"maxConnReceiveWindow" validate:"omitempty"`
	MaxIdleTimeout              *string `yaml:"maxIdleTimeout" json:"maxIdleTimeout" validate:"omitempty"`
	MaxIncomingStreams          *int64  `yaml:"maxIncomingStreams" json:"maxIncomingStreams" validate:"omitempty"`
	DisablePathMTUDiscovery     *bool   `yaml:"disablePathMTUDiscovery" json:"disablePathMTUDiscovery" validate:"omitempty"`
}

type serverConfigBandwidth struct {
	Up   *string `yaml:"up" json:"up" validate:"required"`
	Down *string `yaml:"down" json:"down" validate:"required"`
}

type ServerConfigAuthHTTP struct {
	URL      *string `yaml:"url" json:"url" validate:"required"`
	Insecure *bool   `yaml:"insecure" json:"insecure" validate:"required"`
}

type ServerConfigAuth struct {
	Type     *string               `yaml:"type" json:"type" validate:"omitempty"`
	Password *string               `yaml:"password" json:"password" validate:"omitempty"`
	UserPass map[string]string     `yaml:"userpass" json:"userpass" validate:"omitempty"`
	HTTP     *ServerConfigAuthHTTP `yaml:"http" json:"http" validate:"omitempty"`
	Command  *string               `yaml:"command" json:"command" validate:"omitempty"`
}

type serverConfigResolverTCP struct {
	Addr    *string `yaml:"addr" json:"addr" validate:"required"`
	Timeout *string `yaml:"timeout" json:"timeout" validate:"required"`
}

type serverConfigResolverUDP struct {
	Addr    *string `yaml:"addr" json:"addr" validate:"required"`
	Timeout *string `yaml:"timeout" json:"timeout" validate:"required"`
}

type serverConfigResolverTLS struct {
	Addr     *string `yaml:"addr" json:"addr" validate:"required"`
	Timeout  *string `yaml:"timeout" json:"timeout" validate:"required"`
	SNI      *string `yaml:"sni" json:"sni" validate:"required"`
	Insecure *bool   `yaml:"insecure" json:"insecure" validate:"required"`
}

type serverConfigResolverHTTPS struct {
	Addr     *string `yaml:"addr" json:"addr" validate:"required"`
	Timeout  *string `yaml:"timeout" json:"timeout" validate:"required"`
	SNI      *string `yaml:"sni" json:"sni"  validate:"required"`
	Insecure *bool   `yaml:"insecure" json:"insecure" validate:"required"`
}

type serverConfigResolver struct {
	Type  *string                    `yaml:"type" json:"type" validate:"omitempty"`
	TCP   *serverConfigResolverTCP   `yaml:"tcp" json:"tcp" validate:"omitempty"`
	UDP   *serverConfigResolverUDP   `yaml:"udp" json:"udp" validate:"omitempty"`
	TLS   *serverConfigResolverTLS   `yaml:"tls" json:"tls" validate:"omitempty"`
	HTTPS *serverConfigResolverHTTPS `yaml:"https" json:"https" validate:"omitempty"`
}

type serverConfigACL struct {
	File              *string  `yaml:"file" json:"file" validate:"omitempty"`
	Inline            []string `yaml:"inline" json:"inline" validate:"omitempty"`
	GeoIP             *string  `yaml:"geoip" json:"geoip" validate:"omitempty"`
	GeoSite           *string  `yaml:"geosite" json:"geosite" validate:"omitempty"`
	GeoUpdateInterval *string  `yaml:"geoUpdateInterval" json:"geoUpdateInterval" validate:"omitempty"`
}

type serverConfigOutboundDirect struct {
	Mode       *string `yaml:"mode" json:"mode" validate:"required"`
	BindIPv4   *string `yaml:"bindIPv4" json:"bindIPv4" validate:"required"`
	BindIPv6   *string `yaml:"bindIPv6" json:"bindIPv6" validate:"required"`
	BindDevice *string `yaml:"bindDevice" json:"bindDevice" validate:"required"`
}

type serverConfigOutboundSOCKS5 struct {
	Addr     *string `yaml:"addr" json:"addr" validate:"required"`
	Username *string `yaml:"username" json:"username" validate:"omitempty"`
	Password *string `yaml:"password" json:"password" validate:"omitempty"`
}

type serverConfigOutboundHTTP struct {
	URL      *string `yaml:"url" json:"url" validate:"required"`
	Insecure *bool   `yaml:"insecure" json:"insecure" validate:"required"`
}

type serverConfigOutboundEntry struct {
	Name   *string                     `yaml:"name" json:"name" validate:"required"`
	Type   *string                     `yaml:"type" json:"type" validate:"omitempty"`
	Direct *serverConfigOutboundDirect `yaml:"direct" json:"direct" validate:"omitempty"`
	SOCKS5 *serverConfigOutboundSOCKS5 `yaml:"socks5" json:"socks5" validate:"omitempty"`
	HTTP   *serverConfigOutboundHTTP   `yaml:"http" json:"http" validate:"omitempty"`
}

type ServerConfigTrafficStats struct {
	Listen *string `yaml:"listen" json:"listen" validate:"required"`
	Secret *string `yaml:"secret" json:"secret" validate:"omitempty"`
}

type serverConfigMasqueradeFile struct {
	Dir *string `yaml:"dir" json:"dir" validate:"required"`
}

type serverConfigMasqueradeProxy struct {
	URL         *string `yaml:"url" json:"url" validate:"required"`
	RewriteHost *bool   `yaml:"rewriteHost" json:"rewriteHost" validate:"required"`
}

type serverConfigMasqueradeString struct {
	Content    *string           `yaml:"content" json:"content" validate:"required"`
	Headers    map[string]string `yaml:"headers" json:"headers" validate:"omitempty"`
	StatusCode *int              `yaml:"statusCode" json:"statusCode" validate:"omitempty"`
}

type serverConfigMasquerade struct {
	Type        *string                       `yaml:"type" json:"type" validate:"omitempty"`
	File        *serverConfigMasqueradeFile   `yaml:"file" json:"file" validate:"omitempty"`
	Proxy       *serverConfigMasqueradeProxy  `yaml:"proxy" json:"proxy" validate:"omitempty"`
	String      *serverConfigMasqueradeString `yaml:"string" json:"string" validate:"omitempty"`
	ListenHTTP  *string                       `yaml:"listenHTTP" json:"listenHTTP" validate:"omitempty"`
	ListenHTTPS *string                       `yaml:"listenHTTPS" json:"listenHTTPS" validate:"omitempty"`
	ForceHTTPS  *bool                         `yaml:"forceHTTPS" json:"forceHTTPS" validate:"omitempty"`
}
