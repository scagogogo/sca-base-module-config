package sca_base_module_config

import "github.com/spf13/viper"

// Configuration 对应着一个映射文件
type Configuration struct {
	Database Database `yaml:"database" mapstructure:"database"`
	Redis    Redis    `yaml:"redis" mapstructure:"redis"`
	Logger   Logger   `yaml:"logger" mapstructure:"logger"`
	OSS      OSS      `yaml:"oss" mapstructure:"oss"`
	COS      COS      `yaml:"cos" mapstructure:"cos"`
	Crawler  Crawler  `yaml:"crawler" mapstructure:"crawler"`

	// 把原来的viper也获取一下放在这里，随时都可以通过viper的方法来获取
	*viper.Viper

	// 未明确声明的就放在这里吧...
	ViperRemainMap map[string]any `mapstructure:",remain"`
}

func (x *Configuration) GetStringOrDefault(name, defaultValue string) string {
	return OrDefault(x.Viper.GetString(name), defaultValue)
}
