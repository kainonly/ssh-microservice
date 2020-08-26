package types

type Config struct {
	Debug  bool   `yaml:"debug"`
	Listen string `yaml:"listen"`
	Pool   uint32 `yaml:"pool"`
}
