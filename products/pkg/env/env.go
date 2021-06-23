package env

import (
	"os"
	"strings"
)

var envList = []string{"PORT", "HOST"}

var Env = make(map[string]string)

func init() {
	for _, s := range envList {
		Env[strings.ToLower(s)] = os.Getenv(s)
	}
}