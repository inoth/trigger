package util

import (
	"strings"

	"github.com/google/uuid"
)

func UUID(ns ...int) string {
	n := 16
	if len(ns) > 0 {
		n = ns[0]
	}
	uuidStr := uuid.New().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	return uuidStr[0:n]
}
