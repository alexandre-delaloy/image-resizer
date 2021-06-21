package helpers

import (
	"os"
	"strings"
)

func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "Create user"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
