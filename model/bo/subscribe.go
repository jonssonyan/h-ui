package bo

type Hysteria2 struct {
	Name           string `yaml:"name"`
	Type           string `yaml:"type"`
	Server         string `yaml:"server"`
	Port           string `yaml:"port"`
	Ports          string `yaml:"ports"`
	Password       string `yaml:"password"`
	Up             string `yaml:"up,omitempty"`
	Down           string `yaml:"down,omitempty"`
	Obfs           string `yaml:"obfs,omitempty"`
	ObfsPassword   string `yaml:"obfs-password,omitempty"`
	Sni            string `yaml:"sni,omitempty"`
	SkipCertVerify bool   `yaml:"skip-cert-verify,omitempty"`
}

type ProxyGroup struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}

type ClashConfig struct {
	Proxies     []interface{} `yaml:"proxies"`
	ProxyGroups []ProxyGroup  `yaml:"proxy-groups"`
}
