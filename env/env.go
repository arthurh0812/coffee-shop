package env

import "os"

func init() {
	initEnv()
}

var envMapping = map[string]string{
	"PORT": "port",
	"HOST": "host",
}

var Env = make(map[string]string)

func initEnv() {
	for key, val := range envMapping {
		Env[val] = os.Getenv(key)
	}
}
