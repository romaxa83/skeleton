package helpers

import "strings"

func NetworkIP(ip string) string {
	arr := strings.Split(ip, ".")
	arr[len(arr) - 1] = "0/24"

	return strings.Join(arr,".")
}

func ConvertBoolFromString(value string) bool {
	if value == "Yes" {
		return true
	}

	return false
}
