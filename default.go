package sca_base_module_config

import (
	reflect_utils "github.com/golang-infrastructure/go-reflect-utils"
)

// OrDefault 如果给定的值为空的话则返回默认值
func OrDefault[T any](value, defaultValue T) T {
	if reflect_utils.IsNil(value) {
		return defaultValue
	} else {
		return value
	}
}
