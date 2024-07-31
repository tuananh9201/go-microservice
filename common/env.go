package common

import (
	"fmt"
	"syscall"
)

func EnvString(key, fallback string) string {
	fmt.Println("key", key)
	if val, ok := syscall.Getenv(key); ok {
		return val
	}
	return fallback
}
