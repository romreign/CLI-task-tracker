package util

import (
	"os"
)

func ParseCli() []string {
	Args := make([]string, len(os.Args)-1)
	copy(Args, os.Args[1:])
	return Args
}
