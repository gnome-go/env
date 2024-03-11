//go:build windows

package env

import (
	"errors"
	"os"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func Username() string {
	return os.Getenv("USERNAME")
}

func GetPath() string {
	return os.Getenv("Path")
}

func SetWin(key string, value string, target int) error {
	switch target {
	case WinProcess:
		return os.Setenv(key, value)

	case WinMachine:
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.SET_VALUE)
		if err != nil {
			return err
		}

		defer k.Close()

		return k.SetStringValue(key, value)

	case WinUser:
		k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.SET_VALUE)
		if err != nil {
			return err
		}

		defer k.Close()

		return k.SetStringValue(key, value)
	default:

		return errors.New("invalid target: " + getTargetName(target))
	}
}

func GetWin(key string, target int) (string, error) {
	switch target {
	case WinProcess:
		return os.Getenv(key), nil

	case WinMachine:
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.QUERY_VALUE)
		if err != nil {
			return "", err
		}

		defer k.Close()

		v, _, err := k.GetStringValue(key)
		if err != nil {
			return "", err
		}

		return v, nil

	case WinUser:
		k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE)
		if err != nil {
			return "", err
		}

		defer k.Close()

		v, _, err := k.GetStringValue(key)
		if err != nil {
			return "", err
		}

		return v, nil

	default:
		return "", errors.New("invalid target: " + getTargetName(target))
	}
}

func SetPath(path string) error {
	return os.Setenv("Path", path)
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

func getTargetName(target int) string {
	switch target {
	case WinProcess:
		return "process"
	case WinMachine:
		return "machine"
	case WinUser:
		return "user"
	default:
		return "unknown"
	}
}

func hasPath(path string, paths []string) bool {
	for _, p := range paths {
		if strings.EqualFold(p, path) {
			return true
		}
	}
	return false
}
