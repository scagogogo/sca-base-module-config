package sca_base_module_config

type COS struct {
	Endpoint        string `yaml:"endpoint" mapstructure:"endpoint"`
	SecretId     string `yaml:"secret-id" mapstructure:"secret-id"`
	SecretKey string `yaml:"secret-key" mapstructure:"secret-key"`
}
