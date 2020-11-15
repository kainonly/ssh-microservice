package options

type Ssh struct {
	Host       string
	Port       uint32
	Username   string
	Password   string
	Key        []byte
	PassPhrase []byte
}
