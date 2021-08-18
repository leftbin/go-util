package shell

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func RunCmd(cmd *exec.Cmd) error {
	log.Infof("\nrunning command %s\n\n", cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
