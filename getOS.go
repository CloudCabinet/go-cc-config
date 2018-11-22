package config

import "os"

func getOS(keyos string, def string) string {
	value, err := os.LookupEnv(keyos)
	if (err) {
		return value
	} else {
		return def
	}
}
