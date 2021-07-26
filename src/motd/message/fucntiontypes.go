package message

import "fmt"

func greeting(name, message string) string {
	return fmt.Sprint("%s, %s", message, name)
}