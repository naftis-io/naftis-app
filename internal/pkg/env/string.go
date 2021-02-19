package env

import "os"

func String(key string, defaultValue string) string {
	v := os.Getenv(key)

	if len(v) == 0 {
		v = defaultValue
	}

	return v
}
