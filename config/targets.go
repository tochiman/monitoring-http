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
		HttpConfig{Name: "Google", Host: "google.com", Path: "/", Proto: "http", Domain: "google.com"},
		HttpConfig{Name: "Google", Host: "google.com", Path: "/no-such-path", Proto: "http", Domain: "google.com"},
		HttpConfig{Name: "NiseGoogle", Host: "nisegoogle.com", Path: "/", Proto: "http", Domain: "nisegoogle.com"},
	}
}
