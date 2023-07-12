package sca_base_module_config

type Crawler struct {
	Proxy CrawlerProxy `yaml:"proxy" mapstructure:"proxy"`
	Queue CrawlerQueue `yaml:"queue" mapstructure:"queue"`
	Consumer CrawlerConsumer `yaml:"consumer" mapstructure:"consumer"`
}

type CrawlerQueue struct {
	Name string `yaml:"name" mapstructure:"name"`
}

type CrawlerConsumer struct {
	WorkerNum int `yaml:"worker-num" mapstructure:"worker-num"`
}
