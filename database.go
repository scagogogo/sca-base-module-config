package sca_base_module_config

type Database struct {
	Mysql Mysql `yaml:"mysql" mapstructure:"mysql"`
}
