package types

type AppOption struct {
	Debug  bool   `yaml:"debug"`
	Listen string `yaml:"listen"`
	Pool   uint32 `yaml:"pool"`
}
