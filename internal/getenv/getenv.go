package getenv

import (
	"os"
	"strconv"
)

const (
	prefix = "GORESEARCH_"
)

const (
	keySize      = "SIZE"
	keyTerminate = "TERMINATE"
)

func envVar(key string) string {
	return prefix + key
}

func Size(def int) (int, error) {
	return Int(keySize, def)
}

func Terminate(def bool) (bool, error) {
	return Boolean(keyTerminate, def)
}

func Int(key string, def int) (int, error) {
	env := os.Getenv(prefix + key)

	if env == "" {
		return def, nil
	}

	return strconv.Atoi(env)
}

func Boolean(key string, def bool) (bool, error) {
	env := os.Getenv(prefix + key)

	if env == "" {
		return def, nil
	}

	return strconv.ParseBool(env)
}
