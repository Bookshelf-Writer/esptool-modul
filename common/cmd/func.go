package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

//###########################################################//

func Run(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s: %v", out.String(), err)
	}
	return out.String(), nil
}
