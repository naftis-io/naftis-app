package env

import (
	"os"
	"strconv"
)

func Uint64(key string, defaultValue uint64) uint64 {
	vStr := os.Getenv(key)

	v, err := strconv.ParseUint(vStr, 10, 64)
	if err != nil {
		return defaultValue
	}

	return uint64(v)
}
