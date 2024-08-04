package getenv

import (
	"os"
	"strconv"

	"github.com/akramarenkov/goresearch/internal/consts"
)

const (
	sizeKey = "SIZE"
)

func Size(defaultSize int) (int, error) {
	env := os.Getenv(consts.EnvPrefix + sizeKey)

	if env == "" {
		return defaultSize, nil
	}

	return strconv.Atoi(env)
}
