// +build !windows

package unixexec

import (
	"os/exec"
)

func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
