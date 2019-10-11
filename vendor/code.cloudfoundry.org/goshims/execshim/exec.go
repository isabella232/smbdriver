// This file was generated by counterfeiter
// with command: counterfeiter -p os/exec
// BUT THEN WE MODIFIED IT SO MAKE SURE TO COPY THOSE MODIFICATIONS FORWARDS
package execshim

import (
	"context"
)

//go:generate counterfeiter -o exec_fake/fake_exec.go . Exec
type Exec interface {
	Command(name string, arg ...string) Cmd
	CommandContext(ctx context.Context, name string, arg ...string) Cmd
	LookPath(file string) (string, error)
}
