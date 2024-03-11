package env

import (
	"os"
	"strings"
)

const (
	WinProcess = 0
	WinMachine = 2
	WinUser    = 1
)

func Get(key string) string {
	return os.Getenv(key)
}

func Set(key, value string) error {
	return os.Setenv(key, value)
}

func Delete(key string) error {
	return os.Unsetenv(key)
}

func Has(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

func GetAll() map[string]string {
	kv := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if len(pair) == 2 && len(pair[1]) > 0 {
			kv[pair[0]] = pair[1]
		}
	}

	return kv
}

func TempDir() string {
	return os.TempDir()
}

func UserHomeDr() (string, error) {
	return os.UserHomeDir()
}

func UserCacheDr() (string, error) {
	return os.UserCacheDir()
}

func UserConfigDr() (string, error) {
	return os.UserConfigDir()
}

func Hostname() (string, error) {
	return os.Hostname()
}

func SplitPath() []string {
	return strings.Split(GetPath(), string(os.PathListSeparator))
}

func JoinPath(paths ...string) string {
	return strings.Join(paths, string(os.PathListSeparator))
}
