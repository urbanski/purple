package purple

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

type BashExecutor struct {
	Exec TestExecutionSpec
}

func (b *BashExecutor) Init (spec TestExecutionSpec) {
	b.Exec = spec
}

func (b *BashExecutor) Run() error {
	// run pre-req commands if they exist
	if b.Exec.PrereqCommand != "" {
		b.exec(b.Exec.PrereqCommand)
	}

	// run cmd if it exists
	if b.Exec.Command == "" {
		return errors.New("command not specified")
	}
	b.exec(b.Exec.Command)

	// run cleanup if it exists
	if b.Exec.CleanupCommand != "" {
		b.exec(b.Exec.CleanupCommand)
	}
	return nil
}

func (b *BashExecutor) exec(command string) {
	log.Info("Executing ", command)
	_, err := osExec(command, true)
	if err != nil {
		log.Error(err)
	}
}