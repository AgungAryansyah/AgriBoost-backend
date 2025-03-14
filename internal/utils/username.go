package utils

import "strings"

func GetUsername(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) > 1 {
		return parts[0]
	}
	return ""
}
