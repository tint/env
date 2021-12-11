package env

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Path 根据应用根目录获取路径
func Path(paths ...string) string {
	if len(paths) == 0 {
		return RootPath
	}
	return filepath.Join(RootPath, strings.Join(paths, "/"))
}

func Lookup(name string) (val string, exists bool) {
	val, exists = os.LookupEnv(name)
	if exists && len(val) == 0 {
		exists = false
	}
	return
}

func Exists(name string) bool {
	_, exists := Lookup(name)
	return exists
}

func Unknown(name string, value ...interface{}) interface{} {
	if val, ok := Lookup(name); ok {
		return val
	}
	if len(value) > 0 {
		return value[0]
	}
	return nil
}

func String(name string, value ...string) string {
	if val, ok := Lookup(name); ok {
		return val
	}
	if len(value) > 0 {
		return value[0]
	}
	return ""
}

func Bytes(name string, value ...[]byte) []byte {
	if val, ok := Lookup(name); ok {
		return []byte(val)
	}
	if len(value) > 0 {
		return value[0]
	}
	return []byte{}
}

func Int(name string, value ...int) int {
	if val, ok := Lookup(name); ok {
		n, err := strconv.Atoi(val)
		if err == nil {
			return n
		}
	}
	if len(value) > 0 {
		return value[0]
	}
	return 0
}

func Duration(name string, value ...time.Duration) time.Duration {
	if val, ok := Lookup(name); ok {
		n, err := strconv.Atoi(val)
		if err == nil {
			return time.Duration(n)
		}
	}
	if len(value) > 0 {
		return value[0]
	}
	return time.Duration(0)
}

func Bool(name string, value ...bool) bool {
	if val, ok := Lookup(name); ok {
		bl, err := strconv.ParseBool(val)
		if err == nil {
			return bl
		}
	}
	if len(value) > 0 {
		return value[0]
	}
	return false
}

func Map(prefix string) map[string]string {
	prefix = prefix + "_"
	result := map[string]string{}
	for _, value := range os.Environ() {
		if strings.HasPrefix(value, prefix) {
			parts := strings.SplitN(value, "=", 2)
			name := strings.TrimPrefix(parts[0], prefix)
			result[name] = strings.TrimSpace(parts[1])
		}
	}
	return result
}

func Where(filter func(name string, value string) bool) map[string]string {
	result := map[string]string{}
	for _, value := range os.Environ() {
		parts := strings.SplitN(value, "=", 2)
		name := parts[0]
		val := strings.TrimSpace(parts[1])
		if filter(name, val) {
			result[name] = val
		}
	}
	return result
}

func Environ() []string {
	return os.Environ()
}
