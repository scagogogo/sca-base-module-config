package sca_base_module_config

type Mysql struct {

	// Example: "xxxx:xxxxxxxx@tcp(xxxx)/xxxx?parseTime=true&loc=Local&charset=utf8mb4"
	DSN string `yaml:"dsn" mapstructure:"dsn"`
}
