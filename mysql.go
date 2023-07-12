package sca_base_module_config

type Mysql struct {

	// 自动初始化MySQL驱动
	AutoInit *bool `yaml:"auto-init" mapstructure:"auto-init"`

	// Example: "xxxx:xxxxxxxx@tcp(xxxx)/xxxx?parseTime=true&loc=Local&charset=utf8mb4"
	DSN string `yaml:"dsn" mapstructure:"dsn"`

	Driver MysqlDriverEnable `yaml:"driver" mapstructure:"driver"`
}

type MysqlDriverEnable struct {
	SqlX *bool `yaml:"sqlx" mapstructure:"sqlx"`
	Gorm *bool `yaml:"gorm" mapstructure:"gorm"`
}
