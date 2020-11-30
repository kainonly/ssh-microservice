package config

type Config struct {
	Debug   string `yaml:"debug"`
	Listen  string `yaml:"listen"`
	Gateway string `yaml:"gateway"`
}
