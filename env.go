package sca_base_module_config

import (
	"github.com/golang-infrastructure/go-pointer"
	"os"
	"strconv"
	"strings"
)

const ScaAutoInitConfigEnvName = "SCA_AUTO_INIT_CONFIG"

func GetEnvString(name string) *string {
	env, b := os.LookupEnv(name)
	if !b {
		return nil
	}
	return pointer.ToPointer(env)
}

func GetEnvStringOrDefault(name, defaultValue string) string {
	return pointer.FromPointerOrDefault(GetEnvString(name), defaultValue)
}

func GetEnvBool(name string) *bool {
	env, b := os.LookupEnv(name)
	if !b {
		return nil
	}
	return pointer.ToPointer(strings.ToLower(env) == "true")
}

func GetEnvBoolOrDefault(name string, defaultValue bool) bool {
	return pointer.FromPointerOrDefault(GetEnvBool(name), defaultValue)
}

func GetEnvInt(name string) *int {
	env, b := os.LookupEnv(name)
	if !b {
		return nil
	}
	atoi, err := strconv.Atoi(env)
	if err != nil {
		return nil
	}
	return pointer.ToPointer(atoi)
}
func GetEnvIntOrDefault(name string, defaultValue int) int {
	return pointer.FromPointerOrDefault(GetEnvInt(name), defaultValue)
}
