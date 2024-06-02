package utils

import "fmt"

func InterfaceToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}
