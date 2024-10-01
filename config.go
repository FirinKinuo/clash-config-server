package main

type Config struct {
	ExtraPaths map[string]ExtraPath `yaml:"extra_paths"`
}

type ExtraPath struct {
	Path      string    `yaml:"path"`
	BasicAuth BasicAuth `yaml:"basic_auth"`
}

type BasicAuth struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
