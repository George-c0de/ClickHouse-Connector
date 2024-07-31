package configuration

type Config struct {
	Database     string `env:"DATABASE"`
	Username     string `env:"Username"`
	Password     string `env:"PASSWORD,unset"`
	IsProduction bool   `env:"PRODUCTION" envDefault:"false"`
	PortClick    string `env:"PORT_CLICK" envDefault:"8123"`
	AddrClick    string `env:"ADDR_CLICK"`
}
