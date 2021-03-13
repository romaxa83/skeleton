package helpers

func ConvertBoolFromString(value string) bool {
	if value == "Yes" {
		return true
	}

	return false
}
