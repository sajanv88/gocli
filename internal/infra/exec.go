package infra

import (
	"errors"
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

func ResolveNodePackageManager() (string, error) {
	if err := CheckToolAvailable("npm", "Node 20.19+/22.12+"); err == nil {
		return "npm", nil
	}
	if err := CheckToolAvailable("pnpm", "8+"); err == nil {
		return "pnpm", nil
	}

	return "", errors.New("neither npm nor pnpm found on PATH — install Node.js or pnpm")
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
