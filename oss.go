package sca_base_module_config

type OSS struct {
	Endpoint        string `yaml:"endpoint" mapstructure:"endpoint"`
	AccessKeyId     string `yaml:"access-key-id" mapstructure:"access-key-id"`
	AccessKeySecret string `yaml:"access-key-secret" mapstructure:"access-key-secret"`
}
