package options

type ClientOption struct {
	Identity   string         `yaml:"identity"`
	Host       string         `yaml:"host"`
	Port       uint32         `yaml:"port"`
	Username   string         `yaml:"username"`
	Password   string         `yaml:"password"`
	PrivateKey string         `yaml:"key"`        // base64 string
	Passphrase string         `yaml:"passphrase"` // base64 string
	Tunnels    []TunnelOption `yaml:"tunnels"`
}
