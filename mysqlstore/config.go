package mysqlstore

type Config struct {
	DSN     string `yaml:"dsn"`
	Level   string `yaml:"level"`
	MaxConn int    `yaml:"max_conn"`
	MaxOpen int    `yaml:"max_open"`
	Timeout int64  `yaml:"timeout"`
}
