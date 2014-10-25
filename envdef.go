package envdef

import (
	"os"
)

func Get(key, def string) string {
	val := os.Getenv(key)
	if len(val) > 0 {
		return val
	}

	return def
}

func GetWithExpand(key, def string) string {
	return os.ExpandEnv(Get(key, def))
}
