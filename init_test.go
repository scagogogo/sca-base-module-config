package sca_base_module_config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	c, err := Init("config.example")
	assert.Nil(t, err)
	assert.NotNil(t, c)
}
