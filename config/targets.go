package config

type HttpConfigSlice struct {
	Http []HttpConfig
}

type HttpConfig struct {
	Name   string
	Host   string
	Path   string
	Proto  string
	Domain string
}

func HttpTargets() []HttpConfig {
	return []HttpConfig{
		HttpConfig{Name: "Nextcloud", Host: "nextcloud.tochiman.com", Path: "/", Proto: "https", Domain: "nextcloud.tochiman.com"},
		HttpConfig{Name: "CodiMD", Host: "codimd.tochiman.com", Path: "/", Proto: "https", Domain: "codimd.tochiman.com"},
	}
}
