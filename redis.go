package sca_base_module_config

type Redis struct {
	AutoInit *bool  `yaml:"auto-init"  mapstructure:"auto-init"`
	Address  string `yaml:"address" mapstructure:"address"`
	Passwd   string `yaml:"passwd" mapstructure:"passwd"`
}

func (x Redis) IsOk() bool {
	return x.Address != ""
}
