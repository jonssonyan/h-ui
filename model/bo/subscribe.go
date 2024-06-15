package bo

type Hysteria2 struct {
	Name           string `yaml:"name"`
	Type           string `yaml:"type"`
	Server         string `yaml:"server"`
	Port           uint   `yaml:"port"`
	Up             int    `yaml:"up,omitempty"`
	Down           int    `yaml:"down,omitempty"`
	Password       string `yaml:"password"`
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
