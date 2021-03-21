package helpers

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func NetworkIP(ip string) string {
	arr := strings.Split(ip, ".")
	arr[len(arr) - 1] = "0/24"

	return strings.Join(arr,".")
}

func GetUser() *user.User {
	//uid := os.Getuid()
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return user
}

func ConvertBoolFromString(value string) bool {
	if value == "Yes" {
		return true
	}

	return false
}
