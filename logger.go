package sca_base_module_config

type Logger struct {
	Enable    LoggerEnable `yaml:"enable" mapstructure:"enable"`
	Directory string       `yaml:"directory" mapstructure:"directory"`
}

type LoggerEnable struct {
	Stdout *bool `yaml:"stdout" mapstructure:"stdout"`
	File   *bool `yaml:"file" mapstructure:"file"`
}
