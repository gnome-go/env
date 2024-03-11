//go:build !windows

package env

import (
	"errors"
	"os"
)

func SetWin(key string, value string, target int) error {
	return errors.New("windows only")
}

func GetWin(key string, target int) (string, error) {
	return "", errors.New("windows only")
}

func Username() string {
	return os.Getenv("USER")
}

func GetPath() string {
	return os.Getenv("PATH")
}

func SetPath(path string) error {
	return os.Setenv("PATH", path)
}

func HasPath(path string) bool {
	return hasPath(path, SplitPath())
}

func AppendPath(path string) error {
	paths := SplitPath()
	if hasPath(path, paths) {
		return nil
	}
	paths = append(paths, path)
	return SetPath(JoinPath(paths...))
}

func PrependPath(path string) error {
	paths := SplitPath()
	if hasPath(path, paths) {
		return nil
	}
	paths = append([]string{path}, paths...)
	return SetPath(JoinPath(paths...))
}

func hasPath(path string, paths []string) bool {
	for _, p := range paths {
		if p == path {
			return true
		}
	}
	return false
}
