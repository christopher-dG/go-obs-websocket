package obsws

type Config struct {
	Addr           string `yaml:"addr"`
	Password       string `yaml:"password"`
	ReceiveTimeout string `yaml:"receiveTimeout"`
}
