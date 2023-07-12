package sca_base_module_config

import (
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
	"github.com/spf13/viper"
)

const DefaultConfigName = "config"

var Config *Configuration

// TODO 此处的自动初始化处理得不好
func init() {
	// 如果配置了不自动初始化，则跳过此初始化过程
	if !GetEnvBoolOrDefault(ScaAutoInitConfigEnvName, true) {
		return
	}
	err := InitDefault()
	if err != nil {
		panic(err)
	}
}

// InitDefault 初始化默认
func InitDefault() error {
	var err error
	Config, err = Init(DefaultConfigName)
	return err
}

// Init 初始化配置文件
func Init(configName string) (*Configuration, error) {

	// 如果是相对路径的话，认为是从项目根目录开始的相对路径
	projectRootPath, err := project_root_directory.GetRootDirectory()
	if err != nil {
		return nil, err
	}

	configFileViper := viper.New()
	configFileViper.AddConfigPath(projectRootPath) //设置读取的文件路径
	configFileViper.SetConfigName(configName)      //设置读取的文件名
	configFileViper.SetConfigType("yaml")          //设置文件的类型
	//尝试进行配置读取
	if err := configFileViper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Configuration{}
	err = configFileViper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	config.Viper = configFileViper

	return config, nil
}
