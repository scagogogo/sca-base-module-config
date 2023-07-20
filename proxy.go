package sca_base_module_config

import (
	"context"
	"fmt"
	"github.com/crawler-go-go-go/go-requests"
	"strings"
)

// CrawlerProxy 代理设置
type CrawlerProxy struct {

	// 动态代理
	DynamicProxy string `json:"dynamic-proxy" mapstructure:"dynamic-proxy"`

	// 提取型短效代理，配置的是一个接口的URL，待认证，对这个接口发送GET请求返回样例如下：
	// 1.1.1.1:10086
	// 2.3.4.5:10086
	// 192.168.178.1:10086
	// 198.3.134.156:10086
	ShortProxy string `json:"short-proxy" mapstructure:"short-proxy"`
}

// AcquireShortProxy 尝试提取短效代理
func (x CrawlerProxy) AcquireShortProxy(ctx context.Context) ([]string, error) {
	if x.ShortProxy == "" {
		return nil, fmt.Errorf("not config short proxy")
	}
	responseBody, err := requests.GetString(ctx, x.ShortProxy)
	if err != nil {
		return nil, err
	}
	split := strings.Split(responseBody, "\n")
	proxySlice := make([]string, 0)
	for _, line := range split {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		proxySlice = append(proxySlice, line)
	}
	return proxySlice, nil
}

// AcquireOneShortProxy 尝试获取一个短效代理
func (x CrawlerProxy) AcquireOneShortProxy(ctx context.Context) (string, error) {
	proxySlice, err := x.AcquireShortProxy(ctx)
	if err != nil {
		return "", err
	}
	if len(proxySlice) == 0 {
		return "", nil
	}
	return proxySlice[0], nil
}
