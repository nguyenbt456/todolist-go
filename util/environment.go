package util

import (
	"os"
	"strconv"
)

func getEVString(key string, flashback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return flashback
	}
	return value
}

// EVString get environment variable string
func EVString(key string, flashback string) string {
	return getEVString(key, flashback)
}

func getEVInt64(key string, flashback int64) int64 {
	value := os.Getenv(key)

	if len(value) == 0 {
		return flashback
	}

	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

// EVInt64 get environment variable int64
func EVInt64(key string, flashback int64) int64 {
	return getEVInt64(key, flashback)
}

// EVInt get environment variable int
func EVInt(key string, flashback int) int {
	return int(getEVInt64(key, int64(flashback)))
}
