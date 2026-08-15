package infra

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CheckToolAvailable(tool, minVersionHint string) error {
	if _, err := exec.LookPath(tool); err != nil {
		return fmt.Errorf("%s is not found on PATH - install it (%s) first", tool, minVersionHint)
	}
	return nil
}

func Run(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s failed: %s", strings.Join(append([]string{name}, args...), " "), err)
	}
	return nil
}
